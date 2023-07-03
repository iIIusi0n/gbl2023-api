package booth

import (
	"gbl-api/data"
	"gorm.io/gorm"
)

func GetBooths() ([]Booth, error) {
	db := data.GetDatabase()
	var booths []Booth
	err := db.Find(&booths).Error
	return booths, err
}

func GetBooth(bid string) (Booth, error) {
	db := data.GetDatabase()
	var booth Booth
	err := db.Where("bid = ?", bid).First(&booth).Error
	return booth, err
}

func GetBoothByPassword(password string) (Booth, error) {
	db := data.GetDatabase()
	var boothPw BoothPassword
	err := db.Where("password = ?", password).First(&boothPw).Error
	if err == gorm.ErrRecordNotFound {
		return Booth{}, nil
	} else if err != nil {
		return Booth{}, err
	} else {
		return GetBooth(boothPw.BID)
	}
}

func DeleteBooth(bid string) error {
	db := data.GetDatabase()
	return db.Delete(&Booth{}, "bid = ?", bid).Error
}
