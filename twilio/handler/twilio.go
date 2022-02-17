package handler

import (
	"context"
	"github.com/micro/micro/v3/service/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
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

type TwilioOptions struct {
	sid           string
	token         string
	msgServiceSid string
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
	}, nil
}

func ToString(ptr *string) (s string) {
	if ptr == nil {
		return s
	}
	return *ptr
}

func TwilioMsgToMessage(resp *openapi.ApiV2010Message) (*pb.Message, error) {
	msg := pb.Message{}

	dateCreated, err := time.Parse(RFC2822, ToString(resp.DateCreated))
	if err != nil {
		return nil, err
	}
	dateSent, err := time.Parse(RFC2822, ToString(resp.DateSent))
	dateUpdated, err := time.Parse(RFC2822, ToString(resp.DateUpdated))
	if err != nil {
		return nil, err
	}

	msg.Body = ToString(resp.Body)
	msg.DateCreated = timestamppb.New(dateCreated)

	if dateSent.IsZero() {
		msg.DateSent = nil
	} else {
		msg.DateSent = timestamppb.New(dateSent)
	}

	msg.DateUpdated = timestamppb.New(dateUpdated)

	msg.Direction = ToString(resp.Direction)
	msg.From = ToString(resp.From)
	msg.NumMedia = ToString(resp.NumMedia)
	msg.NumSegments = ToString(resp.NumSegments)
	msg.Price = ToString(resp.Price)
	msg.PriceUnit = ToString(resp.PriceUnit)
	msg.Sid = ToString(resp.Sid)

	switch ToString(resp.Status) {
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

	msg.To = ToString(resp.To)
	msg.Uri = ToString(resp.Uri)

	return &msg, nil
}
