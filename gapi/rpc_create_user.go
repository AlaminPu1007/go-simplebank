package gapi

import (
	"context"

	db "github.com/alaminpu1007/simplebank/db/sqlc"
	"github.com/alaminpu1007/simplebank/pb"
	"github.com/alaminpu1007/simplebank/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashPassword, err := util.HashedPassword(req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// log.Println((pqErr.Code.Name()))
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user name already exists %s", err)
			}
		}

		return nil, status.Errorf(codes.Internal, "Failed to create user %s", err)
	}

	// create a return response object and return to the user
	res := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return res, nil
}
