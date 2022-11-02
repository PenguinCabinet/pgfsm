/*
This is the state machine library for Ebiten.

これはGoのゲームライブラリEbiten用のステートマシンライブラリのutilityです。
*/
package pgfsm

import "github.com/hajimehoshi/ebiten/v2"

/*
The structure of the state for the selection scene.

選択画面用のステートスタックマシーンの構造体です。
*/
type SelectState struct {
	/*
		Coordinates where the first selection is placed.

		一番最初の選択肢がおかれる座標。
	*/
	OrignX int
	OrignY int

	/*
		Spacing between alternatives.
		選択肢と選択肢の間隔。
	*/
	IntervalX int
	IntervalY int

	/*
		Function to process initialization.
		初期化の処理をする関数。
	*/
	InitFunc func()

	/*
		Functions that describes a selection.

		選択肢を描写する関数。
	*/
	DrawFuncs []func(*ebiten.Image, int, int, bool)

	/*
		Function to be executed when a decision is made.

		決定したときに実行される関数。
	*/
	DecisionFunc []func() Result

	/*
		Function to be executed when canceled.

		キャンセルしたときに実行される関数。
	*/
	CancelFunc func() Result

	/*
		Index of the currently selected option.

		現在選択中の選択肢のインデックス。
	*/
	Index int
	/*
		Whether to loop if said before the end of the current selection.

		現在選択肢の端より先に言った場合、ループするかどうか。
	*/
	IsLoop bool

	/*
		Discriminant function to go to the next selection.

		次の選択肢へ行くための判別関数。
	*/
	PressedNextKeyFunc func() bool

	/*
		Discriminant function to go to the previous selection.

		前の選択肢へ行くための判別関数。
	*/
	PressedBackKeyFunc func() bool

	/*
		Discriminant function for decision.

		決定のための判別関数。
	*/
	PressedDecisionKeyFunc func() bool

	/*
		Discriminant function for cancellation.

		キャンセルのための判別関数。
	*/
	PressedCancelKeyFunc func() bool
}

/*
Function to advance the the current selection.

選択した選択肢を進める関数。
*/
func (s *SelectState) Next() {
	s.Index += 1
	if s.Index >= len(s.DrawFuncs) {
		if s.IsLoop {
			s.Index = 0
		} else {
			s.Index = len(s.DrawFuncs) - 1
		}
	}
}

/*
Function to return the the current selection.

選択した選択肢を戻る関数。
*/
func (s *SelectState) Back() {
	s.Index -= 1
	if s.Index < 0 {
		if s.IsLoop {
			s.Index = len(s.DrawFuncs) - 1
		} else {
			s.Index = 0
		}
	}
}

func (s *SelectState) Init(StackIndex int) {
	if s.InitFunc != nil {
		s.InitFunc()
	}
}

func (s *SelectState) Update(StackIndex int) Result {
	if s.PressedNextKeyFunc != nil {
		if s.PressedNextKeyFunc() {
			s.Next()
		}
	}
	if s.PressedBackKeyFunc != nil {
		if s.PressedBackKeyFunc() {
			s.Back()
		}
	}
	if s.PressedDecisionKeyFunc != nil {
		if s.PressedDecisionKeyFunc() {
			if s.DecisionFunc[s.Index] != nil {
				return s.DecisionFunc[s.Index]()
			}
		}
	}

	if s.PressedCancelKeyFunc != nil {
		if s.PressedCancelKeyFunc() {
			if s.CancelFunc != nil {
				return s.CancelFunc()
			}
		}
	}
	return Result{
		Code:      CodeNil,
		NextState: nil,
	}
}
func (s *SelectState) Draw(screen *ebiten.Image, StackIndex int) {
	x := s.OrignX
	y := s.OrignY
	for i, e := range s.DrawFuncs {
		e(screen, x, y, i == s.Index)
		x += s.IntervalX
		y += s.IntervalY
	}
}
