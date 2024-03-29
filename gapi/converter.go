package gapi

import (
	db "github.com/Streamfair/streamfair_session_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_session_svc/pb/session"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertSession(session db.SessionSvcSession) *pb.Session {
	return &pb.Session{
		Uuid:         session.Uuid.String(),
		Username:     session.Username,
		RefreshToken: session.RefreshToken,
		UserAgent:    session.UserAgent,
		ClientIp:     session.ClientIp,
		IsBlocked:    session.IsBlocked,
		ExpiresAt:    timestamppb.New(session.ExpiresAt),
		CreatedAt:    timestamppb.New(session.CreatedAt),
	}
}
