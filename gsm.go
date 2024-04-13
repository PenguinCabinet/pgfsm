/*
This is the state machine library for Ebiten.

これはGoのゲームライブラリEbiten用のステートマシンライブラリです。
*/
package pgfsm

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
The type of values for the type of operations the stack.

If its value is changed, In example,

Do nothing and continue.
Change the current running state to NextState.
Add NextState on the current running state in the stack...etc.

You can operate the stack.


これは現在実行中のステートから返される、スタックを操作する種類を決めるための変数型です

この値を変えると、例えば、

何もせずそのまま処理実行。
現在実行しているStateとNextStateを入れ替える。
現在スタックに乗っているStateをすべて消去して、NextStateを乗せる。

とスタックを操作することができます。
*/
type Code int

const (
	/*Do nothing and continue
	何もせずそのまま処理実行*/
	CodeNil Code = iota

	/*Change the current running state to NextState
	現在実行しているStateとNextStateを入れ替える*/
	CodeChange

	/*Add NextState on the current running state in the stack
	現在実行しているStateの上にNextStateを乗せる*/
	CodeAdd

	/*Delete the current running state
	現在実行しているStateを消去する*/
	CodeDelete

	/*Delete all states in the stack and add NextState in the stack.
	現在スタックに乗っているStateをすべて消去して、NextStateを乗せる*/
	CodeAllDeleteAndChange

	/*Insert NextState behiend the current running state
	現在実行しているStateの後ろに、NextStateを挿入する*/
	CodeInsertBack
)

/*
The type of values for operations the stack.
If it is needed,set NextState.

これは現在実行中のステートから返される、スタックを操作するための変数型です
必要ならば、NextStateに値を入れてください
*/
type Result struct {
	Code      Code
	NextState State
}

/*
The interface of state.

ステートのインターフェース
*/
type State interface {
	Init(int)
	Update(int) Result
	Draw(*ebiten.Image, int)
}

/*
The struct of the stack and state machine.

ステートスタックマシーンの構造体です
*/
type Machine struct {
	StateStack []State

	/*Please set the window size here.
	It is used in the layout function for Ebiten
	Windowのサイズを入れてください。
	Ebiten用のレイアウト関数で使用されます*/
	LayoutWidth, LayoutHeight int
}

/*This is the function for processing that is called in every frames.It is used internally.

これは処理のため毎フレーク呼び出される関数です。内部的にEbitenで使います*/
func (g *Machine) Update() error {

	if g.Empty() {
		return nil
	}

	res := g.StateStack[len(g.StateStack)-1].Update(len(g.StateStack) - 1)

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
		res.NextState.Init(len(g.StateStack) - 1)
		index := len(g.StateStack) - 1
		g.StateStack = append(g.StateStack[:index+1], g.StateStack[index:]...)
		g.StateStack[index] = res.NextState

	}

	return nil
}

/*This is the function for drawing that is called in every frames.It is used internally.

これは描写のため毎フレーク呼び出される関数です。内部的にEbitenで使います*/
func (g *Machine) Draw(screen *ebiten.Image) {
	for d, e := range g.StateStack {
		e.Draw(screen, d)
	}
}

/*This returns true if the stack is empty, false otherwise

stackが空ならtrueを返し、それ以外ならfalseを返します*/
func (g *Machine) Empty() bool {
	return len(g.StateStack) == 0
}

/*This add argument v to the stack

stackに引数vを追加します*/
func (g *Machine) StateAdd(v State) {
	g.StateStack = append(g.StateStack, v)
	v.Init(len(g.StateStack) - 1)
}

/*This changes the current running state in the stack to the argument v

実行中のステートを引数vに変更します*/
func (g *Machine) StateChange(v State) {
	g.StateStack[len(g.StateStack)-1] = v
	v.Init(len(g.StateStack) - 1)
}

/*The function for debugging.It displays the all states of the stack on the console.

デバッグ用の関数です。コンソールにスタック内のステートをすべて表示します*/
func (g *Machine) DebugLogprint() {
	for i, e := range g.StateStack {
		fmt.Printf("%d %+v\n", i, e)
	}
	fmt.Printf("\n")
}

/*This is the function for Ebiten.It is used internally.
内部的にEbitenで使います*/
func (g *Machine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.LayoutWidth, g.LayoutHeight
}

/*This is the function for Ebiten.It is used internally.
内部的にEbitenで使います*/
func (g *Machine) Init() {
}
