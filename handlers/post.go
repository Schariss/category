package handlers

import (
	"github.com/Schariss/category-api/data"
	"net/http"
)

func (q *Categories) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the category from the context
	cat := r.Context().Value(KeyProduct{}).(data.Category)

	q.l.Printf("[DEBUG] Inserting category: %#v\n", cat)
	data.AddCategory(cat)
}
