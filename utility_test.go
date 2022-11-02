/*
This is the state machine library for Ebiten.

これはGoのゲームライブラリEbiten用のステートマシンライブラリのutilityです。
*/
package pgfsm

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

func TestSelectState(t *testing.T) {
	SelectedIndex := -1
	s := SelectState{
		OrignX:    100,
		OrignY:    100,
		IntervalX: 0,
		IntervalY: 30,
		DrawFuncs: []func(*ebiten.Image, int, int, bool){
			func(screen *ebiten.Image, x, y int, b bool) {
				assert.Equal(t, x, 100)
				assert.Equal(t, y, 100)
			},
			func(screen *ebiten.Image, x, y int, b bool) {
				assert.Equal(t, x, 130)
				assert.Equal(t, y, 100)
			},
			func(screen *ebiten.Image, x, y int, b bool) {
				assert.Equal(t, x, 160)
				assert.Equal(t, y, 100)
			},
		},
		DecisionFunc: []func() Result{
			func() Result {
				SelectedIndex = 0
				return Result{
					Code:      CodeNil,
					NextState: nil,
				}
			},
			func() Result {
				SelectedIndex = 1
				return Result{
					Code:      CodeNil,
					NextState: nil,
				}
			},
			func() Result {
				SelectedIndex = 2
				return Result{
					Code:      CodeNil,
					NextState: nil,
				}
			},
		},
		IsLoop: true,
	}

	assert.Equal(t, s.Index, 0)
	s.Next()
	assert.Equal(t, s.Index, 1)
	s.Next()
	assert.Equal(t, s.Index, 2)
	s.Next()
	assert.Equal(t, s.Index, 0)

	s.DecisionFunc[s.Index]()
	assert.Equal(t, SelectedIndex, 0)

	s.Back()
	assert.Equal(t, s.Index, 2)
	s.Back()
	assert.Equal(t, s.Index, 1)
	s.Back()
	assert.Equal(t, s.Index, 0)

}
