package handlers

import (
    "net/http"

    "github.com/moosch/go-api-example/internal/middleware"
)

func Handler(m *http.ServeMux) {
    m.Handle("GET /account", middleware.Authorization(http.HandlerFunc(GetAccount)))
    m.Handle("POST /account", middleware.Authorization(http.HandlerFunc(UpdateAccount)))
    m.Handle("DELETE /account", middleware.Authorization(http.HandlerFunc(DeleteAccount)))

    m.Handle("POST /account/login", http.HandlerFunc(Login))
    m.Handle("POST /account/logout", http.HandlerFunc(Logout))
}

