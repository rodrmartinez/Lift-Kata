package lift

import (
	"sort"
	"strconv"
)

// Direction ..
type Direction int

// Directions ..
const (
	Up Direction = iota
	Down
)

// Call ..
type Call struct {
	Floor     int
	Direction Direction
}

//Request
type Request struct {
	Lift  string
	Floor int
}

// Lift ..
type Lift struct {
	ID        string
	Floor     int
	Requests  []int
	DoorsOpen bool
	Monitor   string
}

// System ..
type System struct {
	floors []int
	lifts  []Lift
	calls  []Call
}

// NewSystem ..
func NewSystem() *System {
	return &System{floors: []int{}, lifts: []Lift{}, calls: []Call{}}
}

// AddFloors ..
func (s *System) AddFloors(floors ...int) {
	s.floors = append(s.floors, floors...)
}

// AddLifts ..
func (s *System) AddLifts(lifts ...Lift) {
	s.lifts = append(s.lifts, lifts...)
}

// AddCalls ..
func (s *System) AddCalls(calls ...Call) {
	s.calls = append(s.calls, calls...)
}

// AddRequest ..
func (s *System) AddRequest(request Request) {
	for i, lift := range s.lifts {
		if lift.ID == request.Lift {
			s.lifts[i].Requests = append(s.lifts[i].Requests, request.Floor)
		}
	}
}

// MoveToRequest ..
func (s *System) MoveToRequest() (output string) {
	output += PrintLiftStatus(s)
	for i, lift := range s.lifts {
		sort.Ints(s.lifts[i].Requests) //sort lift request to answer in order
		for _, request := range lift.Requests {
			for s.lifts[i].Floor != request {
				s.lifts[i].Tick(request)
				output += PrintLiftStatus(s)
			}
			s.lifts[i].Tick(request) //Open doors
			output += PrintLiftStatus(s)
			s.lifts[i].DoorsOpen = false //Close doors
			output += PrintLiftStatus(s)
		}
	}
	return
}

// MoveToCall ..
func (s *System) MoveToCall() (output string) {
	output += PrintLiftStatus(s)
	var responder Lift
	for _, call := range s.calls {
		responder = s.lifts[0]
		for i, lift := range s.lifts {
			if responder.Floor != lift.Floor {
				switch {
				case lift.Floor > call.Floor:
					responder = s.lifts[i]
				case lift.Floor < call.Floor:
					responder = s.lifts[i]
				}
			}
		}
		s.AddRequest(Request{responder.ID, call.Floor})
		output += s.MoveToRequest()
	}
	return
}

// CallsFor ..
func (s System) CallsFor(floor int) (calls []Call) {
	calls = []Call{}
	for _, c := range s.calls {
		if c.Floor == floor {
			calls = append(calls, c)
		}
	}
	return calls
}

//Tick
func (l *Lift) Tick(request int) {
	switch {
	case request == l.Floor:
		l.DoorsOpen = true
		l.Monitor = "*"
	case request > l.Floor:
		l.Floor += 1
		l.Monitor = strconv.Itoa(l.Floor)
	case request < l.Floor:
		l.Floor -= 1
		l.Monitor = strconv.Itoa(l.Floor)
	}
}
