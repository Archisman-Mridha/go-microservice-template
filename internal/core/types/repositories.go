package types

import (
	"context"

	"github.com/Archisman-Mridha/go-microservice-template/pkg/types"
)

type (
	UsersRepository interface {
		Create(ctx context.Context, args *CreateUserArgs) (types.ID, error)

		FindByEmail(ctx context.Context, email string) (*FindUserByOperationOutput, error)
		FindByUsername(ctx context.Context, username string) (*FindUserByOperationOutput, error)
		UserIDExists(ctx context.Context, id types.ID) (bool, error)
	}

	CreateUserArgs struct {
		Name,
		Email,
		Username,
		HashedPassword string
	}

	FindUserByOperationOutput struct {
		ID             types.ID
		HashedPassword string
	}
)
