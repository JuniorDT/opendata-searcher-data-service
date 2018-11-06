package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JuniorDT/opendata-searcher-data-service/pkg/tests/common_service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type commonTestResultRouter struct {
	CSVParseResultService commontestresults.CSVParseResultService
}

func NewTestResultRouter(r commontestresults.CSVParseResultService, router *mux.Router) *mux.Router {
	CommonTestResultRouter := commonTestResultRouter{CSVParseResultService: r}

	router.HandleFunc("/parse_result/", CommonTestResultRouter.createCSVParseResultHandler).Methods("PUT")
	router.HandleFunc("/parse_result/{id}", CommonTestResultRouter.getCSVParseResultHandler).Methods("GET")
	return router
}

func(trr *commonTestResultRouter) createCSVParseResultHandler(w http.ResponseWriter, r *http.Request) {
	result, err := decodeCSVParseResult(r)
	if err != nil {
		fmt.Println(err)
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = trr.CSVParseResultService.Create(&result)
	if err != nil {
		fmt.Println(err)
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, http.StatusOK, err)
}

func(trr *commonTestResultRouter) getCSVParseResultHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	log.Println(vars)
	id := vars["id"]

	result, err := trr.CSVParseResultService.GetById(id)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, result)
}

func decodeCSVParseResult(rr *http.Request) (r commontestresults.CSVParseResult, err error) {

	if rr.Body == nil {
		return r, errors.New("no request body")
	}
	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&r)
	return r, err
}

func Error(w http.ResponseWriter, code int, message string) {
	Json(w, code, map[string]string{"error": message})
}

func Json(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}