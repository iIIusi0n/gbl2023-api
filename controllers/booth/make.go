package booth

import "gbl-api/data"

func MakeBooth(b Booth) error {
	db := data.GetDatabse()
	return db.Create(&b).Error
}
