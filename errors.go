package category_nurzhas_store

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrCreateCategoryUnknown  = com.NewMiddleError(errors.New("could not create category: unknown error"), 500, 1)
	ErrCategoryNotFound       = com.NewMiddleError(errors.New("category not found"), 404, 2)
	ErrNothingToUpdate        = com.NewMiddleError(errors.New("nothing to update"), 400, 3)
	ErrCategoryIdNotProvided  = com.NewMiddleError(errors.New("category id is not provided"), 400, 4)
	ErrCreateUserUnknown      = com.NewMiddleError(errors.New("could not create user: unknown error"), 500, 5)
	ErrUserNotFound           = com.NewMiddleError(errors.New("user not found"), 404, 6)
	ErrUserIdNotProvided      = com.NewMiddleError(errors.New("user id is not provided"), 400, 7)
	ErrUserPasswordNotCorrect = com.NewMiddleError(errors.New("user password not correct"), 500, 8)
	ErrCreateTelegramUnknown  = com.NewMiddleError(errors.New("could not create telegram bot:unknown error"), 500, 9)
	ErrTelegramNotFound       = com.NewMiddleError(errors.New("telegram not found"), 404, 10)
)
