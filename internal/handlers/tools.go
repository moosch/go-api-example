package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"

	"github.com/moosch/go-api-example/api"
)

func parseRequestUrlQuery[T interface{}](query url.Values, s *T) error {
    var decoder = schema.NewDecoder()
    var err error

    err = decoder.Decode(s, query)
    if err != nil {
        log.Error(err)
        return errors.New("Invalid URL query.")
    }
    return nil
}

func parseRequestBody[T interface{}](body io.Reader, s *T) error {
    var decoder = json.NewDecoder(body)
    if err := decoder.Decode(s); err != nil {
        log.Error(err)
        return errors.New("Invalid Body.")
    }
    return nil
}

func sendResponse[T interface{}](w http.ResponseWriter, response T) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
    }
}

func sendEmptyResponse(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNoContent)
}

