package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/api/v1/users/1", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{
			Code:      "OK",
			Message:   "success",
			RequestID: "req-" + time.Now().Format("20060102150405"),
			Data: map[string]interface{}{
				"user_id":  1,
				"username": "gopher",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	log.Println("listening on :8891")
	log.Fatal(http.ListenAndServe(":8891", nil))
}
