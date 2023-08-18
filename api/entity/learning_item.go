package entity

import (
	db "lla/db/sqlc"
	"time"
)

type LearningItem struct {
	ID               string    `json:"id"`
	ImageLink        string    `json:"image_link"`
	EnglishWord      string    `json:"english_word"`
	VietnameseWord   string    `json:"vietnamese_word"`
	EnglishSentences []string  `json:"english_sentences"`
	CreatedAt        time.Time `json:"created_at"`
	UserID           string    `json:"user_id"`
	TopicID          string    `json:"topic_id"`
}

func CreateLearningItemFromLI(li *db.LearningItem) *LearningItem {
	return &LearningItem{
		ID:               li.ID,
		ImageLink:        li.ImageLink,
		EnglishWord:      li.EnglishWord,
		VietnameseWord:   li.VietnameseWord.String,
		EnglishSentences: li.EnglishSentences,
		CreatedAt:        li.CreatedAt,
		UserID:           li.UserID.String,
		TopicID:          li.TopicID.String,
	}
}

func CreateLearningItemsFromLIs(lis []db.LearningItem) []*LearningItem {
	learningItems := make([]*LearningItem, len(lis))
	for i, li := range lis {
		learningItems[i] = CreateLearningItemFromLI(&li)
	}
	return learningItems
}
