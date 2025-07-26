package models

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Score    int    `json:"score"`
}

// SignupPayload represents the data needed to create a new user
type SignupPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginPayload represents the data needed for user login
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AnswerPayload struct {
	UserID     int    `json:"user_id"`
	QuestionID int    `json:"question"`
	Answer     string `json:"answer"`
}
