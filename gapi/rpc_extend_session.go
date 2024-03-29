package gapi

import (
	"context"

	db "github.com/Streamfair/streamfair_session_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_session_svc/pb/session"
	"github.com/Streamfair/streamfair_session_svc/validator"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ExtendSession(ctx context.Context, req *pb.ExtendSessionRequest) (*pb.ExtendSessionResponse, error) {
	violations := validateExtendSessionByIdRequest(req)
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

	session, err := server.store.ExtendSession(ctx, db.ExtendSessionParams{
		Uuid:      uuid,
		ExpiresAt: req.GetExpiresAt().AsTime(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to extend session: %v", err)
	}

	rsp := &pb.ExtendSessionResponse{
		Session: convertSession(session),
	}

	return rsp, nil
}

func validateExtendSessionByIdRequest(req *pb.ExtendSessionRequest) (violations []*CustomError) {
	if err := validator.ValidateUuid(req.GetUuid()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("uuid", err))
	}

	if err := validator.ValidateExpiration(req.GetExpiresAt().AsTime()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("expires_at", err))
	}
	return violations
}
