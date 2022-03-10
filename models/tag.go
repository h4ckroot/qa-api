package models

func (b *Tag) TableName() string {
	return "tag"
}

type Tag struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}
