
package server

import (
	"fmt"
	"github.com/JuniorDT/opendata-searcher-data-service/pkg/server/routers"
	"github.com/JuniorDT/opendata-searcher-data-service/pkg/services_test_results/common_service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(u commontestresults.FileParseResultService) *Server {
	s := Server{router: mux.NewRouter()}
	server.NewTestResultRouter(u, s.newSubRouter("/test_results"))
	// check routers
	s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})
	return &s
}

func (s *Server) Start() {
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) newSubRouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}