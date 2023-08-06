// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import ()

type LearningItem struct {
	ID               string   `json:"id"`
	ImageLink        string   `json:"image_link"`
	EnglishWord      string   `json:"english_word"`
	VietnameseWord   string   `json:"vietnamese_word"`
	EnglishSentences []string `json:"english_sentences"`
}

type Test struct {
	TestID string `json:"test_id"`
}

type User struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
