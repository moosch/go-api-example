package api

import (
    "encoding/json"
    "net/http"
)

type LoginParams struct {
    Username string
}

type LoginResponse struct {
    Username  string
    AuthToken string
}

type LogoutParams struct {
    Username string
}

type GetAccountParams struct {
    Username string
}

type GetAccountResponse struct {
    Username string
    Age      int
}

type UpdateAccountParams struct {
    Username string
    Age      int
}

type UpdateAccountBody struct {
    Username string
    Age      int
}

type UpdateAccountResponse struct {
    Username string
    Age      int
}

type DeleteAccountParams struct {
    Username int
}

type Error struct {
    Code int
    Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
    resp := Error{
        Code: code,
        Message: message,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)

    json.NewEncoder(w).Encode(resp)
}

var (
    NotFoundErrorHandler = func(w http.ResponseWriter, message string) {
        writeError(w, message, http.StatusNotFound)
    }
    RequestErrorHandler = func(w http.ResponseWriter, err error) {
        writeError(w, err.Error(), http.StatusBadRequest)
    }
    InternalErrorHandler = func(w http.ResponseWriter) {
        writeError(w, "An unexpected error occurred.", http.StatusInternalServerError)
    }
)

