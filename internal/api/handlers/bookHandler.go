package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Werp-ES/book_store/internal/services"
)

func FindBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(string(r.URL.Path[len("/book/")]))
	if err != nil {
		http.Error(w, "That is not a number!", http.StatusBadRequest)
		return
	}

	b, err := services.GetBookById(id)

	if err != nil {
		fmt.Println(err.Error())
	}

	bb, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, string(bb))

}
