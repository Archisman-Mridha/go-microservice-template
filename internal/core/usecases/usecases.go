package usecases

import (
	coreTypes "github.com/Archisman-Mridha/go-microservice-template/internal/core/types"
	"github.com/Archisman-Mridha/go-microservice-template/internal/core/validators"
	"github.com/Archisman-Mridha/go-microservice-template/internal/token"
	"github.com/Archisman-Mridha/go-microservice-template/pkg/types"
	"github.com/Archisman-Mridha/go-microservice-template/pkg/utils"
	"github.com/go-playground/validator/v10"
	goValidator "github.com/go-playground/validator/v10"
)

type Usecases struct {
	validator *validator.Validate

	cache           types.KVStore
	usersRepository coreTypes.UsersRepository

	tokenService token.TokenService
}

func NewUsecases(
	validator *validator.Validate,
	cache types.KVStore,
	usersRespository coreTypes.UsersRepository,
	tokenService token.TokenService,
) *Usecases {
	utils.RegisterCustomFieldValidators(validator, map[string]goValidator.Func{
		"name":     validators.NameFieldValidator,
		"email":    validators.EmailFieldValidator,
		"username": validators.UsernameFieldValidator,
		"password": validators.PasswordFieldValidator,
	})

	return &Usecases{
		validator,
		cache,
		usersRespository,
		tokenService,
	}
}
