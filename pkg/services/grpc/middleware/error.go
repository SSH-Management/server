package middleware

import (
	"context"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/SSH-Management/server/pkg/log"
)


type (
	ErrorResponse struct {
		Message string `json:"message,omitempty"`
	}

	ValidationErrorResponse struct {
		Message string              `json:"message,omitempty"`
		Errors  map[string][]string `json:"errors,omitempty"`
	}
)

func handleValidationError(ut ut.Translator, validations validator.ValidationErrors) (err error) {
	st := status.New(codes.InvalidArgument, "invalid data")
	if validations == nil {
		return status.Error(codes.InvalidArgument, "invalid data")
	}

	l := len(validations)

	fields := make([]*errdetails.BadRequest_FieldViolation, 0, l)

	for key, validation := range validations.Translate(ut) {
		fields = append(fields, &errdetails.BadRequest_FieldViolation{
			Field:       key,
			Description: validation,
		})
	}

	badRequest := &errdetails.BadRequest{
		FieldViolations: fields,
	}

	st, err = st.WithDetails(badRequest)

	if err != nil {
		return status.Error(codes.Internal, "internal error")
	}

	return st.Err()
}

func handleError(ut ut.Translator, err error) error {
	if validations, ok := err.(validator.ValidationErrors); ok {
		return handleValidationError(ut, validations)
	}

	return status.Error(codes.Internal, "Something went wrong during the request")
}

func UnaryErrorHandler(ut ut.Translator, errorLogger *log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		res, err := handler(ctx, req)
		if err != nil {
			errorLogger.Err(err).
				Interface("request", req).
				Str("method", info.FullMethod).
				Msg("Something went wrong")

			return nil, handleError(ut, err)
		}

		return res, nil
	}
}
