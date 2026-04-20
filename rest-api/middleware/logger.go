package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

// nested handlers are only used  for middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("START %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("END %s %s", r.Method, r.URL.Path)
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			//For catching unexpected errors, we can use recover() that works in a similar way to try catch thus it stops the error sequence before it terminates the program but still returns the err info

			if err := recover(); err != nil {
				log.Printf("PANIC: %v\n%s", err, debug.Stack())

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error":"internal server error"}`))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
