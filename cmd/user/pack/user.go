package pack

import (
	"github.com/BlueGopher/tiktok-rpc/cmd/user/dal/db"
	usercore "github.com/BlueGopher/tiktok-rpc/kitex_gen/userCore"
)

// User pack user info
func User(user *db.User) *usercore.User {
	if user == nil {
		return nil
	}
	return &usercore.User{
		Id:            int64(user.ID),
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
}

// Users pack list of user info
func Users(us []*db.User) []*usercore.User {
	users := make([]*usercore.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
