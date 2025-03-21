package constants

import (
	"github.com/Archisman-Mridha/go-microservice-template/pkg/utils"
)

var (
	ErrInvalidEmail    = utils.NewAPIError("invalid email")
	ErrInvalidUsername = utils.NewAPIError("invalid username")

	ErrDuplicateEmail    = utils.NewAPIError("email already exists")
	ErrDuplicateUsername = utils.NewAPIError("username already exists")

	ErrInvalidJWT = utils.NewAPIError("invalid JWT")
	ErrExpiredJWT = utils.NewAPIError("expired JWT")

	ErrUserNotFound = utils.NewAPIError("user not found")
)
