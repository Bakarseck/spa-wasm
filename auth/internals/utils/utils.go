package utils

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../db/clash.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RespondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	if _, err := w.Write(jsonData); err != nil {
		return
	}
}

func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println("🚨 " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("🚨 Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}
