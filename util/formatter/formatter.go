package formatter

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorFormatter(w http.ResponseWriter, message string, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(map[string]string{
        "error": message,
    })
	return
}

func DataFormatter(w http.ResponseWriter, data interface{}, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
        "data": data,
    }); err != nil {
		log.Println(err)
		ErrorFormatter(w, "Cannot process", http.StatusInternalServerError)
	}
	return
}

func MessageFormatter(w http.ResponseWriter, data string, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
        "message": data,
    }); err != nil {
		log.Println(err)
		ErrorFormatter(w, "Cannot process", http.StatusInternalServerError)
	}
	return
}
