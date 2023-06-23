package user

import "gbl-api/data"

func RegisterUser(u User) error {
	db := data.GetDatabase()
	return db.Create(&u).Error
}
