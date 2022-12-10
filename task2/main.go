package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// The upperCaseHandler function converts the word parameter to uppercase and returns it.
// Example usage:
// curl http://localhost:8080/upper?word=hello
// HELLO
func upperCaseHandler(w http.ResponseWriter, req *http.Request) {
	word := req.URL.Query().Get("word")
	fmt.Fprint(w, strings.ToUpper(word))
}

// The dbIncrementHandler function increments the request count in the database.
func dbIncrementHandler(w http.ResponseWriter, req *http.Request) {
	username := os.Getenv("DB_USERNAME")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")

	// the database name is hardcoded here, but it could be passed in as an environment variable
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/app_db",
		username, password, host))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO requests (count) VALUES(NULL)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Fprint(w, "Request count incremented.")
}

// The main function starts the HTTP server.
func main() {
	http.HandleFunc("/upper", upperCaseHandler)

	// register the dbIncrementHandler only if the ENABLE_DB environment variable is set to true
	if os.Getenv("ENABLE_DB") == "true" {
		http.HandleFunc("/db", dbIncrementHandler)
	}

	http.ListenAndServe(":8080", nil)
}
