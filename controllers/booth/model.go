package booth

type Booth struct {
	BID          string   `json:"bid"`
	Name         string   `json:"name"`
	Part         string   `json:"part"`
	Complexity   int      `json:"complexity"`
	VideoURL     string   `json:"video_url"`
	ThumbnailURL string   `json:"thumbnail_url"`
	ProblemOrder []string `json:"problem_order"`
}
