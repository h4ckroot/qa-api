package models

import (
	"github.com/rs/zerolog/log"
	"shamilabs.com/qa-api/config"
)

func CreateAQuestion(question *Question) (err error) {
	log.Info().Msgf("Creating new question with title: %s", question.Title)
	if err = config.GetDB().Create(&question).Error; err != nil {
		log.Error().Err(err).Msg("could not create question")
		return err
	}
	return nil
}

func TagQuestion(question *Question, tag *Tag) {
	db := config.GetDB()
	log.Info().Msgf("Updating question(id): %d, with tags: %s", question.ID, tag)
	db.Model(&question).Association("Tags").Append(tag)
	log.Info().Msgf("Found question: %s", question.Text)

}

// Question type
type Question struct {
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Tags    []Tag    `json:"tags" gorm:"many2many:question_tags;"`
	Answers []Answer `json:"answers"`
}

func (b *Question) TableName() string {
	return "question"
}
