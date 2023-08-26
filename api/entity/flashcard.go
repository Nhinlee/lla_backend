package entity

import (
	db "lla/db/sqlc"
	"time"
)

type FlashCard struct {
	ID          string `json:"id"`
	ImageLink   string `json:"image_link"`
	EnglishWord string `json:"english_word"`
	// VietnameseWord   string   `json:"vietnamese_word"`
	EnglishSentences []string  `json:"english_sentences"`
	CreatedAt        time.Time `json:"created_at"`
	CompletedAt      time.Time `json:"completed_at"`
}

func CreateFlashCardFromLI(li *db.LearningItem) *FlashCard {
	fl := &FlashCard{
		ID:               li.ID,
		ImageLink:        li.ImageLink,
		EnglishWord:      li.EnglishWord,
		EnglishSentences: li.EnglishSentences,
		CreatedAt:        li.CompletedAt.Time,
	}

	if li.CompletedAt.Valid {
		fl.CompletedAt = li.CompletedAt.Time
	}

	return fl
}

func CreateFlashCardsFromLIs(lis []db.LearningItem) []*FlashCard {
	flashcards := make([]*FlashCard, len(lis))
	for i, li := range lis {
		flashcards[i] = CreateFlashCardFromLI(&li)
	}
	return flashcards
}
