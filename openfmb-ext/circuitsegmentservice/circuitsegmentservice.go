package breaker

import (
	"log"

	openfmb "github.com/corynguyenfl/openfmb-go/openfmb-ext"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/circuitsegmentservicemodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/commonmodule"
	"google.golang.org/protobuf/proto"
)

func DeviceMRID(msg proto.Message) (mrid string, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:: Failed to extract MRID from profile: ", msg)
			err = openfmb.ErrNoDeviceMRID
		}
	}()

	switch profile := msg.(type) {
	case *circuitsegmentservicemodule.CircuitSegmentControlProfile:
		return profile.ApplicationSystem.MRID, nil
	case *circuitsegmentservicemodule.CircuitSegmentEventProfile:
		return profile.ApplicationSystem.MRID, nil
	case *circuitsegmentservicemodule.CircuitSegmentStatusProfile:
		return profile.ApplicationSystem.MRID, nil

	default:
		panic(openfmb.ErrUnknownError)
	}
}

func MessageMRID(msg proto.Message) (mrid string, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:: Failed to extract message MRID from profile: ", msg)
			err = openfmb.ErrNoMessageMRID
		}
	}()
	switch profile := msg.(type) {
	case *circuitsegmentservicemodule.CircuitSegmentControlProfile:
		return profile.ControlMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil
	case *circuitsegmentservicemodule.CircuitSegmentEventProfile:
		return profile.EventMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil
	case *circuitsegmentservicemodule.CircuitSegmentStatusProfile:
		return profile.EventMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil // BUG

	default:
		panic(openfmb.ErrUnknownError)
	}
}

func MessageTimeStamp(msg proto.Message) (ts *commonmodule.Timestamp, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:: Failed to extract message timestamp from profile: ", msg)
			err = openfmb.ErrNoMessageTimestamp
		}
	}()

	switch profile := msg.(type) {
	case *circuitsegmentservicemodule.CircuitSegmentControlProfile:
		return profile.ControlMessageInfo.MessageInfo.MessageTimeStamp, nil
	case *circuitsegmentservicemodule.CircuitSegmentEventProfile:
		return profile.EventMessageInfo.MessageInfo.MessageTimeStamp, nil
	case *circuitsegmentservicemodule.CircuitSegmentStatusProfile:
		return profile.EventMessageInfo.MessageInfo.MessageTimeStamp, nil

	default:
		panic(openfmb.ErrUnknownError)
	}
}
