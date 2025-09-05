package user

import (
	"epictectus/contract"
	"epictectus/domain"
	"epictectus/repo"
	"epictectus/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
)

type userService struct {
	userRepo            repo.UserRepository
	usernameToUserIdMap *domain.UsernameToUserIdMap
}

type UserService interface {
	GetUserIdByUserName(username string) (int64, error)
	CreateUser(ctx *gin.Context, user *contract.SignUpUser) error
	LoginUser(ctx *gin.Context, user *contract.LoginUser) error
}

func NewUserService(userRepo repo.UserRepository) UserService {
	u := &userService{
		userRepo: userRepo,
	}
	u.usernameToUserIdMap = &domain.UsernameToUserIdMap{M: make(map[string]int64)}
	// Check if there are existing users on Db, if yes, populate them on the map, if no, we will create an empty map
	// Existing in memory map gets cleared when we restart server
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	users, err := userRepo.GetAllUsers(ctx)
	if err != nil || len(users) == 0 {
		return u
	}
	for _, user := range users {
		u.usernameToUserIdMap.Set(user.Username, user.UserId)
	}
	return u
}

// GetUserIdByUserName method returns userId and an error when a user already exists,
// returns the last userId and no error when a user does not exist
func (u *userService) GetUserIdByUserName(username string) (int64, error) {
	userId, lastUserId := u.usernameToUserIdMap.Get(username)
	if userId != 0 {
		return userId, fmt.Errorf("err-username-already-taken")
	}
	return lastUserId, nil
}

func (u *userService) CreateUser(ctx *gin.Context, user *contract.SignUpUser) error {
	lastUserId, err := u.GetUserIdByUserName(user.Username)
	if err != nil {
		return err
	}
	if err := utils.ValidatePassword(user.Password); err != "" {
		return fmt.Errorf(err)
	}
	createErr := u.userRepo.AddNewUser(ctx, &domain.User{
		Id:          primitive.NewObjectID(),
		Name:        user.Name,
		Username:    user.Username,
		Password:    user.Password,
		UserId:      lastUserId + 1,
		PhoneNumber: user.PhoneNumber,
	})
	if createErr != nil {
		return createErr
	}
	u.usernameToUserIdMap.Set(user.Username, lastUserId+1)
	return nil
}

func (u *userService) LoginUser(ctx *gin.Context, userLoginInfo *contract.LoginUser) error {
	userId, err := u.GetUserIdByUserName(userLoginInfo.Username)
	if err == nil {
		return fmt.Errorf("err-username-not-identified")
	}
	userDetails, getUserErr := u.userRepo.GetUserByUserId(ctx, userId)
	if getUserErr != nil {
		return err
	}
	if userDetails.Password != userLoginInfo.Password {
		return fmt.Errorf("err-incorrect-password")
	}
	return nil
}
