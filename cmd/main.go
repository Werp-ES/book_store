package main

import (
	"net/http"

	"github.com/Werp-ES/book_store/internal/api/routers"
	"github.com/Werp-ES/book_store/internal/data"
)

func main() {

	data.Connect()

	mux := routers.NewServeMux()

	http.ListenAndServe(":8081", mux)

}
