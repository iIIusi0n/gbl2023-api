package booth

import "github.com/lib/pq"

type Booth struct {
	BID          string         `json:"bid"`
	Name         string         `json:"name"`
	Part         string         `json:"part"`
	Complexity   int            `json:"complexity"`
	VideoURL     string         `json:"video_url"`
	ThumbnailURL string         `json:"thumbnail_url"`
	ProblemOrder pq.StringArray `gorm:"type:text[]" json:"problem_order"`
	UIDs         []string       `json:"uids"`
}

type BoothPassword struct {
	BID      string `json:"bid"`
	Password string `json:"password"`
}
