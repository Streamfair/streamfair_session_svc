package gapi

import (
	"context"

	pb "github.com/Streamfair/streamfair_session_svc/pb/session"
	"github.com/Streamfair/streamfair_session_svc/validator"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeleteSession(ctx context.Context, req *pb.DeleteSessionRequest) (*emptypb.Empty, error) {
	violations := validateDeleteSessiondRequest(req)
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

	err = server.store.DeleteSession(ctx, uuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete session: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteSessiondRequest(req *pb.DeleteSessionRequest) (violations []*CustomError) {
	if err := validator.ValidateUuid(req.GetUuid()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("uuid", err))
	}

	return violations
}
