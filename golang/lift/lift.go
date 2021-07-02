package lift

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

// MovetoRequest ..
func (s *System) MoveToRequest() {
	for i, lift := range s.lifts {
		for _, request := range lift.Requests {
			for s.lifts[i].Floor != request {
				s.lifts[i].Tick()
			}
		}
		s.lifts[i].Tick() //Open doors
	}
}

// MovetoCall ..
func (s *System) MoveToCall() {
	for _, call := range s.calls {
		for _, lift := range s.lifts {
			if lift.Floor != call.Floor {
				s.AddRequest(Request{lift.ID, call.Floor})
				s.MoveToRequest()
			}
		}
	}
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

// Tick ..
func (s System) Tick() {
	for i, _ := range s.lifts {
		s.lifts[i].Tick()
	}
}

func (l *Lift) Tick() {

	for _, request := range l.Requests {
		switch {
		case request == l.Floor:
			l.Requests = []int{}
			l.DoorsOpen = true
		case request > l.Floor:
			l.Floor += 1
		case request < l.Floor:
			l.Floor -= 1
		}
	}
}
