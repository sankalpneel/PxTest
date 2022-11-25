# PxTest

To run the app (port 5000)
```
go run main.go
```
To run the tests
```
go test ./...
```

Maze Solver
- pkg\controllers\mazeSolver.go  => contains the final code for maze solver
- pkg\controllers\recfactorHistory.go => contians the code and different functions for code refactor and evolution
- pkg\controllers\mazeGenerator.go => contains the code for Random Maze generation

### I have added some comments in the code to explain my approach
- POST /maze/solve - returns a way to exit the maze if possible, else Sorry
Used Depth First Search
Parameter in JSON containg the maze

- GET /maze/generate - Genrates a random maze and returns 
 Parameter: max-children - the maximum number of children per node



## Explain in one or two paragraphs how to generate an infinity of maze. Think of exceptions and weird cases.

We can consider the maze as a tree. A node has a n children. The edges are directed from a parent node to child node. There are no loops.

A room is a Vertex - It might be empty or contail a animal\monster or dead end or exit

i) If a room is empty we can move in the room and then further to its childern.
ii) If a room has an animal/monster or dead end. It wont be having child nodes, if it does have its of no use because we can not go into this room.
iii) If the node is exit then also it wont be having any child nodes that are useable/necessary

Now on creating the maze -> initially we will be in an empty room. We can move using the edges connected to it.
It will be having max n childern and thus max n edges.

To create the maze we will first select the n direction the room has opening to. Now we will define what each direction room will have.
We can select any random description for the room from all the options - monsted/animal/dead end/exit. We can change the probablity if needed.
What can happen is the maze can be created infinitely, to prevent it we can add a limit to number of rooms.

If the room is anyting but an empty room we can just the animal or dragon to it, and dont create its child nodes.
If its a empty room we can call a recursvie function as everything will be the same. We can consider this room as starting room.
