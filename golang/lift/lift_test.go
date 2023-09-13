package lift_test

import (
	"testing"

	approvaltests "github.com/approvals/go-approval-tests"
	"github.com/lift-kata/lift"
)

func TestAddRequest(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 0, []int{}, false, "0"}
	request := lift.Request{Lift: "A", Floor: 3}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA)
	liftSystem.AddRequest(request)

	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))

}

func TestMoveUpToRequest(t *testing.T) {

	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 0, []int{3}, false, "0"}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA)
	output := liftSystem.MoveToRequest()

	approvaltests.VerifyString(t, output)
}

func TestMoveDownToRequest(t *testing.T) {

	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 3, []int{1}, false, "3"}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA)
	output := liftSystem.MoveToRequest()

	approvaltests.VerifyString(t, output)
}

func TestMoveDownToCall(t *testing.T) {

	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 3, []int{}, false, "3"}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA)
	liftSystem.AddCalls(lift.Call{1, lift.Down})
	output := liftSystem.MoveToCall()

	approvaltests.VerifyString(t, output)
}

func TestMoveUpToCall(t *testing.T) {

	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 0, []int{}, false, "0"}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA)
	liftSystem.AddCalls(lift.Call{3, lift.Down})
	output := liftSystem.MoveToCall()

	approvaltests.VerifyString(t, output)
}

func TestMoveManyLiftsToRequests(t *testing.T) {

	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 0, []int{}, false, "0"}
	liftB := lift.Lift{"B", 0, []int{}, false, "0"}
	requestA := lift.Request{Lift: "A", Floor: 3}
	requestB := lift.Request{Lift: "B", Floor: 2}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA, liftB)
	liftSystem.AddRequest(requestA)
	liftSystem.AddRequest(requestB)
	output := liftSystem.MoveToRequest()

	approvaltests.VerifyString(t, output)
}

func TestMoveManyToCall(t *testing.T) {

	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 0, []int{}, false, "0"}
	liftB := lift.Lift{"B", 0, []int{}, false, "0"}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA, liftB)
	liftSystem.AddCalls(lift.Call{3, lift.Up})
	output := liftSystem.MoveToCall()

	approvaltests.VerifyString(t, output)
}

func TestAnserManyCalls(t *testing.T) {

	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 0, []int{}, false, "0"}
	liftB := lift.Lift{"B", 0, []int{}, false, "0"}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA, liftB)
	liftSystem.AddCalls(lift.Call{3, lift.Up})
	liftSystem.AddCalls(lift.Call{2, lift.Up})
	output := liftSystem.MoveToCall()

	approvaltests.VerifyString(t, output)
}

func TestMoveOneLiftToManyRequests(t *testing.T) {

	liftSystem := lift.NewSystem()
	liftA := lift.Lift{"A", 0, []int{}, false, "0"}
	liftB := lift.Lift{"B", 0, []int{}, false, "0"}
	requestA := lift.Request{Lift: "A", Floor: 3}
	requestB := lift.Request{Lift: "A", Floor: 2}

	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddLifts(liftA, liftB)
	liftSystem.AddRequest(requestA)
	liftSystem.AddRequest(requestB)
	output := liftSystem.MoveToRequest()

	approvaltests.VerifyString(t, output)
}
