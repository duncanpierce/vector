package scalable

import (
	"fmt"
	"testing"
)

type State struct {
	x, y, z int
}

func ForVars[T any](vars *T, f func(vars *T)) {
	f(vars)
}

func Test(t *testing.T) {
	vars := &State{}
	vars.x = 10
	vars.y = 20
	ForVars[State](vars, func(v *State) {
		v.z = v.x + v.y
	})
	fmt.Printf("got %v", vars.z)
}

func TestOlderLoop(t *testing.T) {
	x := NewBunch[int]()
	y := NewBunch[int]()
	z := NewBunch[int]()
	For()
}
