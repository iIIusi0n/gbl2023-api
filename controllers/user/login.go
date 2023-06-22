package user

import (
	"gbl-api/data"
)

func IsUserExist(uid string) bool {
	db := data.GetDatabse()
	var user User
	db.Where("uid = ?", uid).First(&user)
	return user.UID != ""
}
