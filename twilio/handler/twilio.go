package handler

import (
	"context"
	"github.com/gotidy/ptr"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/errors"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	log "github.com/micro/micro/v3/service/logger"

	pb "github.com/ansg191/northstars-backend/twilio/proto"
)

const RFC2822 = "Mon, 02 Jan 2006 15:04:05 -0700"

type Twilio struct{}

// SendMessage is a single request handler called via client.Call or the generated client code
func (e *Twilio) SendMessage(ctx context.Context, req *pb.SendMessageRequest, rsp *pb.SendMessageResponse) error {
	log.Info("Received Twilio.SendMessage request")

	opts, err := GetTwilioOpts()
	if err != nil {
		return err
	}

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: opts.sid,
		Password: opts.token,
	})

	params := &openapi.CreateMessageParams{}
	params.SetBody(req.Body)
	params.SetTo(req.To)
	params.SetMessagingServiceSid(opts.msgServiceSid)

	resp, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		return err
	}

	msg, err := TwilioMsgToMessage(resp)
	if err != nil {
		return err
	}

	rsp.Msg = msg

	return nil
}

func (e *Twilio) GetMessage(ctx context.Context, req *pb.GetMessageRequest, rsp *pb.GetMessageResponse) error {
	log.Info("Received Twilio.GetMessage request")

	opts, err := GetTwilioOpts()
	if err != nil {
		return err
	}

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: opts.sid,
		Password: opts.token,
	})

	params := &openapi.FetchMessageParams{}
	params.SetPathAccountSid(opts.sid)

	resp, err := client.ApiV2010.FetchMessage(req.Sid, params)
	if err != nil {
		return err
	}

	msg, err := TwilioMsgToMessage(resp)
	if err != nil {
		return err
	}

	rsp.Msg = msg

	return nil
}

func (e *Twilio) Verify(_ context.Context, req *pb.VerifyRequest, rsp *pb.VerifyResponse) error {
	if number, ok := req.Destination.(*pb.VerifyRequest_Number); ok {
		opts, err := GetTwilioOpts()
		if err != nil {
			return err
		}

		client := twilio.NewRestClientWithParams(twilio.RestClientParams{
			Username: opts.sid,
			Password: opts.token,
		})

		verification, err := client.VerifyV2.CreateVerification(opts.verifySid, &verify.CreateVerificationParams{
			Channel: ptr.String("sms"),
			To:      ptr.String(number.Number),
		})
		if err != nil {
			return err
		}

		rsp.Sid = ptr.ToString(verification.Sid)

		return nil
	} else if _, ok = req.Destination.(*pb.VerifyRequest_Email); ok {
		return errors.NotImplemented("twilio.Verify", "Email not implemented yet")
	} else {
		return errors.BadRequest("twilio.Verify", "No destination provided")
	}
}

func (e *Twilio) CheckVerify(_ context.Context, req *pb.CheckVerifyRequest, rsp *pb.CheckVerifyResponse) error {
	opts, err := GetTwilioOpts()
	if err != nil {
		return err
	}

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: opts.sid,
		Password: opts.token,
	})

	check, err := client.VerifyV2.CreateVerificationCheck(opts.verifySid, &verify.CreateVerificationCheckParams{
		VerificationSid: ptr.String(req.Sid),
		Code:            ptr.String(req.Code),
	})
	if err != nil {
		return err
	}

	switch ptr.ToString(check.Status) {
	case "pending":
		rsp.Status = pb.CheckVerifyResponse_PENDING
	case "approved":
		rsp.Status = pb.CheckVerifyResponse_APPROVED
	case "canceled":
		rsp.Status = pb.CheckVerifyResponse_CANCELED
	default:
		return errors.InternalServerError("twilio.CheckVerify", "Received unknown status %s", ptr.ToString(check.Status))
	}

	return nil
}

type TwilioOptions struct {
	sid           string
	token         string
	msgServiceSid string
	verifySid     string
}

func GetTwilioOpts() (TwilioOptions, error) {
	twilioConfVal, err := config.Get("twilio", config.Secret(true))
	if err != nil {
		return TwilioOptions{}, err
	}
	twilioConf := twilioConfVal.StringMap(nil)

	return TwilioOptions{
		sid:           twilioConf["sid"],
		token:         twilioConf["token"],
		msgServiceSid: twilioConf["msgSid"],
		verifySid:     twilioConf["verifySid"],
	}, nil
}

func TwilioMsgToMessage(resp *openapi.ApiV2010Message) (*pb.Message, error) {
	msg := pb.Message{}

	dateCreated, err := time.Parse(RFC2822, ptr.ToString(resp.DateCreated))
	if err != nil {
		return nil, err
	}
	dateSent, err := time.Parse(RFC2822, ptr.ToString(resp.DateSent))
	dateUpdated, err := time.Parse(RFC2822, ptr.ToString(resp.DateUpdated))
	if err != nil {
		return nil, err
	}

	msg.Body = ptr.ToString(resp.Body)
	msg.DateCreated = timestamppb.New(dateCreated)

	if dateSent.IsZero() {
		msg.DateSent = nil
	} else {
		msg.DateSent = timestamppb.New(dateSent)
	}

	msg.DateUpdated = timestamppb.New(dateUpdated)

	msg.Direction = ptr.ToString(resp.Direction)
	msg.From = ptr.ToString(resp.From)
	msg.NumMedia = ptr.ToString(resp.NumMedia)
	msg.NumSegments = ptr.ToString(resp.NumSegments)
	msg.Price = ptr.ToString(resp.Price)
	msg.PriceUnit = ptr.ToString(resp.PriceUnit)
	msg.Sid = ptr.ToString(resp.Sid)

	switch ptr.ToString(resp.Status) {
	case "accepted":
		msg.Status = pb.Message_ACCEPTED
	case "scheduled":
		msg.Status = pb.Message_SCHEDULED
	case "queued":
		msg.Status = pb.Message_QUEUED
	case "sending":
		msg.Status = pb.Message_SENDING
	case "sent":
		msg.Status = pb.Message_SENT
	case "receiving":
		msg.Status = pb.Message_RECEIVING
	case "received":
		msg.Status = pb.Message_RECEIVED
	case "delivered":
		msg.Status = pb.Message_DELIVERED
	case "undelivered":
		msg.Status = pb.Message_UNDELIVERED
	case "failed":
		msg.Status = pb.Message_FAILED
	case "read":
		msg.Status = pb.Message_READ
	case "canceled":
		msg.Status = pb.Message_CANCELED
	}

	msg.To = ptr.ToString(resp.To)
	msg.Uri = ptr.ToString(resp.Uri)

	return &msg, nil
}
