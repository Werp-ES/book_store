package routers

import (
	"net/http"

	"github.com/Werp-ES/book_store/internal/api/handlers"
)

func NewServeMux() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/book/", handlers.FindBookHandler)

	return mux
}
