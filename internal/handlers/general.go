package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"study-app/internal/database"
	"study-app/internal/models"
)

// HelloHandler handles the root endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the backend!")
}

// StatusHandler handles the API status endpoint
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

func SubmitAnswerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var answerPayload models.AnswerPayload
	err := json.NewDecoder(r.Body).Decode(&answerPayload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if answerPayload.UserID == 0 && answerPayload.QuestionID == 0 && answerPayload.Answer == "" {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	var correctAnswer string
	selectSQL := `SELECT answer FROM questions WHERE id = ?`
	err = database.DB.QueryRow(selectSQL, answerPayload.QuestionID).Scan(&correctAnswer)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve question", http.StatusInternalServerError)
		}
		return
	}

	if answerPayload.Answer != correctAnswer {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"correct": false, "message": "Sorry, that's not the right answer."})
		return
	}

	updateSQL := `UPDATE users SET score = score + 1 WHERE id = ?`
	_, err = database.DB.Exec(updateSQL, answerPayload.UserID)
	if err != nil {
		http.Error(w, "Failed to update user score", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"correct": true, "message": "You got it!"})
}
