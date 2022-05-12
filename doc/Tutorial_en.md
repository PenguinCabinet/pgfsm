# Tutorial
This tutorial is learning how to use The Pen Game Programing Finite State Machine.

# The development environment

```shell
>go version
go version go1.17.5 windows/amd64
```

# First
This library is affected by [How to Build a JRPG: A Primer for Game Developers](https://gamedevelopment.tutsplus.com/articles/how-to-build-a-jrpg-a-primer-for-game-developers--gamedev-6676)  
This is written about game state machine clearly.  
You should read it before you read this tutorial and understand about game state machine roughly.

# How to make your project
First,You need it.  
How to make your project is general. 

```shell
mkdir tutorial
cd tutorial
go mod init tutorial
```

# Install needed libraies

This library is a dependency with Ebiten, so when you put in The Pen Game Programing Finite State Machine, Ebiten will be downloaded with it.
```shell
go get github.com/PenguinCabinet/pgfsm
```

Since this tutorial will also use text display, font-related libraries should also be included.
```shell
go get golang.org/x/image/font
go get golang.org/x/image/font/opentype
```

# Make the title scene

First, let's create the title scene.  
The title scene and game scenes are created as states.  

```go
package main

import (
	"image/color"
	"log"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//This is the title scene state
type TitleGameState struct {
	mplusNormalFont font.Face
}

//This is the function that is called when the state is first executed.
func (sm *TitleGameState) Init(
	stackdeep int, /*Here is the index of where this state is stacked on the stack*/
	delta float64, /*Here is the time that has elapsed between the previous frame and the current frame.*/
) {

	/*Here is the start of the font initialization process of Ebiten*/
	const dpi = 72

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		panic(err)
	}

	sm.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
	/*Here is the end of the font initialization process of Ebiten*/
}

//This is the function that is called every frame.
//Called only when this state is running.
func (sm *TitleGameState) Update(
	screen *ebiten.Image, /*Screen of ebiten, but it is deprecated to describe it in Update*/
	stackdeep int, delta float64,
) pgfsm.Result {
	/*Continue loop by returning an empty pgfsm.Result.
	Change the current running state to the new state by rewriting the returned pgfsm.Result or
	New states can be placed on top of the stack.*/
	return pgfsm.Result{}
}

//This is the function for drawing that is called every frame.
//Even if this state is not running, it will be called if it is on the stack.
func (sm *TitleGameState) Draw(screen *ebiten.Image, stackdeep int, delta float64) {
	text.Draw(screen, "Game Title", sm.mplusNormalFont, 200, 100, color.White)
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pen_Game_State_Machine")

	gms := &pgfsm.Machine{}

	gms.LayoutWidth = 640
	gms.LayoutHeight = 480

	TitleSm := &TitleGameState{}

	/*Add the title scene state to the stack*/
	gms.StateAdd(TitleSm)

	if err := ebiten.RunGame(gms); err != nil {
		log.Fatal(err)
	}
}
```
The execution Result   
![img1](image/img1.png)

Please watch comment in the source code. 
State is specified in an interface called pgfsm.State, which is implemented accordingly.
In this case, the state of the title scene is implemented as TitleGameState.
```go
type State interface {
	Init(int, float64)
	Update(*ebiten.Image, int, float64) Result
	Draw(*ebiten.Image, int, float64)
}
```

Also
```go
    return pgfsm.Result{}
```
Please note this.   
By changing the return value of this Update, you can switch to a new state or put a new state on the stack. 

# The game scene and switching between game scenes.
The title scene has been completed.   
Next, let's implement the title scene and switching between the title scene and the game scene.  
Entering the s key on the title scene,switch from the title scene to the game scene.

```go
package main

import (
	"image/color"
	"log"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//This is the game scene state.
type GameMainState struct {
	mplusNormalFont font.Face
}

//This is the function that is called when the state is first executed.
func (sm *GameMainState) Init(
	stackdeep int, /*Here is the index of where this state is stacked on the stack.*/
	delta float64, /*Here is the time that has elapsed between the previous frame and the current frame.*/
) {
	/*Here is the start of the font initialization process of Ebiten*/
	const dpi = 72

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		panic(err)
	}

	sm.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
	/*Here is the end of the font initialization process of Ebiten*/
}

//This is the function that is called every frame.
//Called only when this state is running.
func (sm *GameMainState) Update(
	screen *ebiten.Image, /*Screen of ebiten, but it is deprecated to describe it in Update*/
	stackdeep int, delta float64,
) pgfsm.Result {
	/*Continue loop by returning an empty pgfsm.Result.
	Change the current running state to the new state by rewriting the returned pgfsm.Result or
	New states can be placed on top of the stack.*/
	return pgfsm.Result{}
}

//This is the function for drawing that is called every frame.
//Even if this state is not running, it will be called if it is on the stack.
func (sm *GameMainState) Draw(screen *ebiten.Image, stackdeep int, delta float64) {
	text.Draw(screen, "Game Main", sm.mplusNormalFont, 200, 100, color.White)
}

//This is the title scene state
type TitleGameState struct {
	mplusNormalFont font.Face
}

//This is the function that is called when the state is first executed.
func (sm *TitleGameState) Init(
	stackdeep int, /*Here is the index of where this state is stacked on the stack*/
	delta float64, /*Here is the time that has elapsed between the previous frame and the current frame.*/
) {
	/*Here is the start of the font initialization process of Ebiten*/
	const dpi = 72

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		panic(err)
	}

	sm.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
	/*Here is the end of the font initialization process of Ebiten*/
}

//This is the function that is called every frame.
//Called only when this state is running.
func (sm *TitleGameState) Update(
	screen *ebiten.Image, /*Screen of ebiten, but it is deprecated to describe it in Update*/
	stackdeep int, delta float64,
) pgfsm.Result {

	/*If the s key is entered*/
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		/*Change the state here
		pgfsm.CodeChange changes the currently running state to NextState.
		Here you changes the currently running title scene state to the game scene state.*/
		return pgfsm.Result{
			Code:      pgfsm.CodeChange,
			NextState: &GameMainState{},
		}
	}
	/*Continue loop by returning an empty pgfsm.Result.
	Change the current running state to the new state by rewriting the returned pgfsm.Result or
	New states can be placed on top of the stack.*/
	return pgfsm.Result{}
}

//This is the function for drawing that is called every frame.
//Even if this state is not running, it will be called if it is on the stack.
func (sm *TitleGameState) Draw(screen *ebiten.Image, stackdeep int, delta float64) {
	text.Draw(screen, "Game Title\nPressing S key,start!", sm.mplusNormalFont, 100, 100, color.White)
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pen_Game_State_Machine")

	gms := &pgfsm.Machine{}

	gms.LayoutWidth = 640
	gms.LayoutHeight = 480

	TitleSm := &TitleGameState{}

	/*Add the title scene state to the stack*/
	gms.StateAdd(TitleSm)

	if err := ebiten.RunGame(gms); err != nil {
		log.Fatal(err)
	}
}
```
![img1](image/img2.gif)  
Press the s key to switch from the title scene to the game scene! (The gif file is set to loop, so it looks like it also switch from the game scene to the title scene, but it does not.)
```go
		return pgfsm.Result{
			Code:       pgfsm.CodeChange,
			NextState: &GameMainState{},
		}
```
Here is essential,It is possible to switch states by changing the returned pgfsm.Result.

# Implementation of the menu scene.
The title scene and the game scene have been completed.  
Next, let's implement putting the map scene and map scene on the stack.  
Pressing the m key during the game scene opens the menu scene, and pressing the m key when the menu scene is open closes it.  

Why can't we simply switch from the game scene to the menu scene?    
That is because the game scene data is retained while the menu scene is open and   
When the menu scene is closed, it is necessary to return to the game scene before opening.    

```go
package main

import (
	"image/color"
	"log"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//This is the menu scene state
type MenuGameState struct {
	mplusNormalFont font.Face
}

//This is the function that is called when the state is first executed.
func (sm *MenuGameState) Init(
	stackdeep int, /*Here is the index of where this state is stacked on the stack*/
	delta float64, /*Here is the time that has elapsed between the previous frame and the current frame.*/
) {
	/*Here is the start of the font initialization process of Ebiten*/
	const dpi = 72

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		panic(err)
	}

	sm.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
	/*Here is the end of the font initialization process of Ebiten*/
}

//This is the function that is called every frame.
//Called only when this state is running.
func (sm *MenuGameState) Update(
	screen *ebiten.Image, /*Screen of ebiten, but it is deprecated to describe it in Update*/
	stackdeep int, delta float64,
) pgfsm.Result {

	/*If m key is entered,the menu is closed*/
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		/*
			Here,delete the current running game scene state.
			The stack is stocked in the order of "the game scene,the menu scene",
			so when you delete the current running game scene state, the contents of the stack will be
			"the game scene" and moves back to the menu scene process.
		*/
		return pgfsm.Result{
			Code:      pgfsm.CodeDelete,
			NextState: nil,
		}
	}
	/*Continue loop by returning an empty pgfsm.Result.
	Change the current running state to the new state by rewriting the returned pgfsm.Result or
	New states can be placed on top of the stack.*/
	return pgfsm.Result{}
}

//This is the function for drawing that is called every frame.
//Even if this state is not running, it will be called if it is on the stack.
func (sm *MenuGameState) Draw(screen *ebiten.Image, stackdeep int, delta float64) {
	text.Draw(screen, "Menu", sm.mplusNormalFont, 300, 240, color.White)
}

//This is the game scene state
type GameMainState struct {
	mplusNormalFont font.Face
}

//This is the function that is called when the state is first executed.
func (sm *GameMainState) Init(
	stackdeep int, /*Here is the index of where this state is stacked on the stack*/
	delta float64, /*Here is the time that has elapsed between the previous frame and the current frame.*/
) {
	/*Here is the start of the font initialization process of Ebiten*/
	const dpi = 72

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		panic(err)
	}

	sm.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
	/*Here is the end of the font initialization process of Ebiten*/
}

//This is the function that is called every frame.
//Called only when this state is running.
//In other words,during opening the menu,The Update function is not being run.
func (sm *GameMainState) Update(
	screen *ebiten.Image, /*Screen of ebiten, but it is deprecated to describe it in Update*/
	stackdeep int, delta float64,
) pgfsm.Result {
	/*If m key is entered,the menu is opened*/
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		/*Here,put the menu state on the top of the current running game scene state.
		The stack is stocked in the order of "game scene",
		so when you add the menu scene, the contents of the stack will be
		"the game scene,the menu scene" and moves to the menu scene process.
		*/
		return pgfsm.Result{
			Code:      pgfsm.CodeAdd,
			NextState: &MenuGameState{},
		}
	}

	/*Continue loop by returning an empty pgfsm.Result.
	Change the current running state to the new state by rewriting the returned pgfsm.Result or
	New states can be placed on top of the stack.*/
	return pgfsm.Result{}
}

//This is the function for drawing that is called every frame.
//Even if this state is not running, it will be called if it is on the stack.
//In other words,during opening the menu,The Drawe function is being run.
func (sm *GameMainState) Draw(screen *ebiten.Image, stackdeep int, delta float64) {
	text.Draw(screen, "Game Main", sm.mplusNormalFont, 200, 100, color.White)
}

//This is the title scene state
type TitleGameState struct {
	mplusNormalFont font.Face
}

//This is the function that is called when the state is first executed.
func (sm *TitleGameState) Init(
	stackdeep int, /*Here is the index of where this state is stacked on the stack*/
	delta float64, /*Here is the time that has elapsed between the previous frame and the current frame.*/
) {
	/*Here is the start of the font initialization process of Ebiten*/
	const dpi = 72

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		panic(err)
	}

	sm.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
	/*Here is the end of the font initialization process of Ebiten*/
}

//This is the function that is called every frame.
//Called only when this state is running.
func (sm *TitleGameState) Update(
	screen *ebiten.Image, /*Screen of ebiten, but it is deprecated to describe it in Update*/
	stackdeep int, delta float64,
) pgfsm.Result {

	/*If the s key is entered*/
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		/*Change the state here
		pgfsm.CodeChange changes the currently running state to NextState.
		Here you changes the currently running title scene state to the game scene state.*/
		return pgfsm.Result{
			Code:      pgfsm.CodeChange,
			NextState: &GameMainState{},
		}
	}
	/*Continue loop by returning an empty pgfsm.Result.
	Change the current running state to the new state by rewriting the returned pgfsm.Result or
	New states can be placed on top of the stack.*/
	return pgfsm.Result{}
}

//This is the function for drawing that is called every frame.
//Even if this state is not running, it will be called if it is on the stack.
func (sm *TitleGameState) Draw(screen *ebiten.Image, stackdeep int, delta float64) {
	text.Draw(screen, "Game Title\nPressing S key,start!", sm.mplusNormalFont, 100, 100, color.White)
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pen_Game_State_Machine")

	gms := &pgfsm.Machine{}

	gms.LayoutWidth = 640
	gms.LayoutHeight = 480

	TitleSm := &TitleGameState{}

	/*Add the title scene state to the stack*/
	gms.StateAdd(TitleSm)

	if err := ebiten.RunGame(gms); err != nil {
		log.Fatal(err)
	}
}

```
![img1](image/img3.gif)  
Pressing the m key on the game scene,The menu scene is opened.
  
While the menu is open, the Update function of the game scene **will not execute**, the Draw function of the game scene **will execute**, and the variables of the game scene **will continue to be held**. (In other words, the data is retained and displayed but stopped.)   
The way they are stacked rather than switched is ideal for JRPG combat, the menu scene displays, etc. 


