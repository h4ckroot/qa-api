package models

import (
	"github.com/rs/zerolog/log"
	"shamilabs.com/qa-api/config"
)

func CreateAnswer(answer *Answer) (err error) {
	log.Info().Msgf("Creating new question with title: %s", answer.Text)
	if err = config.GetDB().Create(answer).Error; err != nil {
		return err
	}
	return nil
}

// Answer type
type Answer struct {
	ID         uint   `json:"id"`
	Text       string `json:"text"`
	QuestionID uint
	Question   Question
}

func (b *Answer) TableName() string {
	return "answer"
}
