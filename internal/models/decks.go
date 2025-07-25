package models

import "time"

type Deck struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Question struct {
	ID           int       `json:"id"`
	DeckID       int       `json:"deckId"`
	QuestionText string    `json:"questionText"`
	Answer       string    `json:"answer"`
	Options      []string  `json:"options"`
	CreatedAt    time.Time `json:"createdAt"`
}
