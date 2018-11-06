package main

import (
	"github.com/JuniorDT/opendata-searcher-data-service/pkg/mongo"
	"github.com/JuniorDT/opendata-searcher-data-service/pkg/server"
	"log"
)

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}
	defer ms.Close()

	u := mongo.NewTestResultService(ms.Copy(), "test_results", "parse_results")
	s := server.NewServer(u)

	s.Start()
}