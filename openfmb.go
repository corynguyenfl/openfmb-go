package openfmb

import (
	"errors"
	"log"
	"reflect"
	"strings"

	"github.com/nats-io/nats.go"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/breakermodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/capbankmodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/circuitsegmentservicemodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/essmodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/generationmodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/loadmodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/metermodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/reclosermodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/regulatormodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/reservemodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/resourcemodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/solarmodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/switchmodule"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Topic struct {
	Module  string
	Profile string
	Mrid    string
}

func TopicFromString(subject string, args ...string) (topic *Topic, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Failed to parse OpenFMB message topic: " + subject)
		}
	}()

	sep := "."
	if len(args) > 0 {
		sep = args[0]
	}

	tokens := strings.Split(subject, sep)

	topic = &Topic{}

	if len(tokens) > 1 {
		topic.Module = tokens[1]
	}
	if len(tokens) > 2 {
		topic.Profile = tokens[2]
	}
	if len(tokens) > 3 {
		topic.Mrid = tokens[3]
	}
	return topic, nil
}

func (topic *Topic) String() string {
	var sb strings.Builder

	sb.WriteString("openfmb")
	if len(topic.Module) > 0 {
		sb.WriteString(".")
		sb.WriteString(topic.Module)

		if len(topic.Profile) > 0 {
			sb.WriteString(".")
			sb.WriteString(topic.Profile)

			if len(topic.Mrid) > 0 {
				sb.WriteString(".")
				sb.WriteString(topic.Mrid)
			} else {
				sb.WriteString(".>")
			}
		} else {
			sb.WriteString(".>")
		}
	} else {
		sb.WriteString(".>")
	}

	return sb.String()
}

func GetSubjectPrefix(profile proto.Message) string {
	return strings.Replace(reflect.TypeOf(profile).String(), "*", "openfmb.", 1)
}

type OpenFMBMessage struct {
	Topic *Topic
	Msg   proto.Message
}

func (m *OpenFMBMessage) ToJson() string {
	return protojson.Format(m.Msg)
}

func FromNatsMessage(natsMsg *nats.Msg) (openfmbMsg *OpenFMBMessage, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("unable to decode message into openfmb")
		}
	}()

	topic, topicErr := TopicFromString(natsMsg.Subject)

	if topicErr != nil {
		return nil, topicErr
	}

	log.Printf("Topic: %v\n", topic.String())

	switch topic.Profile {
	case "BreakerDiscreteControlProfile":
		var profile breakermodule.BreakerDiscreteControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "BreakerEventProfile":
		var profile breakermodule.BreakerEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "BreakerReadingProfile":
		var profile breakermodule.BreakerReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "BreakerStatusProfile":
		var profile breakermodule.BreakerStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "CapBankControlProfile":
		var profile capbankmodule.CapBankControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "CapBankDiscreteControlProfile":
		var profile capbankmodule.CapBankDiscreteControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "CapBankEventProfile":
		var profile capbankmodule.CapBankEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "CapBankReadingProfile":
		var profile capbankmodule.CapBankReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "CapBankStatusProfile":
		var profile capbankmodule.CapBankStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "CircuitSegmentControlProfile":
		var profile circuitsegmentservicemodule.CircuitSegmentControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "CircuitSegmentEventProfile":
		var profile circuitsegmentservicemodule.CircuitSegmentEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "CircuitSegmentStatusProfile":
		var profile circuitsegmentservicemodule.CircuitSegmentStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	// case "ESSCapabilityOverrideProfile":
	// 	var profile essmodule.ESSCapabilityOverrideProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	// case "ESSCapabilityProfile":
	// 	var profile essmodule.ESSCapabilityProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	case "ESSControlProfile":
		var profile essmodule.ESSControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	// case "ESSDiscreteControlProfile":
	// 	var profile essmodule.ESSDiscreteControlProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	case "ESSEventProfile":
		var profile essmodule.ESSEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "ESSReadingProfile":
		var profile essmodule.ESSReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "ESSStatusProfile":
		var profile essmodule.ESSStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	// case "GenerationCapabilityOverrideProfile":
	// 	var profile generationmodule.GenerationCapabilityOverrideProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	// case "GenerationCapabilityProfile":
	// 	var profile generationmodule.GenerationCapabilityProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	case "GenerationControlProfile":
		var profile generationmodule.GenerationControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "GenerationDiscreteControlProfile":
		var profile generationmodule.GenerationDiscreteControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "GenerationReadingProfile":
		var profile generationmodule.GenerationReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "GenerationEventProfile":
		var profile generationmodule.GenerationEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "GenerationStatusProfile":
		var profile generationmodule.GenerationStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	// case "InterconnectionPlannedScheduleProfile":
	// 	var profile interconnectionmodule.InterconnectionPlannedScheduleProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	// case "InterconnectionRequestedScheduleProfile":
	// 	var profile interconnectionmodule.InterconnectionRequestedScheduleProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	case "LoadControlProfile":
		var profile loadmodule.LoadControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "LoadEventProfile":
		var profile loadmodule.LoadEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "LoadReadingProfile":
		var profile loadmodule.LoadReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "LoadStatusProfile":
		var profile loadmodule.LoadStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "MeterReadingProfile":
		var profile metermodule.MeterReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "RecloserDiscreteControlProfile":
		var profile reclosermodule.RecloserDiscreteControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "RecloserEventProfile":
		var profile reclosermodule.RecloserEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "RecloserReadingProfile":
		var profile reclosermodule.RecloserReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "RecloserStatusProfile":
		var profile reclosermodule.RecloserStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "RegulatorControlProfile":
		var profile regulatormodule.RegulatorControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "RegulatorDiscreteControlProfile":
		var profile regulatormodule.RegulatorDiscreteControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "RegulatorEventProfile":
		var profile regulatormodule.RegulatorEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "RegulatorReadingProfile":
		var profile regulatormodule.RegulatorReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "RegulatorStatusProfile":
		var profile regulatormodule.RegulatorStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "ReserveAvailabilityProfile":
		var profile reservemodule.ReserveAvailabilityProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "ReserveRequestProfile":
		var profile reservemodule.ReserveRequestProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "ResourceDiscreteControlProfile":
		var profile resourcemodule.ResourceDiscreteControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "ResourceReadingProfile":
		var profile resourcemodule.ResourceReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "ResourceEventProfile":
		var profile resourcemodule.ResourceEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "ResourceStatusProfile":
		var profile resourcemodule.ResourceStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	// case "SolarCapabilityOverrideProfile":
	// 	var profile solarmodule.SolarCapabilityOverrideProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	// case "SolarCapabilityProfile":
	// 	var profile solarmodule.SolarCapabilityProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	case "SolarControlProfile":
		var profile solarmodule.SolarControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	// case "SolarDiscreteControlProfile":
	// 	var profile solarmodule.SolarDiscreteControlProfile
	// 	err = proto.Unmarshal(natsMsg.Data, &profile)
	// 	if err == nil {
	// 		return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
	// 	}
	case "SolarEventProfile":
		var profile solarmodule.SolarEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "SolarReadingProfile":
		var profile solarmodule.SolarReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "SolarStatusProfile":
		var profile solarmodule.SolarStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "SwitchDiscreteControlProfile":
		var profile switchmodule.SwitchDiscreteControlProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "SwitchEventProfile":
		var profile switchmodule.SwitchEventProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "SwitchReadingProfile":
		var profile switchmodule.SwitchReadingProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	case "SwitchStatusProfile":
		var profile switchmodule.SwitchStatusProfile
		err = proto.Unmarshal(natsMsg.Data, &profile)
		if err == nil {
			return &OpenFMBMessage{Topic: topic, Msg: &profile}, nil
		}
	}

	return nil, errors.New("invalid openfmb profile: " + topic.Profile)
}
