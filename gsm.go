//State Machine
package Pen_Game_State_Machine

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Game_State_result_code_t int

const (
	Game_State_result_nil                   Game_State_result_code_t = iota //何もせずそのまま処理実行
	Game_State_result_change                                                //現在実行しているStateとnext_Stateを入れ替える
	Game_State_result_add                                                   //現在実行しているStateの上にnext_Stateを乗せる
	Game_State_result_delete                                                //現在実行しているStateを消去する
	Game_State_result_all_delete_and_change                                 //現在スタックに乗っているStateをすべて消去して、next_Stateを乗せる
	Game_State_result_insert_back                                           //現在実行しているStateの後ろに、next_Stateを挿入する
	/*
		Game_State_result_delete_and_insert_back                                 //現在実行しているStateの後ろに、next_Stateを挿入する
		Game_State_result_insert2_back
	*/
)

var Game_State_result_map map[string]Game_State_result_code_t = map[string]Game_State_result_code_t{
	"Game_State_result_nil":                   Game_State_result_nil,
	"Game_State_result_change":                Game_State_result_change,
	"Game_State_result_add":                   Game_State_result_add,
	"Game_State_result_delete":                Game_State_result_delete,
	"Game_State_result_all_delete_and_change": Game_State_result_all_delete_and_change,
	"Game_State_result_insert_back":           Game_State_result_insert_back,
	/*
		"Game_State_result_delete_and_insert_back": Game_State_result_delete_and_insert_back,
		"Game_State_result_insert2_back":           Game_State_result_insert2_back,
	*/
}

type Game_State_result_t struct {
	Code       Game_State_result_code_t
	Next_State Game_State_t
}

type Game_State_t interface {
	Init(int, float64)
	Update(*ebiten.Image, int, float64) Game_State_result_t
	Draw(*ebiten.Image, int, float64)
}

type Game_State_Machine_t struct {
	Game_State                  []Game_State_t
	deltaTime                   float64
	old_time_for_deltaTime      int64
	Layout_Width, Layout_Height int
}

func (g *Game_State_Machine_t) Update(screen *ebiten.Image) error {
	new_time_for_deltaTime := time.Now().UnixNano()
	g.deltaTime = float64(new_time_for_deltaTime-g.old_time_for_deltaTime) * 0.001
	g.old_time_for_deltaTime = new_time_for_deltaTime

	if g.Empty() {
		return nil
	}

	res := g.Game_State[len(g.Game_State)-1].Update(screen, len(g.Game_State)-1, g.deltaTime)

	switch res.Code {
	case Game_State_result_change:
		g.State_Change(res.Next_State)
	case Game_State_result_add:
		g.State_Add(res.Next_State)
	case Game_State_result_delete:
		g.Game_State = g.Game_State[0 : len(g.Game_State)-1]
	case Game_State_result_all_delete_and_change:
		g.Game_State = []Game_State_t{}
		g.State_Add(res.Next_State)
	case Game_State_result_insert_back:
		res.Next_State.Init(len(g.Game_State)-1, g.deltaTime)
		index := len(g.Game_State) - 2
		g.Game_State = append(g.Game_State[:index], g.Game_State[index:]...)
		g.Game_State[index] = res.Next_State
		/*
			case Game_State_result_delete_and_insert_back:
				g.Game_State = g.Game_State[0 : len(g.Game_State)-1]
				index := len(g.Game_State) - 1
				res.Next_State.Init()
				g.Game_State = append(g.Game_State[:index+1], g.Game_State[index:]...)
				g.Game_State[index] = res.Next_State

			case Game_State_result_insert2_back:
				res.Next_State.Init()
				index := len(g.Game_State) - 2
				g.Game_State = append(g.Game_State[:index], g.Game_State[index:]...)
				g.Game_State[index] = res.Next_State
		*/

	}

	return nil
}

func (g *Game_State_Machine_t) Draw(screen *ebiten.Image) {
	for d, e := range g.Game_State {
		e.Draw(screen, d, g.deltaTime)
	}
	//ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game_State_Machine_t) Empty() bool {
	return len(g.Game_State) == 0
}

func (g *Game_State_Machine_t) State_Add(v Game_State_t) {
	g.Game_State = append(g.Game_State, v)
	v.Init(len(g.Game_State)-1, g.deltaTime)
}

func (g *Game_State_Machine_t) State_Change(v Game_State_t) {
	g.Game_State[len(g.Game_State)-1] = v
	v.Init(len(g.Game_State)-1, g.deltaTime)
}

func (g *Game_State_Machine_t) Debug_Log_print() {
	for i, e := range g.Game_State {
		fmt.Printf("%d %+v\n", i, e)
	}
	fmt.Printf("\n")
}

func (g *Game_State_Machine_t) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Layout_Width, g.Layout_Height
}

func (g *Game_State_Machine_t) Init() {
	g.old_time_for_deltaTime = time.Now().UnixNano()
}
