package endpoint

import (
	"context"
	"encoding/json"
	"net/http"
	translateEntities "translator/endpoint/translate"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpHandler(endpoints EndPoints) http.Handler {
	translateHandler := httptransport.NewServer(endpoints.Translate, DecodeTranslateRequest, EncodeTranslateResponse)

	r := mux.NewRouter()
	r.Handle("/translator/translate", translateHandler).Methods(http.MethodPost)

	return r
}

func DecodeTranslateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req translateEntities.Request
	err := json.NewDecoder(r.Body).Decode(&req)

	return req, err
}

func EncodeTranslateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := response.(translateEntities.Response)
	encodeError(ctx, res.TextoTraducido, w)
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, texto string, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	if texto != "" {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusAccepted)
	}
}
