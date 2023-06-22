package score

import (
	"gbl-api/data"
	"gorm.io/gorm"
)

func IsUserParticipated(bid, uid string) (bool, error) {
	db := data.GetDatabse()
	var score Participation
	err := db.Where("bid = ? AND uid = ?", bid, uid).First(&score).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
