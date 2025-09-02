package user

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"to-do/contract"
)

type UserServiceMock struct {
	mock.Mock
}

func (u *UserServiceMock) LoginUser(ctx *gin.Context, user *contract.LoginUser) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

func (u *UserServiceMock) GetUserIdByUserName(username string) (int64, error) {
	args := u.Called(username)
	return int64(args.Int(0)), args.Error(1)
}

func (u *UserServiceMock) CreateUser(ctx *gin.Context, userSignUpRequest *contract.SignUpUser) error {
	args := u.Called(ctx, userSignUpRequest)
	return args.Error(0)
}
