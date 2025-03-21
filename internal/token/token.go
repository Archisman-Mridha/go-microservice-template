package token

import (
	"github.com/Archisman-Mridha/go-microservice-template/pkg/types"
)

type (
	Token = string

	TokenService interface {
		Issue(userID types.ID) (*Token, error)
		GetUserIDFromToken(token Token) (*types.ID, error)
	}
)
