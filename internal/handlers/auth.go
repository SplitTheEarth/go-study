package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"study-app/internal/database"
	"study-app/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// SignupHandler handles user registration
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var signupData models.SignupPayload

	err := json.NewDecoder(r.Body).Decode(&signupData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if signupData.Username == "" || signupData.Email == "" || signupData.Password == "" {
		http.Error(w, "Username, email, and password are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error, unable to create your account.", http.StatusInternalServerError)
		return
	}

	insertSQL := `INSERT INTO users(username, email, password_hash) VALUES (?, ?, ?)`
	_, err = database.DB.Exec(insertSQL, signupData.Username, signupData.Email, string(hashedPassword))
	if err != nil {
		// This can fail if the username or email is already taken (due to the UNIQUE constraint)
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			http.Error(w, "Username or email already exists", http.StatusConflict)
		} else {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
		}
		return
	}

	fmt.Printf("New user created: Username=%s, Email=%s\n", signupData.Username, signupData.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

// LoginHandler handles user authentication
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginPayload models.LoginPayload
	err := json.NewDecoder(r.Body).Decode(&loginPayload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if loginPayload.Email == "" || loginPayload.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	var storedHash string
	var userID int
	var username string
	checkSQL := `SELECT id, username, password_hash FROM users WHERE email = ?`
	err = database.DB.QueryRow(checkSQL, loginPayload.Email).Scan(&userID, &username, &storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "Error logging in", http.StatusInternalServerError)
		}
		return
	}

	storedHash = strings.TrimSpace(storedHash)
	fmt.Printf("Login attempt for user: %s (ID: %d)\n", username, userID)

	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(loginPayload.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	fmt.Printf("Successful login for user: %s\n", username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Login successful",
		"user_id":  userID,
		"username": username,
	})
}
