package controllers

import (
	"encoding/json"
	"fmt"
)

// So first I solved on a test case where I  JSONs value can be only string and there exsists a exit
// Below is the example
// I just traverse over the map and find the exit
func MazeSolver01() {
	s := `{"forward": "exit","upstairs": "tiger"}`
	data := []byte(s)
	var Path map[string]string
	json.Unmarshal(data, &Path)

	currentPath := make([]string, 0)

	for direction, pathDescription := range Path {
		if pathDescription == "exit" {
			currentPath = append(currentPath, direction)
			break
		}
	}

	fmt.Println(currentPath)

}

// After that I solve for a test which works with the above condition and Where there are no exits
// Handaling the condtion where we need to print Sorry
func MazeSolver02() {
	s := `{"forward": "exit","upstairs": "tiger"}`
	data := []byte(s)
	var Path map[string]string
	json.Unmarshal(data, &Path)

	currentPath := make([]string, 0)

	for direction, pathDescription := range Path {
		if pathDescription == "exit" {
			currentPath = append(currentPath, direction)
			break
		}
	}

	if len(currentPath) == 0 {
		fmt.Println("Sorry")
	} else {
		fmt.Println(currentPath)
	}

}

// Here Solving the condition
// Splitted the Function here - First PreProcessing the Input and then finally calling the function to handle the cases
// Here I handle the nested json where Values can be map or a string
// Used Interface{} as we are not sure whether its string or map
func MazeSolver03() {

	s := `{"forward": "tiger", "left": {"forward": {"upstairs": "exit"}, "left": "dragon"}, "right": {"forward": "dead end"}}`
	data := []byte(s)
	var Path map[string]interface{}
	json.Unmarshal(data, &Path)
	currentPath := make([]string, 0)
	pathFound := false
	ans := MazeSolverfinal(Path, currentPath, &pathFound)
	if len(ans) == 0 {
		fmt.Println("Sorry")
	} else {
		fmt.Println(ans)
	}

}

// Handlaing the condition to check if its a map or string using switch case
func MazeSolverfinal(Path interface{}, currentPath []string, pathFound *bool) []string {
	PathMap, _ := Path.(map[string]interface{})
	for direction, nextPath := range PathMap {
		switch nextPath.(type) {
		case string:
			if nextPath == "exit" {
				*pathFound = true
				currentPath = append(currentPath, direction)
				return currentPath
			} else {
				continue
			}
		case map[string]interface{}:
			pathTaken := append(currentPath, direction)
			pathRecieved := MazeSolverfinal(nextPath, pathTaken, pathFound)
			if *pathFound {
				return pathRecieved
			}
		}

	}
	return (make([]string, 0))
}

// I have assumed that we only need to return a single path even there exsits multiple path
// The previous one I did already take care of the condition
// If we need multiple path we can have a slice where we store the all possible paths and not end the recurrsion just as we found 1st path
func MazeSolver04() {

	s := `{"forward": "tiger", "left": {"forward": {"upstairs": "exit"}, "left": "exit"}, "right": {"forward": "dead end"}}`
	data := []byte(s)
	var Path map[string]interface{}
	json.Unmarshal(data, &Path)
	currentPath := make([]string, 0)
	pathFound := false
	ans := MazeSolverfinal(Path, currentPath, &pathFound)
	if len(ans) == 0 {
		fmt.Println("Sorry")
	} else {
		fmt.Println(ans)
	}

}
