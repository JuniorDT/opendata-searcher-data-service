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
	FileParseResultService commontestresults.FileParseResultService
}

func NewTestResultRouter(r commontestresults.FileParseResultService, router *mux.Router) *mux.Router {
	CommonTestResultRouter := commonTestResultRouter{FileParseResultService: r}

	router.HandleFunc("/parse_result/", CommonTestResultRouter.createFileParseResultHandler).Methods("PUT")
	router.HandleFunc("/parse_result/{id}", CommonTestResultRouter.getFileParseResultHandler).Methods("GET")
	return router
}

func(trr *commonTestResultRouter) createFileParseResultHandler(w http.ResponseWriter, r *http.Request) {
	result, err := decodeFileParseResult(r)
	if err != nil {
		fmt.Println(err)
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = trr.FileParseResultService.Create(&result)
	if err != nil {
		fmt.Println(err)
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	Json(w, http.StatusOK, err)
}

func(trr *commonTestResultRouter) getFileParseResultHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	log.Println(vars)
	id := vars["id"]

	result, err := trr.FileParseResultService.GetById(id)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, result)
}

func decodeFileParseResult(rr *http.Request) (r commontestresults.FileParseResult, err error) {

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