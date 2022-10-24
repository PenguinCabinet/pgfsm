/*
This is the state machine library for Ebiten.

これはGoのゲームライブラリEbiten用のステートマシンライブラリのutilityです。
*/
package pgfsm

import "github.com/hajimehoshi/ebiten"

/*
The structure of the state for the selection scene.

選択画面用のステートスタックマシーンの構造体です。
*/
type SelectState struct {
	OrignX int
	OrignY int

	IntervalX int
	IntervalY int

	InitFunc  func()
	DrawFuncs []func(*ebiten.Image, int, int, bool)

	DecisionFunc []func() Result
	CancelFunc   func() Result

	Index                  int
	Is_Loop                bool
	PressedNextKeyFunc     func() bool
	PressedBackKeyFunc     func() bool
	PressedDecisionKeyFunc func() bool
	PressedCancelKeyFunc   func() bool
}

func (s *SelectState) Next() {
	s.Index += 1
	if s.Index >= len(s.DrawFuncs) {
		if s.Is_Loop {
			s.Index = 0
		} else {
			s.Index = len(s.DrawFuncs) - 1
		}
	}
}
func (s *SelectState) Back() {
	s.Index -= 1
	if s.Index < 0 {
		if s.Is_Loop {
			s.Index = len(s.DrawFuncs) - 1
		} else {
			s.Index = 0
		}
	}
}

func (s *SelectState) Init(StackIndex int, delta float64) {
	s.InitFunc()
}

func (s *SelectState) Update(screen *ebiten.Image, StackIndex int, delta float64) Result {
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
func (s *SelectState) Draw(screen *ebiten.Image, StackIndex int, delta float64) {
	x := s.OrignX
	y := s.OrignY
	for i, e := range s.DrawFuncs {
		e(screen, x, y, i == s.Index)
		x += s.IntervalX
		y += s.IntervalY
	}
}
