package user

import "gbl-api/data"

func RegisterUser(u User) error {
	db := data.GetDatabse()
	return db.Create(&u).Error
}
