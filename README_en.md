[README of JAPANESE](./README.md)
# The Pen Game Programing Finite State Machine

This is the state machine library for [Ebiten](https://ebiten.org/).
## Provide features
* The state and stack machine for the game programming
* The deltatime between the previous frame and the current frame.

# Documents
* [Tutorial](doc/Tutorial_en.md) 
* [pkg.go.dev](https://pkg.go.dev/github.com/PenguinCabinet/pgfsm)
* [Examples](examples/)

# Used by

We are currently accepting games.   
Could you tell me your games using this by issue.

# Quick start

## Install
```shell
go get github.com/PenguinCabinet/pgfsm
```

## Example
```go
package main

import (
	"log"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten"
)

type MyGameState struct {
}

func (sm *MyGameState) Init(
	stackdeep int, /*The index of this state.*/
) {
	//Init
}

func (sm *MyGameState) Update(
	stackdeep int,
) pgfsm.Result {
	//Update
	return pgfsm.Result{
		Code:      pgfsm.CodeNil,
		NextState: nil,
	}
}

func (sm *MyGameState) Draw(screen *ebiten.Image, stackdeep int) {
	//Draw
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("game title")

	gms := &pgfsm.Machine{}

	gms.LayoutWidth = 640
	gms.LayoutHeight = 480

	mySm := &MyGameState{}

	gms.StateAdd(mySm)

	if err := ebiten.RunGame(gms); err != nil {
		log.Fatal(err)
	}
}
```
