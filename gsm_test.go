package pgfsm

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

type StateForTest struct {
	R  Result
	ID string
}

func (sm *StateForTest) Init(
	stackdeep int,
) {
}

func (sm *StateForTest) Update(
	stackdeep int,
) Result {
	return sm.R
}

func (sm *StateForTest) Draw(screen *ebiten.Image, stackdeep int) {
}

func TestEmptyOfMachine(t *testing.T) {

	gms := &Machine{}

	assert.Equal(t, gms.Empty(), true, "Run Empty func when empty.")
	gms.StateAdd(&StateForTest{R: Result{}})
	assert.Equal(t, gms.Empty(), false, "Run Empty func when not empty.")
}

func TestCodeNilOfMachine(t *testing.T) {

	gms := &Machine{}

	gms.StateAdd(&StateForTest{
		R:  Result{},
		ID: "State1",
	})
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State1", "Check the id of the current state")
	//continue one frame.
	_ = gms.Update()
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State1", "Check the id of the current state")
}

func TestCodeChangeOfMachine(t *testing.T) {

	gms := &Machine{}

	gms.StateAdd(&StateForTest{
		ID: "State1",
		R: Result{
			Code: CodeChange,
			NextState: &StateForTest{
				ID: "State2",
				R:  Result{},
			},
		},
	})
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State1", "Check the id of the current state")
	//continue one frame.
	_ = gms.Update()
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State2", "Check the id of the current state")
}

func TestCodeAddOfMachine(t *testing.T) {

	gms := &Machine{}

	gms.StateAdd(&StateForTest{
		ID: "State1",
		R: Result{
			Code: CodeAdd,
			NextState: &StateForTest{
				ID: "State2",
				R:  Result{},
			},
		},
	})

	assert.Equal(t, len(gms.StateStack), 1, "Check the length of the states")
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State1", "Check the id of the state[0]")
	//continue one frame.
	_ = gms.Update()
	assert.Equal(t, len(gms.StateStack), 2, "Check the length of the states")
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State1", "Check the id of the state[0]")
	assert.Equal(t, gms.StateStack[1].(*StateForTest).ID, "State2", "Check the id of the state[1]")
}

func TestCodeDeleteOfMachine(t *testing.T) {

	gms := &Machine{}

	gms.StateAdd(&StateForTest{
		ID: "State1",
		R: Result{
			Code:      CodeDelete,
			NextState: nil,
		},
	})
	assert.Equal(t, gms.Empty(), false, "Run Empty func when not empty.")
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State1", "Check the id of the current state")
	//continue one frame.
	_ = gms.Update()
	assert.Equal(t, gms.Empty(), true, "Run Empty func when empty.")
}

func TestCodeAllDeleteAndChangeOfMachine(t *testing.T) {

	gms := &Machine{}

	gms.StateAdd(&StateForTest{
		R:  Result{},
		ID: "State1",
	})
	gms.StateAdd(&StateForTest{
		R:  Result{},
		ID: "State2",
	})
	gms.StateAdd(&StateForTest{
		R:  Result{},
		ID: "State3",
	})

	gms.StateAdd(&StateForTest{
		ID: "State4",
		R: Result{
			Code: CodeAllDeleteAndChange,
			NextState: &StateForTest{
				ID: "State5",
				R:  Result{},
			},
		},
	})

	assert.Equal(t, len(gms.StateStack), 4, "Check the length of the states")
	//continue one frame.
	_ = gms.Update()
	assert.Equal(t, len(gms.StateStack), 1, "Check the length of the states")
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State5", "Check the id of the current state")
}

func TestCodeInsertBackOfMachine(t *testing.T) {

	gms := &Machine{}

	gms.StateAdd(&StateForTest{
		ID: "State1",
		R: Result{
			Code: CodeInsertBack,
			NextState: &StateForTest{
				ID: "State2",
				R:  Result{},
			},
		},
	})

	assert.Equal(t, len(gms.StateStack), 1, "Check the length of the states")
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State1", "Check the id of the state[0]")
	//continue one frame.
	_ = gms.Update()
	assert.Equal(t, len(gms.StateStack), 2, "Check the length of the states")
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State2", "Check the id of the state[0]")
	assert.Equal(t, gms.StateStack[1].(*StateForTest).ID, "State1", "Check the id of the state[1]")

	gms = &Machine{}

	gms.StateAdd(&StateForTest{
		R:  Result{},
		ID: "State1",
	})
	gms.StateAdd(&StateForTest{
		ID: "State2",
		R: Result{
			Code: CodeInsertBack,
			NextState: &StateForTest{
				ID: "State3",
				R:  Result{},
			},
		},
	})

	assert.Equal(t, len(gms.StateStack), 2, "Check the length of the states")
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State1", "Check the id of the state[0]")
	assert.Equal(t, gms.StateStack[1].(*StateForTest).ID, "State2", "Check the id of the state[1]")
	//continue one frame.
	_ = gms.Update()
	assert.Equal(t, len(gms.StateStack), 3, "Check the length of the states")
	assert.Equal(t, gms.StateStack[0].(*StateForTest).ID, "State1", "Check the id of the state[0]")
	assert.Equal(t, gms.StateStack[1].(*StateForTest).ID, "State3", "Check the id of the state[1]")
	assert.Equal(t, gms.StateStack[2].(*StateForTest).ID, "State2", "Check the id of the state[2]")
}
