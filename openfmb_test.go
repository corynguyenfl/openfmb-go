package openfmb

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/corynguyenfl/openfmb-go/openfmb-ext/breaker"

	"github.com/corynguyenfl/openfmb-go/utils/ts"
	"github.com/google/uuid"
	"github.com/nats-io/nats-server/v2/server"
	natsserver "github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/breakermodule"
	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/commonmodule"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func RunServer() *server.Server {
	opts := natsserver.DefaultTestOptions
	opts.Cluster.Name = "testing"
	return natsserver.RunServer(&opts)
}

func TestFromNatsMessage(t *testing.T) {
	s := RunServer()
	defer s.Shutdown()

	subject := "openfmb.breakermodule.BreakerStatusProfile.e6768784-48ad-40e9-af2a-9676413d4d6a"
	topic, _ := TopicFromString(subject)

	profile := createBreakerStatus("e6768784-48ad-40e9-af2a-9676413d4d6a")

	data, _ := proto.Marshal(profile)

	nc, _ := nats.Connect(nats.DefaultURL)

	ch := make(chan error, 1)
	nc.Subscribe(subject, func(m *nats.Msg) {
		openfmbMessage, _ := FromNatsMessage(m)

		mrid, _ := breaker.DeviceMRID(openfmbMessage.Msg)

		if mrid != "e6768784-48ad-40e9-af2a-9676413d4d6a" {
			t.Errorf("Expected MRID %s doesn't match actual %s", "e6768784-48ad-40e9-af2a-9676413d4d6a", mrid)
		}

		var err error
		if !reflect.DeepEqual(profile.ProtoReflect(), openfmbMessage.Msg.ProtoReflect()) {
			err = errors.New("Did not receive the correct openfmb message")
		}
		if !reflect.DeepEqual(topic, openfmbMessage.Topic) {
			err = errors.New("Did not receive the correct topic in openfmb message")
		}
		ch <- err
	})

	nc.Publish(subject, data)

	select {
	case e := <-ch:
		if e != nil {
			t.Fatal(e.Error())
		}
	case <-time.After(time.Second):
		t.Fatal("Failed to receive message")
	}
}

func createBreakerStatus(mrid string) *breakermodule.BreakerStatusProfile {
	return &breakermodule.BreakerStatusProfile{
		StatusMessageInfo: &commonmodule.StatusMessageInfo{
			MessageInfo: &commonmodule.MessageInfo{
				IdentifiedObject: &commonmodule.IdentifiedObject{
					MRID: &wrapperspb.StringValue{
						Value: uuid.New().String(),
					},
				},
				MessageTimeStamp: ts.Now(),
			},
		},
		Breaker: &breakermodule.Breaker{
			ConductingEquipment: &commonmodule.ConductingEquipment{
				MRID: mrid,
			},
		},
		BreakerStatus: &breakermodule.BreakerStatus{
			StatusAndEventXCBR: &commonmodule.StatusAndEventXCBR{
				Pos: &commonmodule.PhaseDPS{
					Phs3: &commonmodule.StatusDPS{
						StVal: commonmodule.DbPosKind_DbPosKind_closed,
					},
				},
			},
		},
	}
}
