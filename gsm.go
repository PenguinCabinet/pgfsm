//State Machine
package pgfsm

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Code int

const (
	CodeNil                Code = iota //何もせずそのまま処理実行
	CodeChange                         //現在実行しているStateとnext_Stateを入れ替える
	CodeAdd                            //現在実行しているStateの上にnext_Stateを乗せる
	CodeDelete                         //現在実行しているStateを消去する
	CodeAllDeleteAndChange             //現在スタックに乗っているStateをすべて消去して、next_Stateを乗せる
	CodeInsertBack                     //現在実行しているStateの後ろに、next_Stateを挿入する
	/*
		GameState_result_delete_and_insert_back                                 //現在実行しているStateの後ろに、next_Stateを挿入する
		GameState_result_insert2_back
	*/
)

/*
var GameState_result_map map[string]GameState_result_code_t = map[string]GameState_result_code_t{
	"GameState_result_nil":                   GameState_result_nil,
	"GameState_result_change":                GameState_result_change,
	"GameState_result_add":                   GameState_result_add,
	"GameState_result_delete":                GameState_result_delete,
	"GameState_result_all_delete_and_change": GameState_result_all_delete_and_change,
	"GameState_result_insert_back":           GameState_result_insert_back,
	/*
		"GameState_result_delete_and_insert_back": GameState_result_delete_and_insert_back,
		"GameState_result_insert2_back":           GameState_result_insert2_back,
}
*/

type Result struct {
	Code      Code
	NextState State
}

type State interface {
	Init(int, float64)
	Update(*ebiten.Image, int, float64) Result
	Draw(*ebiten.Image, int, float64)
}

type Machine struct {
	StateStack                []State
	deltaTime                 float64
	oldTimeForDeltaTime       int64
	LayoutWidth, LayoutHeight int
}

func (g *Machine) Update(screen *ebiten.Image) error {
	newTimeForDeltaTime := time.Now().UnixNano()
	g.deltaTime = float64(newTimeForDeltaTime-g.oldTimeForDeltaTime) * 0.001
	g.oldTimeForDeltaTime = newTimeForDeltaTime

	if g.Empty() {
		return nil
	}

	res := g.StateStack[len(g.StateStack)-1].Update(screen, len(g.StateStack)-1, g.deltaTime)

	switch res.Code {
	case CodeChange:
		g.StateChange(res.NextState)
	case CodeAdd:
		g.StateAdd(res.NextState)
	case CodeDelete:
		g.StateStack = g.StateStack[0 : len(g.StateStack)-1]
	case CodeAllDeleteAndChange:
		g.StateStack = []State{}
		g.StateAdd(res.NextState)
	case CodeInsertBack:
		res.NextState.Init(len(g.StateStack)-1, g.deltaTime)
		index := len(g.StateStack) - 2
		g.StateStack = append(g.StateStack[:index], g.StateStack[index:]...)
		g.StateStack[index] = res.NextState
		/*
			case State_result_delete_and_insert_back:
				g.State = g.State[0 : len(g.State)-1]
				index := len(g.State) - 1
				res.Next_State.Init()
				g.State = append(g.State[:index+1], g.State[index:]...)
				g.State[index] = res.Next_State

			case State_result_insert2_back:
				res.Next_State.Init()
				index := len(g.State) - 2
				g.State = append(g.State[:index], g.State[index:]...)
				g.State[index] = res.Next_State
		*/

	}

	return nil
}

func (g *Machine) Draw(screen *ebiten.Image) {
	for d, e := range g.StateStack {
		e.Draw(screen, d, g.deltaTime)
	}
	//ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Machine) Empty() bool {
	return len(g.StateStack) == 0
}

func (g *Machine) StateAdd(v State) {
	g.StateStack = append(g.StateStack, v)
	v.Init(len(g.StateStack)-1, g.deltaTime)
}

func (g *Machine) StateChange(v State) {
	g.StateStack[len(g.StateStack)-1] = v
	v.Init(len(g.StateStack)-1, g.deltaTime)
}

func (g *Machine) DebugLogprint() {
	for i, e := range g.StateStack {
		fmt.Printf("%d %+v\n", i, e)
	}
	fmt.Printf("\n")
}

func (g *Machine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.LayoutWidth, g.LayoutHeight
}

func (g *Machine) Init() {
	g.oldTimeForDeltaTime = time.Now().UnixNano()
}
