package handler

import (
	"context"
	"fmt"
	"github.com/ansg191/northstars-backend/touchpad/utils"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/store"

	log "github.com/micro/micro/v3/service/logger"

	touchpad "github.com/ansg191/northstars-backend/touchpad/proto"
)

type Touchpad struct{}

// GetMeetInfo is a single request handler called via client.Call or the generated client code
func (e *Touchpad) GetMeetInfo(ctx context.Context, req *touchpad.GetMeetInfoRequest, rsp *touchpad.GetMeetInfoResponse) error {
	log.Info("Received Touchpad.GetMeetInfo request")

	if req.MeetID == 0 {
		return errors.BadRequest("touchpad.GetMeetInfo.MissingMeetID", "MeetID Missing")
	}

	apiInfo, err := utils.GetMeetInfo(int(req.MeetID))
	if err != nil {
		return err
	}

	*rsp, err = apiInfo.ToProto()
	if err != nil {
		return err
	}

	return nil
}

func (e *Touchpad) GetMeetEvents(ctx context.Context, req *touchpad.GetMeetEventsRequest, rsp *touchpad.GetMeetEventsResponse) error {
	log.Info("Received Touchpad.GetMeetEvents request")

	if req.MeetID == 0 {
		return errors.BadRequest("touchpad.GetMeetInfo.MissingMeetID", "MeetID Missing")
	}

	apiEvents, err := utils.GetMeetEvents(int(req.MeetID))
	if err != nil {
		return err
	}

	for _, event := range apiEvents {
		protoEvent, err := event.ToProto()
		if err != nil {
			return err
		}

		rsp.Events = append(rsp.Events, &protoEvent)
	}

	return nil
}

func (e *Touchpad) GetIEvent(ctx context.Context, req *touchpad.GetIEventRequest, rsp *touchpad.GetIEventResponse) error {
	log.Info("Received Touchpad.GetIEvent request")

	if req.MeetID == 0 {
		return errors.BadRequest("touchpad.GetIEvent.MissingMeetID", "MeetID Missing")
	}
	if req.EventID == 0 {
		return errors.BadRequest("touchpad.GetIEvent.MissingEventID", "EventID Missing")
	}

	apiTimings, err := utils.GetIndividualEvent(int(req.MeetID), int(req.EventID))
	if err != nil {
		return err
	}

	for _, event := range apiTimings {
		protoTiming, err := event.ToProto()
		if err != nil {
			return err
		}

		rsp.Timings = append(rsp.Timings, &protoTiming)
	}

	return nil
}

func (e *Touchpad) CheckMeetDiff(ctx context.Context, req *touchpad.CheckMeetDiffRequest, rsp *touchpad.CheckMeetDiffResponse) error {
	log.Info("Received Touchpad.CheckMeetDiff request")

	if req.MeetID == 0 {
		return errors.BadRequest("touchpad.CheckMeetDiff.MissingMeetID", "MeetID Missing")
	}

	apiEvents, err := utils.GetMeetEvents(int(req.MeetID))
	if err != nil {
		return err
	}

	list, err := store.Read(fmt.Sprintf("%d", req.MeetID), store.ReadPrefix())
	if err != nil {
		return err
	}
	if len(list) == 0 {
		// Meet not initialized
		for _, event := range apiEvents {
			protoEvent, err := event.ToProto()
			if err != nil {
				return err
			}

			err = store.Write(store.NewRecord(fmt.Sprintf("%d/%d", req.MeetID, protoEvent.Id), protoEvent))
			if err != nil {
				return err
			}
		}
		rsp.Initial = true
		rsp.DiffExists = false

		return nil
	}

	for _, event := range apiEvents {
		protoEvent, err := event.ToProto()
		if err != nil {
			return err
		}

		record, err := store.Read(fmt.Sprintf("%d/%d", req.MeetID, protoEvent.Id))
		if err != nil {
			return err
		}
		if len(record) == 0 {
			err = store.Write(store.NewRecord(fmt.Sprintf("%d/%d", req.MeetID, protoEvent.Id), protoEvent))
			if err != nil {
				return err
			}
			continue
		}

		var storeEvent touchpad.Event
		err = record[0].Decode(&storeEvent)
		if err != nil {
			return err
		}

		if storeEvent.Status != protoEvent.Status {
			// Event Status Changed
			rsp.DiffExists = true

			log.Infof("%d transitioned from %v to %v", protoEvent.EventNumber, storeEvent.Status, protoEvent.Status)

			diff := touchpad.CheckMeetDiffResponse_Diff{
				Old: &storeEvent,
				New: &protoEvent,
			}
			rsp.Diffs = append(rsp.Diffs, &diff)

			err = store.Write(store.NewRecord(fmt.Sprintf("%d/%d", req.MeetID, protoEvent.Id), protoEvent))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
