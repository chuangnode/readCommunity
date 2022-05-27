package service

import (
	"readCommunity/internal/userserver/model"
)

func CheckAuth(username, password string) (bool, error) {
	var auth model.User
	return auth.CheckAuth(username, password)
}
