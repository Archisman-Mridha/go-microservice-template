package usecases

import (
	"context"

	"github.com/Archisman-Mridha/go-microservice-template/pkg/types"
)

func (u *Usecases) GetUserIDFromJWT(ctx context.Context, jwt string) (*types.ID, error) {
	userID, err := u.tokenService.GetUserIDFromToken(jwt)
	if err != nil {
		return nil, err
	}

	// Verify that the user exists in the database.
	_, err = u.usersRepository.UserIDExists(ctx, *userID)
	if err != nil {
		return nil, err
	}

	return userID, nil
}
