package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

const shortBaseURL = "http://localhost:8080/"

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req URLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.OriginalURL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortCode := generateCode(6)

	// lock for writing
	urlStore.Lock()
	urlStore.m[shortCode] = req.OriginalURL
	urlStore.Unlock()

	resp := URLResponse{ShortURL: shortBaseURL + shortCode}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:] // trim "/"

	// lock for reading
	urlStore.RLock()
	originalURL, ok := urlStore.m[code]
	urlStore.RUnlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func generateCode(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}
