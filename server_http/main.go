package main

import (
	"encoding/json"
	"fmt"
	jwt "github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = make([]Book, 0)
var jwtKey = []byte("segredo")

func setupRoutes() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		bookHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				createBookHandler(w, r)
			case http.MethodGet:
				listBooksHandler(w, r)
			default:
				http.Error(w, "Invalid method", 404)
			}
		})
		middlewareHandler := applyMiddleware(bookHandler, loggingMiddleware, jwtMiddleware)
		middlewareHandler.ServeHTTP(w, r)
	})
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	books = append(books, newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func listBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	author := r.URL.Query().Get("author")
	if author != "" {
		books = getBooksByAuthor(author)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func getBooksByAuthor(author string) []Book {
	var result []Book
	for _, book := range books {
		if book.Author == author {
			result = append(result, book)
		}
	}

	return result
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("receive request:%s %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func applyMiddleware(handler http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		handler = m(handler)
	}

	return handler
}

func generateJWT() (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	token, err := generateJWT()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const bearerPrefix = "Bearer "
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(bearerPrefix):]
		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	setupRoutes()
	fmt.Println("livrago server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("failed to start server", err)
	}
}
