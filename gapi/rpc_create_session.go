package gapi

import (
	"context"
	"errors"

	db "github.com/Streamfair/streamfair_session_svc/db/sqlc"
	pb "github.com/Streamfair/common_proto/SessionService/pb/session"
	"github.com/Streamfair/streamfair_session_svc/validator"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateSession(ctx context.Context, req *pb.CreateSessionRequest) (*pb.CreateSessionResponse, error) {
	violations := validateCreateSessionRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	_, err := server.localTokenMaker.VerifyLocalToken(req.RefreshToken)
	if err != nil {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("refresh_token", errors.New("refresh_token is invalid"))
		return nil, invalidArgumentError(violation)
	}

	uuid, err := uuid.Parse(req.GetUuid())
	if err != nil {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("UUID", err)
		return nil, invalidArgumentError(violation)
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		Uuid:         uuid,
		Username:     req.GetUsername(),
		RefreshToken: req.GetRefreshToken(),
		UserAgent:    req.GetUserAgent(),
		ClientIp:     req.GetClientIp(),
		IsBlocked:    req.GetIsBlocked(),
		ExpiresAt:    req.GetExpiresAt().AsTime(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session: %v", err)
	}

	rsp := &pb.CreateSessionResponse{
		Session: convertSession(session),
	}

	return rsp, nil
}

func validateCreateSessionRequest(req *pb.CreateSessionRequest) (violations []*CustomError) {
	if err := validator.ValidateUuid(req.GetUuid()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("uuid", err))
	}

	if err := validator.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("username", err))
	}

	if err := validator.ValidateUserAgent(req.GetUserAgent()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("user_agent", err))
	}

	if err := validator.ValidateClientIp(req.GetClientIp()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("client_ip", err))
	}

	if err := validator.ValidateExpiration(req.GetExpiresAt().AsTime()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("expires_at", err))
	}
	return violations
}
