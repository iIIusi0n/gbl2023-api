package score

type Participation struct {
	BID   string `json:"bid"`
	UID   string `json:"uid"`
	PID   string `json:"pid"`
	Score int    `json:"score"`
}
