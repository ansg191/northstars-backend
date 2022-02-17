package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	twilio "twilio/proto"
)

type Twilio struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Twilio) Call(ctx context.Context, req *twilio.Request, rsp *twilio.Response) error {
	log.Info("Received Twilio.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Twilio) Stream(ctx context.Context, req *twilio.StreamingRequest, stream twilio.Twilio_StreamStream) error {
	log.Infof("Received Twilio.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&twilio.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Twilio) PingPong(ctx context.Context, stream twilio.Twilio_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&twilio.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
