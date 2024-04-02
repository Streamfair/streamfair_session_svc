package gapi

import (
	"context"
	"time"

	pb "github.com/Streamfair/streamfair_session_svc/common_proto/SessionService/pb/session"
	"github.com/Streamfair/streamfair_session_svc/validator"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifySession(ctx context.Context, req *pb.VerifySessionRequest) (*pb.VerifySessionResponse, error) {
	violations := validateVerifySessionRequest(req)
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

	if session.IsBlocked {
		return nil, status.Errorf(codes.PermissionDenied, "session is blocked")
	}
	if session.ExpiresAt.Before(time.Now()) {
		return nil, status.Errorf(codes.PermissionDenied, "session is expired")
	}

	rsp := &pb.VerifySessionResponse{
		Session: convertSession(session),
	}

	return rsp, nil
}

func validateVerifySessionRequest(req *pb.VerifySessionRequest) (violations []*CustomError) {
	if err := validator.ValidateUuid(req.GetUuid()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("uuid", err))
	}

	return violations
}
