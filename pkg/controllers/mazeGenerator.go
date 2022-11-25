package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
##Explain in one or two paragraphs how to generate an infinity of maze. Think of exceptions and weird cases.

We can consider the maze as a tree. A node has a n children. The edges are directed from a parent node to child node.
There are no loops.

A room is a Vertex - It might be empty or contail a animal\monster or dead end or exit

i) If a room is empty we can move in the room and then further to its childern.
ii) If a room has an animal/monster or dead end. It wont be having child nodes, if it does have its of no use because we can not
	go into this room.
iii) If the node is exit then also it wont be having any child nodes that are useable/necessary

Now on creating the maze -> initially we will be in an empty room. We can move using the edges connected to it.
It will be having max n childern and thus max n edges.

To create the maze we will first select the n direction the room has opening to. Now we will define what each direction room will have.
We can select any random description for the room from all the options - monsted/animal/dead end/exit. We can change the probablity if needed.
What can happen is the maze can be created infinitely, to prevent it we can add a limit to number of rooms.

If the room is anyting but an empty room we can just the animal or dragon to it, and dont create its child nodes.
If its a empty room we can call a recursvie function as everything will be the same. We can consider this room as starting room.

*/

// Function To handle GET Request "/maze/generate?max-children=<int>"
func MazeGenerator(w http.ResponseWriter, r *http.Request) {

	//Reading the parameter
	maxChildren, err := strconv.Atoi(r.URL.Query().Get("max-children"))
	if err != nil || maxChildren > 5 {
		http.Error(w, "Hi, Please select max childen upto 5", http.StatusBadRequest)
		return
	}
	maxRooms := 10   //To stop it from going into Infinite loop
	roomCreated := 0 //calculate count of room
	Maze := roomCreator(maxChildren, maxRooms, &roomCreated)
	fmt.Fprint(w, Maze)
}

// Recursive Function to generate The Maze
// This function is called recusively when we one room leads to another empty room and does not contain animal,dead end or exit
func roomCreator(maxChildren, maxRooms int, roomsCreated *int) string {
	Maze := make(map[string]string)

	//we choose the N children room directions
	directions := chooseRandomNDirection(maxChildren)

	for i := 0; i < maxChildren; i++ {

		//to prevent an infinite loop I have put this here
		if maxRooms == *roomsCreated {
			break
		}
		*roomsCreated++

		//We get the room description for the ith direction
		//It can be a animal, monster, dead end, exit or another empty room we can move to
		roomDescription := chooseRandomNRoomDescriptions()

		switch roomDescription {
		case "room":
			//We Need to Create a new room in the ith direction
			//Handling a case where if this the last room we need to create, we can not add a new room in this direction so just added dead end
			if *roomsCreated == maxRooms {
				Maze[directions[i]] = "dead end"
				continue
			}

			//We call the roomCreator function recursively to create a new room in this direction
			newRoom := roomCreator(maxChildren, maxRooms, roomsCreated)
			Maze[directions[i]] = newRoom
		default:
			//Otherwise we just add whatever the description we get
			Maze[directions[i]] = roomDescription
		}

	}

	//We change the Map to string so we can send it in the response
	//Using the predefined function lead to problems with esape charecters of the string
	roomDesription := roomsMapToStringParser(Maze)
	return roomDesription

}

func roomsMapToStringParser(Maze map[string]string) string {
	var roomSlice []string
	for key, val := range Maze {
		var s string
		if val[0] == '{' {
			s = fmt.Sprintf("\"%s\": %s", key, val)
		} else {
			s = fmt.Sprintf("\"%s\": \"%s\"", key, val)
		}

		roomSlice = append(roomSlice, s)
	}
	var parsedString string = "{" + strings.Join(roomSlice, ", ") + "}"

	return string(parsedString)
}

//Random Fucntions Can be improved

// Returns back n random directions
func chooseRandomNDirection(n int) []string {
	directions := [5]string{"left", "right", "upstairs", "downstairs", "forward"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })
	return directions[:n]

}

// returns back a random room description
func chooseRandomNRoomDescriptions() string {
	RoomDes := [6]string{"tiger", "ogre", "room", "dragon", "exit", "dead end"}
	randomIndex := rand.Intn(len(RoomDes))
	pick := RoomDes[randomIndex]
	return pick
}
