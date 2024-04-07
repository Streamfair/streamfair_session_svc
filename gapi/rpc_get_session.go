package gapi

import (
	"context"

	pb "github.com/Streamfair/common_proto/SessionService/pb/session"
	"github.com/Streamfair/streamfair_session_svc/validator"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetSession(ctx context.Context, req *pb.GetSessionRequest) (*pb.GetSessionResponse, error) {
	violations := validateGetSessionRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	uuid, err := uuid.Parse(req.GetUuid())
	if err != nil {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("UUID", err)
		return nil, invalidArgumentError(violation)
	}

	session, err := server.store.GetSession(ctx, uuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get session: %v", err)
	}

	rsp := &pb.GetSessionResponse{
		Session: convertSession(session),
	}

	return rsp, nil
}

func validateGetSessionRequest(req *pb.GetSessionRequest) (violations []*CustomError) {
	if err := validator.ValidateUuid(req.GetUuid()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("uuid", err))
	}

	return violations
}
