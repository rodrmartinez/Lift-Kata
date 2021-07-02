package lift_test

import (
	"testing"

	approvaltests "github.com/approvals/go-approval-tests"
	"github.com/lift-kata/lift"
)

func TestAddRequest(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 0, []int{}, false}
	request := lift.Request{Lift: "A", Floor: 3}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA)
	liftSystem.AddRequest(request)

	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))

}

func TestMoveOneToRequest(t *testing.T) {

	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 0, []int{3}, false}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA)
	liftSystem.Tick()

	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}
