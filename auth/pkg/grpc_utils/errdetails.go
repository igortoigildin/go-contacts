package grpcutils

import (
	"buf.build/go/protovalidate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func convertProtovalidateValidationErrorToErrdetailsBadRequest(valErr *protovalidate.ValidationError) *errdetails.BadRequest {
	return &errdetails.BadRequest{
		FieldViolations: protovalidateVialationsToGoogleViolations(valErr.Violations),
	}
}

func protovalidateVialationsToGoogleViolations(vs []*protovalidate.Violation) []*errdetails.BadRequest_FieldViolation {
	res := make([]*errdetails.BadRequest_FieldViolation, len(vs))
	for i, v := range vs {
		res[i] = &errdetails.BadRequest_FieldViolation{
			Field:       v.FieldValue.String(),
			Description: v.Proto.GetMessage(),
		}
		
	}

	return res
}