package booth

import "gbl-api/data"

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
