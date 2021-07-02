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
	panic("Implement this method")
}
