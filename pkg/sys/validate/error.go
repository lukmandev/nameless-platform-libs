package validate

import (
	"encoding/json"

	"github.com/lukmandev/nameless-platform-libs/pkg/sys/codes"
	"github.com/pkg/errors"
	grpcCodes "google.golang.org/grpc/codes"
)

type ValidationErrors struct {
	Messages []string `json:"error_messages"`
}

func (v *ValidationErrors) addError(message string) {
	v.Messages = append(v.Messages, message)
}

func NewValidationErrors(messages ...string) *ValidationErrors {
	return &ValidationErrors{
		Messages: messages,
	}
}

func (v *ValidationErrors) Error() string {
	data, err := json.Marshal(v.Messages)
	if err != nil {
		return err.Error()
	}

	return string(data)
}

func IsValidationError(err error) bool {
	var ve *ValidationErrors
	return errors.As(err, &ve)
}

func ToGRPCCode(code codes.Code) grpcCodes.Code {
	var res grpcCodes.Code

	switch code {
	case codes.OK:
		res = grpcCodes.OK
	case codes.Canceled:
		res = grpcCodes.Canceled
	case codes.InvalidArgument:
		res = grpcCodes.InvalidArgument
	case codes.DeadlineExceeded:
		res = grpcCodes.DeadlineExceeded
	case codes.NotFound:
		res = grpcCodes.NotFound
	case codes.AlreadyExists:
		res = grpcCodes.AlreadyExists
	case codes.PermissionDenied:
		res = grpcCodes.PermissionDenied
	case codes.ResourceExhausted:
		res = grpcCodes.ResourceExhausted
	case codes.FailedPrecondition:
		res = grpcCodes.FailedPrecondition
	case codes.Aborted:
		res = grpcCodes.Aborted
	case codes.OutOfRange:
		res = grpcCodes.OutOfRange
	case codes.Unimplemented:
		res = grpcCodes.Unimplemented
	case codes.Internal:
		res = grpcCodes.Internal
	case codes.Unavailable:
		res = grpcCodes.Unavailable
	case codes.DataLoss:
		res = grpcCodes.DataLoss
	case codes.Unauthenticated:
		res = grpcCodes.Unauthenticated
	default:
		res = grpcCodes.Unknown
	}

	return res
}
