package breaker

import (
	"log"

	openfmb "github.com/corynguyenfl/openfmb-go/openfmb-ext"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/commonmodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/solarmodule"
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
	case *solarmodule.SolarControlProfile:
		return profile.SolarInverter.ConductingEquipment.MRID, nil
	// case *solarmodule.SolarDiscreteControlProfile:
	// 	return profile.SolarInverter.ConductingEquipment.MRID, nil
	case *solarmodule.SolarEventProfile:
		return profile.SolarInverter.ConductingEquipment.MRID, nil
	case *solarmodule.SolarReadingProfile:
		return profile.SolarInverter.ConductingEquipment.MRID, nil
	case *solarmodule.SolarStatusProfile:
		return profile.SolarInverter.ConductingEquipment.MRID, nil

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
	case *solarmodule.SolarControlProfile:
		return profile.ControlMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil
	// case *solarmodule.SolarDiscreteControlProfile:
	// 	return profile.ControlMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil
	case *solarmodule.SolarEventProfile:
		return profile.EventMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil
	case *solarmodule.SolarReadingProfile:
		return profile.ReadingMessageInfo.MessageInfo.IdentifiedObject.MRID.Value, nil
	case *solarmodule.SolarStatusProfile:
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
	case *solarmodule.SolarControlProfile:
		return profile.ControlMessageInfo.MessageInfo.MessageTimeStamp, nil
	// case *solarmodule.SolarDiscreteControlProfile:
	// 	return profile.ControlMessageInfo.MessageInfo.MessageTimeStamp, nil
	case *solarmodule.SolarEventProfile:
		return profile.EventMessageInfo.MessageInfo.MessageTimeStamp, nil
	case *solarmodule.SolarReadingProfile:
		return profile.ReadingMessageInfo.MessageInfo.MessageTimeStamp, nil
	case *solarmodule.SolarStatusProfile:
		return profile.StatusMessageInfo.MessageInfo.MessageTimeStamp, nil

	default:
		panic(openfmb.ErrUnknownError)
	}
}
