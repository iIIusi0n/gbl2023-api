package user

import (
	"gbl-api/data"
)

func IsUserExist(uid string) bool {
	db := data.GetDatabase()
	var user User
	db.Where("uid = ?", uid).First(&user)
	return user.UID != ""
}
