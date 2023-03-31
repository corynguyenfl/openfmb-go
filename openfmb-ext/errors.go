package openfmb_ext

import "errors"

type ErrorCode int32

var (
	ErrNoDeviceMRID       = errors.New("openfmb: no device mRID")
	ErrNoDeviceName       = errors.New("openfmb: no device name")
	ErrNoMessageMRID      = errors.New("openfmb: no message mRID")
	ErrNoMessageTimestamp = errors.New("openfmb: no message timestamp")
	ErrMissingData        = errors.New("openfmb: missing data")
	ErrInvalidProfile     = errors.New("openfmb: invalid OpenFMB profile")
	ErrUnknownError       = errors.New("openfmb: unknown error")
)
