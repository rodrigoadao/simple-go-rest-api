package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func (wr http.ResponseWriter, r *http.Request) {
        wr.Header().Set("Content-type", "application/json")
        next.ServeHTTP(wr, r)
    })
}
