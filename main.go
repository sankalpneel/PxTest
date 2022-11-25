package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sankalpneel/pxtest/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.Maze(r)
	http.Handle("/", r)
	port := ":5000"
	log.Println("Running on port:", port)
	log.Fatal(http.ListenAndServe(port, r))

	/*
		These functions are located in pkg\controllers\recfactorHistory.go
		Its there as per request, as you wanted to see the evolution and refactors
		controllers.MazeSolver01()
		controllers.MazeSolver02()
		controllers.MazeSolver03()
		controllers.MazeSolver04()
	*/

	/*
		To run the test => go test ./...
	*/
}
