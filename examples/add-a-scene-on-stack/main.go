package main

import (
	"image/color"
	"log"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
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
	stackdeep int,
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
func (sm *MenuGameState) Draw(screen *ebiten.Image, stackdeep int) {
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
