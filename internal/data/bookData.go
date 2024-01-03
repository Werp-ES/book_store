package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Werp-ES/book_store/internal/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
	cfgStr := "user=admin password=1234 dbname=books host=localhost port=5432 sslmode=disable"
	fmt.Println("Conecting to database...")

	var err error
	db, err = sql.Open("postgres", cfgStr)
	if err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}

func FindBookById(id int) (models.Book, error) {
	var b models.Book
	row := db.QueryRow("SELECT * FROM books WHERE id = $1", id)

	err := row.Scan(&b.Id, &b.Title, &b.Author, &b.Price)

	return b, err
}
