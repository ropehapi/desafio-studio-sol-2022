package middleware

import "net/http"

/*
	Middleware que interceptará todas as nossas requisições afim de setar o Content-Type da nossa 
	aplicação para devolver JSON.
*/
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}
