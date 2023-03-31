package breaker

import (
	"log"

	openfmb "github.com/corynguyenfl/openfmb-go/openfmb-ext"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/commonmodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/switchmodule"
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
	case *switchmodule.SwitchDiscreteControlProfile:
		return profile.ProtectedSwitch.ConductingEquipment.MRID, nil
	case *switchmodule.SwitchEventProfile:
		return profile.ProtectedSwitch.ConductingEquipment.MRID, nil
	case *switchmodule.SwitchReadingProfile:
		return profile.ProtectedSwitch.ConductingEquipment.MRID, nil
	case *switchmodule.SwitchStatusProfile:
		return profile.ProtectedSwitch.ConductingEquipment.MRID, nil

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
	case *switchmodule.SwitchDiscreteControlProfile:
		return profile.ControlMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil
	case *switchmodule.SwitchEventProfile:
		return profile.EventMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil
	case *switchmodule.SwitchReadingProfile:
		return profile.ReadingMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil
	case *switchmodule.SwitchStatusProfile:
		return profile.StatusMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil

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
	case *switchmodule.SwitchDiscreteControlProfile:
		return profile.ControlMessageInfo.MessageInfo.MessageTimeStamp, nil
	case *switchmodule.SwitchEventProfile:
		return profile.EventMessageInfo.MessageInfo.MessageTimeStamp, nil
	case *switchmodule.SwitchReadingProfile:
		return profile.ReadingMessageInfo.MessageInfo.MessageTimeStamp, nil
	case *switchmodule.SwitchStatusProfile:
		return profile.StatusMessageInfo.MessageInfo.MessageTimeStamp, nil

	default:
		panic(openfmb.ErrUnknownError)
	}
}
