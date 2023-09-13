package lift_test

import (
	"testing"

	approvaltests "github.com/approvals/go-approval-tests"
	"github.com/lift-kata/lift"
)

func TestPrintNoLifts(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddFloors(0, 1, 2, 3)
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}

func TestNoDoors(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(lift.Lift{"A", 3, []int{0}, false, "3"})
	liftSystem.AddCalls(lift.Call{1, lift.Down})
	liftSystem.AddFloors(0, 1, 2, 3)
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewSimplePrinter()))
}

func TestSimpleLiftSystem(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(
		lift.Lift{"A", 3, []int{0}, false, "3"},
		lift.Lift{"B", 2, []int{}, false, "2"},
		lift.Lift{"C", 2, []int{}, true, "2"},
		lift.Lift{"D", 0, []int{0}, false, "0"})
	liftSystem.AddCalls(lift.Call{1, lift.Down})
	liftSystem.AddFloors(0, 1, 2, 3)
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}

func TestIllegalState(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(lift.Lift{"A", 0, []int{0}, true, "0"})
	liftSystem.AddFloors(0, 1)
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}

func TestLargeLiftSystem(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(
		lift.Lift{"A", 0, []int{3, 5, 7}, false, "0"},
		lift.Lift{"B", 2, []int{}, true, "2"},
		lift.Lift{"C", -2, []int{-2, 0}, false, "-2"},
		lift.Lift{"D", 8, []int{0, -1, -2}, true, "8"},
		lift.Lift{"SVC", 10, []int{0, -1}, false, "10"},
		lift.Lift{"F", 8, []int{}, false, "8"})
	liftSystem.AddFloors(-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	liftSystem.AddCalls(
		lift.Call{1, lift.Down},
		lift.Call{6, lift.Down},
		lift.Call{5, lift.Up},
		lift.Call{5, lift.Down},
		lift.Call{-1, lift.Up})
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}
