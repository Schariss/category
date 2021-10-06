package handlers

import (
	"context"
	"github.com/Schariss/category-api/data"
	"net/http"
)

// MiddlewareValidateCategory validates the product in the request and calls next if ok
func (p *Categories) MiddlewareValidateCategory(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cat := &data.Category{}

		err := data.FromJSON(cat, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing category", err)

			rw.WriteHeader(http.StatusBadRequest)
			http.Error(rw, "unable to decode json", http.StatusInternalServerError)
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, cat)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
