package users

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Archisman-Mridha/go-microservice-template/internal/adapters/repositories/users/generated"
	"github.com/Archisman-Mridha/go-microservice-template/internal/constants"
	coreTypes "github.com/Archisman-Mridha/go-microservice-template/internal/core/types"
	"github.com/Archisman-Mridha/go-microservice-template/pkg/connectors"
	"github.com/Archisman-Mridha/go-microservice-template/pkg/types"
	"github.com/Archisman-Mridha/go-microservice-template/pkg/utils"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type UsersRepositoryAdapter struct {
	*connectors.PostgresConnector
	queries *generated.Queries
}

func NewUsersRepositoryAdapter(ctx context.Context,
	args *connectors.NewPostgresConnectorArgs,
) *UsersRepositoryAdapter {
	postgresConnector := connectors.NewPostgresConnector(ctx, args)

	queries := generated.New(postgresConnector.GetConnection())

	return &UsersRepositoryAdapter{
		postgresConnector,
		queries,
	}
}

func (u *UsersRepositoryAdapter) Create(ctx context.Context,
	args *coreTypes.CreateUserArgs,
) (types.ID, error) {
	userID, err := u.queries.CreateUser(ctx, generated.CreateUserParams{
		Name:     args.Name,
		Email:    args.Email,
		Username: args.Username,
		Password: args.HashedPassword,
	})
	if err != nil {
		pgErr := err.(*pgconn.PgError)
		if pgErr.Code == pgerrcode.UniqueViolation {
			switch pgErr.ColumnName {
			case "email":
				return 0, constants.ErrDuplicateEmail

			case "username":
				return 0, constants.ErrDuplicateUsername
			}
		}

		return 0, utils.WrapError(err)
	}
	return userID, nil
}

func (u *UsersRepositoryAdapter) FindByEmail(ctx context.Context,
	email string,
) (*coreTypes.FindUserByOperationOutput, error) {
	userDetails, err := u.queries.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, constants.ErrUserNotFound
		}

		return nil, utils.WrapError(err)
	}

	output := &coreTypes.FindUserByOperationOutput{
		ID:             userDetails.ID,
		HashedPassword: userDetails.Password,
	}
	return output, nil
}

func (u *UsersRepositoryAdapter) FindByUsername(ctx context.Context,
	username string,
) (*coreTypes.FindUserByOperationOutput, error) {
	userDetails, err := u.queries.FindUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, constants.ErrUserNotFound
		}

		return nil, utils.WrapError(err)
	}

	output := &coreTypes.FindUserByOperationOutput{
		ID:             userDetails.ID,
		HashedPassword: userDetails.Password,
	}
	return output, nil
}

func (u *UsersRepositoryAdapter) UserIDExists(ctx context.Context, id types.ID) (bool, error) {
	_, err := u.queries.FindUserByID(ctx, id)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return false, constants.ErrUserNotFound
		}

		return false, utils.WrapError(err)
	}
	return true, nil
}
