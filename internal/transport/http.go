package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hitesh3602/master_languages/internal/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(svc service.LanguageService) http.Handler {
	router := mux.NewRouter()

	router.Methods("POST").Path("/rms/master_languages").Handler(httptransport.NewServer(
		makeCreateLanguageEndpoint(svc),
		decodeCreateLanguageRequest,
		encodeResponse,
	))	

	// router.Methods("GET").Path("/rms/master_languages").Handler(httptransport.NewServer(
	// 	makeGetLanguagesEndpoint(svc),
	// 	decodeEmptyRequest,
	// 	encodeResponse,
	// ))

	// router.Methods("GET").Path("/rms/master_languages/{id}").Handler(httptransport.NewServer(
	// 	makeGetLanguageByIDEndpoint(svc),
	// 	decodeGetLanguageByIDRequest,
	// 	encodeResponse,
	// ))

	router.Methods("GET").Path("/rms/master_languages").Handler(httptransport.NewServer(
		makeGetLanguagesEndpoint(svc),
		decodeGetLanguagesRequest,
		encodeResponse,
	))

	router.Methods("PUT").Path("/rms/master_languages/{id}").Handler(httptransport.NewServer(
		makeUpdateLanguageEndpoint(svc),
		decodeUpdateLanguageRequest,
		encodeResponse,
	))

	router.Methods("DELETE").Path("/rms/master_languages/{id}").Handler(httptransport.NewServer(
		makeDeleteLanguageEndpoint(svc),
		decodeDeleteLanguageRequest,
		encodeResponse,
	))

	return router
}

func decodeCreateLanguageRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req createLanguageRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return nil, err
    }
    return req, nil
}


func decodeEmptyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// func decodeGetLanguageByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var req getLanguageByIDRequest
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.ID = id
// 	return req, nil
// }

func decodeGetLanguagesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getLanguagesRequest
	vars := r.URL.Query()
	if id, ok := vars["id"]; ok && len(id) > 0 {
		id, err := strconv.Atoi(id[0])
		if err != nil {
			return nil, err
		}
		req.ID = &id
	}
	return req, nil
}

func decodeUpdateLanguageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req updateLanguageRequest
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	req.ID = id
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteLanguageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req deleteLanguageRequest
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	req.ID = id
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
