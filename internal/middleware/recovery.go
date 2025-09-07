package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v\n%s", err, debug.Stack())

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				response := map[string]interface{}{
					"status": "error",
					"error":  "Internal server error",
				}
				if encErr := json.NewEncoder(w).Encode(response); encErr != nil {
					log.Printf("Error encoding panic recovery response: %v", encErr)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}