package services

import (
	"database/sql"
	"fmt"

	"github.com/Werp-ES/book_store/internal/data"
	"github.com/Werp-ES/book_store/internal/models"
)

func GetBookById(id int) (models.Book, error) {
	b, err := data.FindBookById(id)

	if err == sql.ErrNoRows {
		return b, fmt.Errorf("there is no such book with id: %d", id)
	} else if err != nil {
		return b, fmt.Errorf("%q: %v", id, err)
	}

	return b, err
}
