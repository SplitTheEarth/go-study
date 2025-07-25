package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"study-app/internal/database"
	"study-app/internal/models"
	"time"
)

// CreateDeckHandler handles creating a new deck
func CreateDeckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var deckData models.Deck
	err := json.NewDecoder(r.Body).Decode(&deckData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if deckData.Name == "" || deckData.Description == "" {
		http.Error(w, "Name and description are required", http.StatusBadRequest)
		return
	}

	insertSQL := `INSERT INTO decks (name, description) VALUES (?, ?)`
	res, err := database.DB.Exec(insertSQL, deckData.Name, deckData.Description)
	if err != nil {
		http.Error(w, "Failed to create deck", http.StatusInternalServerError)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to create deck", http.StatusInternalServerError)
		return
	}
	deckData.ID = int(id)
	// We can't get the DB-generated CreatedAt without another query,
	// so we'll approximate it. For a perfect value, we'd need to re-fetch the deck.
	deckData.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(deckData)
}

func GetDeckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	selectSQL := `SELECT id, name, description, created_at FROM decks WHERE id = ?`
	row := database.DB.QueryRow(selectSQL, id)

	var deck models.Deck
	err := row.Scan(&deck.ID, &deck.Name, &deck.Description, &deck.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Deck not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deck)
}

func GetDecksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	selectSQL := `SELECT id, name, description, created_at FROM decks`
	rows, err := database.DB.Query(selectSQL)
	if err != nil {
		http.Error(w, "Failed to fetch decks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var decks []models.Deck
	for rows.Next() {
		var deck models.Deck
		err := rows.Scan(&deck.ID, &deck.Name, &deck.Description, &deck.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to scan deck", http.StatusInternalServerError)
			return
		}
		decks = append(decks, deck)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(decks)
}

func AddQuestionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var question models.Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		http.Error(w, "Failed to decode question", http.StatusBadRequest)
		return
	}
	options, err := json.Marshal(question.Options)
	if err != nil {
		http.Error(w, "Failed to encode options", http.StatusInternalServerError)
		return
	}
	question.CreatedAt = time.Now()
	insertSQL := `
		INSERT INTO questions(deck_id, question_text, answer, options)
		VALUES (?, ?, ?, ?)
	`
	result, err := database.DB.Exec(insertSQL, question.DeckID, question.QuestionText, question.Answer, options)
	if err != nil {
		http.Error(w, "Failed to add question", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to get last inserted ID", http.StatusInternalServerError)
		return
	}
	question.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}

func ListQuestionsByDeckHandler(w http.ResponseWriter, r *http.Request) {
	deckIDSrt := r.URL.Query().Get("deck_id")
	if deckIDSrt == "" {
		http.Error(w, "Deck ID is required", http.StatusBadRequest)
		return
	}

	deckID, err := strconv.Atoi(deckIDSrt)
	if err != nil {
		http.Error(w, "Invalid deck ID", http.StatusBadRequest)
		return
	}

	selectSQL := `
		SELECT id, deck_id, question_text, answer, options, created_at
		FROM questions
		WHERE deck_id = ?
	`
	rows, err := database.DB.Query(selectSQL, deckID)
	if err != nil {
		http.Error(w, "Failed to query questions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var q models.Question
		var optionsStr string // Temporary variable to hold the options JSON
		err := rows.Scan(&q.ID, &q.DeckID, &q.QuestionText, &q.Answer, &optionsStr, &q.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to scan question", http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal([]byte(optionsStr), &q.Options); err != nil {
			http.Error(w, "Failed to unmarshal options", http.StatusInternalServerError)
			return
		}

		questions = append(questions, q)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)

}
