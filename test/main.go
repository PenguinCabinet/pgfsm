package main

import (
	"image/color"
	"log"

	"github.com/PenguinCabinet/Pen_Game_State_Machine"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//これがメニュー画面のステート
type Menu_Game_State_t struct {
	mplusNormalFont font.Face
}

//これがステートが最初に実行されたときに呼び出される関数
func (sm *Menu_Game_State_t) Init(
	stack_deep int, /*ここにはこのステートがスタックのどの位置に積まれているかインデックスが入っています*/
	delta float64, /*ここには前のフレームと今のフレーム間で経過した時間が入っています*/
) {
	/*ここから Ebitenのフォントの初期化処理*/
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
	/*ここまで Ebitenのフォントの初期化処理*/
}

//これはマイフレーム呼び出される関数です
//このステートが実行されている時のみ、呼び出されます
func (sm *Menu_Game_State_t) Update(
	screen *ebiten.Image, /*ebitenのscreenですが、Updateで描写するのは非推奨です*/
	stack_deep int, delta float64,
) Pen_Game_State_Machine.Game_State_result_t {

	/*mキーが入力された場合 メニューを閉じる*/
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		/*ここで現在実行しているメニュー画面のステートマシンを消去します
		「ゲーム画面、メニュー画面」の順でスタックにストックされているので、消去するとスタックの中身は
		「ゲーム画面」となってゲーム画面に戻ります
		*/
		return Pen_Game_State_Machine.Game_State_result_t{
			Code:       Pen_Game_State_Machine.Game_State_result_delete,
			Next_State: nil,
		}
	}
	/*空のPen_Game_State_Machine.Game_State_result_tを返却することでループを継続します
	Pen_Game_State_Machine.Game_State_result_tを書き換えることで、実行するものを新しいステートに変えたり
	新しいステートをスタックの上に乗せたりすることができます*/
	return Pen_Game_State_Machine.Game_State_result_t{}
}

//これはマイフレーム呼び出される描写用の関数です
//このステートが実行されていなくても、スタック上にあれば呼び出されます
func (sm *Menu_Game_State_t) Draw(screen *ebiten.Image, stack_deep int, delta float64) {
	text.Draw(screen, "Menu", sm.mplusNormalFont, 300, 240, color.White)
}

//これがゲーム画面のステート
type Game_Main_State_t struct {
	mplusNormalFont font.Face
}

//これがステートが最初に実行されたときに呼び出される関数
func (sm *Game_Main_State_t) Init(
	stack_deep int, /*ここにはこのステートがスタックのどの位置に積まれているかインデックスが入っています*/
	delta float64, /*ここには前のフレームと今のフレーム間で経過した時間が入っています*/
) {
	/*ここから Ebitenのフォントの初期化処理*/
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
	/*ここまで Ebitenのフォントの初期化処理*/
}

//これはマイフレーム呼び出される関数です
//このステートが実行されている時のみ、呼び出されます
//つまりメニューを開いている間は、ゲーム画面のUpdate関数が実行されません
func (sm *Game_Main_State_t) Update(
	screen *ebiten.Image, /*ebitenのscreenですが、Updateで描写するのは非推奨です*/
	stack_deep int, delta float64,
) Pen_Game_State_Machine.Game_State_result_t {
	/*mキーが入力された場合 メニューを開く*/
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		/*ここで現在実行しているゲーム画面の上にメニュー画面のステートをのせます
		「ゲーム画面」の順でスタックにストックされているので、追加するとスタックの中身は
		「ゲーム画面、メニュー画面」となってメニュー画面の処理に移ります
		*/
		return Pen_Game_State_Machine.Game_State_result_t{
			Code:       Pen_Game_State_Machine.Game_State_result_add,
			Next_State: new(Menu_Game_State_t),
		}
	}

	/*空のPen_Game_State_Machine.Game_State_result_tを返却することでループを継続します
	Pen_Game_State_Machine.Game_State_result_tを書き換えることで、実行するものを新しいステートに変えたり
	新しいステートをスタックの上に乗せたりすることができます*/
	return Pen_Game_State_Machine.Game_State_result_t{}
}

//これはマイフレーム呼び出される描写用の関数です
//このステートが実行されていなくても、スタック上にあれば呼び出されます
//つまりメニューを開いている間も、ゲーム画面のdraw関数が実行されます
func (sm *Game_Main_State_t) Draw(screen *ebiten.Image, stack_deep int, delta float64) {
	text.Draw(screen, "Game Main", sm.mplusNormalFont, 200, 100, color.White)
}

//これがタイトル画面のステート
type Title_Game_State_t struct {
	mplusNormalFont font.Face
}

//これがステートが最初に実行されたときに呼び出される関数
func (sm *Title_Game_State_t) Init(
	stack_deep int, /*ここにはこのステートがスタックのどの位置に積まれているかインデックスが入っています*/
	delta float64, /*ここには前のフレームと今のフレーム間で経過した時間が入っています*/
) {
	/*ここから Ebitenのフォントの初期化処理*/
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
	/*ここまで Ebitenのフォントの初期化処理*/
}

//これはマイフレーム呼び出される関数です
//このステートが実行されている時のみ、呼び出されます
func (sm *Title_Game_State_t) Update(
	screen *ebiten.Image, /*ebitenのscreenですが、Updateで描写するのは非推奨です*/
	stack_deep int, delta float64,
) Pen_Game_State_Machine.Game_State_result_t {

	/*sキーが入力された場合*/
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		/*ここでステートマシンを切り替えます
		Pen_Game_State_Machine.Game_State_result_changeは現在実行しているステートを
		Next_Stateに切り替わります
		ここでは現在実行中のタイトル画面のステートからゲーム画面のステートに切り替えています*/
		return Pen_Game_State_Machine.Game_State_result_t{
			Code:       Pen_Game_State_Machine.Game_State_result_change,
			Next_State: new(Game_Main_State_t),
		}
	}
	/*空のPen_Game_State_Machine.Game_State_result_tを返却することでループを継続します
	Pen_Game_State_Machine.Game_State_result_tを書き換えることで、実行するものを新しいステートに変えたり
	新しいステートをスタックの上に乗せたりすることができます*/
	return Pen_Game_State_Machine.Game_State_result_t{}
}

//これはマイフレーム呼び出される描写用の関数です
//このステートが実行されていなくても、スタック上にあれば呼び出されます
func (sm *Title_Game_State_t) Draw(screen *ebiten.Image, stack_deep int, delta float64) {
	text.Draw(screen, "Game Title\nPressing S key,start!", sm.mplusNormalFont, 100, 100, color.White)
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pen_Game_State_Machine")

	gms := new(Pen_Game_State_Machine.Game_State_Machine_t)

	gms.Layout_Width = 640
	gms.Layout_Height = 480

	Title_sm := new(Title_Game_State_t)

	/*スタックにmy_smを追加します*/
	gms.State_Add(Title_sm)

	if err := ebiten.RunGame(gms); err != nil {
		log.Fatal(err)
	}
}
