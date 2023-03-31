package ts

import (
	"time"

	"gitlab.com/openfmb/psm/ops/protobuf/go-openfmb-ops-protobuf/v2/openfmb/commonmodule"
)

func Now() *commonmodule.Timestamp {
	nano := time.Now().UTC().UnixNano()
	seconds := float64(nano) / 1e9
	return &commonmodule.Timestamp{
		Seconds:     uint64(seconds),
		Nanoseconds: 0,
	}
}
