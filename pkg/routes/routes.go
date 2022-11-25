package routes

import (
	"github.com/gorilla/mux"
	"github.com/sankalpneel/pxtest/pkg/controllers"
)

var Maze = func(router *mux.Router) {

	//Maze Solver
	router.HandleFunc("/maze/solve", controllers.MazeSolver).Methods("POST")

	//Maze Generator
	router.HandleFunc("/maze/generate", controllers.MazeGenerator).Methods("GET")
}
