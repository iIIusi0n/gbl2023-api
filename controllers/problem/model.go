package problem

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Problem struct {
	BID         string `json:"bid"`
	PID         string `json:"pid"`
	PdfURL      string `json:"pdf_url"`
	ProblemType string `json:"problem_type"`
	Question    string `json:"question"`
	Answer      string `json:"answer"`
	Score       int    `json:"score"`
}

func (p *Problem) BeforeCreate(tx *gorm.DB) (err error) {
	p.PID = uuid.New().String()
	return
}
