package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
	"to-do/contract"
	"to-do/domain"
	"to-do/repo"
)

func Test_userService_GetUserIdByUserName(t *testing.T) {
	type fields struct {
		usernameToUserIdMap *domain.UsernameToUserIdMap
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test user already exists",
			fields: fields{
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M: map[string]int64{
						"user": 4,
					},
				},
			},
			args: args{
				username: "user",
			},
			want:    4,
			wantErr: true,
		},
		{
			name: "test user does not exist",
			fields: fields{
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M: map[string]int64{
						"user": 4,
					},
					LastUserId: 4,
				},
			},
			args: args{
				username: "user_123",
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				usernameToUserIdMap: tt.fields.usernameToUserIdMap,
			}
			got, err := u.GetUserIdByUserName(tt.args.username)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, got, "GetUserIdByUserName(%v)", tt.args.username)
		})
	}
}

func Test_userService_CreateUser(t *testing.T) {
	type fields struct {
		userRepo            repo.UserRepoMock
		usernameToUserIdMap *domain.UsernameToUserIdMap
	}
	type args struct {
		ctx  *gin.Context
		user *contract.SignUpUser
	}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test username taken",
			fields: fields{
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M:          map[string]int64{"user": 4},
					LastUserId: 4,
				},
			},
			args: args{
				ctx: ctx,
				user: &contract.SignUpUser{
					Username: "user",
				},
			},
			wantErr: true,
		},
		{
			name: "test new user, invalid password",
			fields: fields{
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M:          map[string]int64{"user": 4},
					LastUserId: 4,
				},
			},
			args: args{
				ctx: ctx,
				user: &contract.SignUpUser{
					Username: "user@123",
					Password: "user0213#",
				},
			},
			wantErr: true,
		},
		{
			name: "test new user, add user fail",
			fields: fields{
				userRepo: repo.UserRepoMock{},
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M:          map[string]int64{"user": 4},
					LastUserId: 4,
				},
			},
			args: args{
				ctx: ctx,
				user: &contract.SignUpUser{
					Name:        "Venkat",
					Username:    "user@123",
					Password:    "User01@#",
					PhoneNumber: "9900923821",
				},
			},
			wantErr: true,
		},
		{
			name: "test new user, add user success",
			fields: fields{
				userRepo: repo.UserRepoMock{},
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M:          map[string]int64{"user": 4},
					LastUserId: 4,
				},
			},
			args: args{
				ctx: ctx,
				user: &contract.SignUpUser{
					Name:        "Venkat",
					Username:    "user@123",
					Password:    "User01@#",
					PhoneNumber: "9900923821",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepo:            &tt.fields.userRepo,
				usernameToUserIdMap: tt.fields.usernameToUserIdMap,
			}
			if tt.name == "test new user, add user fail" {
				tt.fields.userRepo.On("AddNewUser", tt.args.ctx, mock.Anything).Return(fmt.Errorf("err")).Once()
			}
			if tt.name == "test new user, add user success" {
				tt.fields.userRepo.On("AddNewUser", tt.args.ctx, mock.Anything).Return(nil).Once()
			}
			err := u.CreateUser(tt.args.ctx, tt.args.user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_userService_LoginUser(t *testing.T) {
	type fields struct {
		userRepo            repo.UserRepoMock
		usernameToUserIdMap *domain.UsernameToUserIdMap
	}
	type args struct {
		ctx           *gin.Context
		userLoginInfo *contract.LoginUser
	}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test user not identified",
			fields: fields{
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M: map[string]int64{"user": 4},
				},
			},
			args: args{
				ctx: ctx,
				userLoginInfo: &contract.LoginUser{
					Username: "venkxy7codes1",
					Password: "Venkxycodes@123",
				},
			},
			wantErr: true,
		},
		{
			name: "test user identified, get user details error",
			fields: fields{
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M: map[string]int64{"venkxy7codes1": 4},
				},
				userRepo: repo.UserRepoMock{},
			},
			args: args{
				ctx: ctx,
				userLoginInfo: &contract.LoginUser{
					Username: "venkxy7codes1",
					Password: "Venkxycodes@123",
				},
			},
			wantErr: true,
		},
		{
			name: "test user identified, password mismatch",
			fields: fields{
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M: map[string]int64{"venkxy7codes1": 4},
				},
				userRepo: repo.UserRepoMock{},
			},
			args: args{
				ctx: ctx,
				userLoginInfo: &contract.LoginUser{
					Username: "venkxy7codes1",
					Password: "Venkxycodes@123",
				},
			},
			wantErr: true,
		},
		{
			name: "test successful login",
			fields: fields{
				usernameToUserIdMap: &domain.UsernameToUserIdMap{
					M: map[string]int64{"venkxy7codes1": 4},
				},
				userRepo: repo.UserRepoMock{},
			},
			args: args{
				ctx: ctx,
				userLoginInfo: &contract.LoginUser{
					Username: "venkxy7codes1",
					Password: "Venkxycodes@123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepo:            &tt.fields.userRepo,
				usernameToUserIdMap: tt.fields.usernameToUserIdMap,
			}
			if tt.name == "test user identified, get user details error" {
				var user *domain.User
				tt.fields.userRepo.On("GetUserByUserId", tt.args.ctx, int64(4)).Return(user, fmt.Errorf("err")).Once()
			}
			if tt.name == "test user identified, password mismatch" {
				tt.fields.userRepo.On("GetUserByUserId", tt.args.ctx, int64(4)).Return(&domain.User{
					Username: "venkxy7codes1",
					UserId:   int64(4),
					Password: "venkxycodes@123",
				}, nil).Once()
			}
			if tt.name == "test successful login" {
				tt.fields.userRepo.On("GetUserByUserId", tt.args.ctx, int64(4)).Return(&domain.User{
					Username: "venkxy7codes1",
					UserId:   int64(4),
					Password: "Venkxycodes@123",
				}, nil).Once()
			}
			err := u.LoginUser(tt.args.ctx, tt.args.userLoginInfo)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
