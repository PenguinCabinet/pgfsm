[README of ENGLISH](./README_en.md)
# The Pen Game Programing Finite State Machine

これは[GoのゲームライブラリEbiten用](https://ebiten.org/)のステートマシンライブラリです。
## 提供する機能
* ゲームプログラミングのためのスタック型のステートマシン

# ドキュメント
* [チュートリアル(JP)](doc/Tutorial.md)   
* [Utilityチュートリアル(JP)](doc/Utility_Tutorial.md)   
* [pkg.go.dev](https://pkg.go.dev/github.com/PenguinCabinet/pgfsm)
* [Examples](examples/)

# 採用実績
絶賛募集中です。   
このライブラリを採用しているゲームで載せてもいいよという方は是非Issueで教えてください。

# Quick start

## インストール
```shell
go get github.com/PenguinCabinet/pgfsm
```

## Example

```go
package main

import (
	"log"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten/v2"
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

