package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Function To handle PUT Request "/maze/solve"
func MazeSolver(w http.ResponseWriter, r *http.Request) {

	//Reading the input from the JSON
	decoder := json.NewDecoder(r.Body)
	var Path map[string]interface{}
	err := decoder.Decode(&Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//We are currently in the starting room, not inside the given maze yet
	currentPath := make([]string, 0)
	pathFound := false
	ans := mazeTraverer(Path, currentPath, &pathFound)

	var ansString string

	if len(ans) > 0 {
		//If there is a Path
		ansString = "[" + strings.Join(ans, ", ") + "]"
	} else {
		//If no paths were were found
		ansString = "Sorry"
	}
	w.Write([]byte(ansString))
}

// This Function Will be called Recursively
// We are traversing the maze and all the paths possible
// As Soon as we find our first path we are returning it, it may not be the shortest path
func mazeTraverer(Path interface{}, currentPath []string, pathFound *bool) []string {

	//Converting it for map traversal
	//The PathMap variable is a Map Data Type where
	//Key is the direction, Value is the part of the maze in the choosen direction
	PathMap, _ := Path.(map[string]interface{})

	//We Move in the available directions from the current room we are in
	for direction, nextPath := range PathMap {
		switch nextPath.(type) {
		case string:
			//If the value is a string, this means we can not got to a new room from this room
			//Its either the Exit, Or its a Deamon/Animal or a dead end
			if nextPath == "exit" {
				//We have foudn our path and successfully
				//Returning the ans
				*pathFound = true
				currentPath = append(currentPath, direction)
				return currentPath
			} else {
				//If its not an exit, its a Deamon or Dead End
				//We have no choice to go from here, so we discard this path
				continue
			}

		case map[string]interface{}:
			//If the value is a Map, it means we can move to a new room in this direction
			//We add the direction we move in to our current path
			pathTaken := append(currentPath, direction)

			//We call the recursive funtion to move ahead and complete the path
			pathRecieved := mazeTraverer(nextPath, pathTaken, pathFound)
			//pathRecieved has the full path we took
			if *pathFound {
				//If we have found a exit we return the path
				return pathRecieved
			}
		}

	}

	//If we reach here, we could not find the path
	return (make([]string, 0))
}
