package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId      int64              `json:"user_id" bson:"user_id"`
	Name        string             `json:"name" bson:"name"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
}

type UsernameToUserIdMap struct {
	sync.RWMutex
	M          map[string]int64
	LastUserId int64
}

func (u *UsernameToUserIdMap) Get(username string) (int64, int64) {
	u.RLock()
	defer u.RUnlock()
	if userId, isPresent := u.M[username]; isPresent {
		return userId, u.LastUserId
	} else {
		return 0, u.LastUserId
	}
}

func (u *UsernameToUserIdMap) Set(username string, userId int64) {
	u.Lock()
	defer u.Unlock()
	if username != "" {
		u.M[username] = userId
		u.LastUserId += 1
	}
}
