package lld

import "fmt"

type Elevator struct {
	ID           int
	CurrentFloor int
	State        ElevatorState
	Capacity     int
	CurrentLoad  int
	DoorOpen     bool
}

type Request struct {
	Floor     int
	Direction Direction
}

type Building struct {
	Floors    int
	Elevators []*Elevator
	Requests  []*Request
}

type ElevatorState int

const (
	Idle ElevatorState = iota
	MovingUp
	MovingDown
	DoorsOpen
)

type Direction int

const (
	Up Direction = iota
	Down
)

func (e *Elevator) IsAvailable() bool {

	if e.State == Idle && e.CurrentLoad < e.Capacity {
		return true
	}
	return false

}

func (b *Building) RequestElevator(r *Request) *Elevator {
	// find nearest available elevator
	// assign it
	// return it

	minValue := 999
	var nearest *Elevator
	for _, e := range b.Elevators {
		if e.IsAvailable() {
			distance := e.CurrentFloor - r.Floor
			if distance < 0 {
				distance = -distance // absolute value
			}
			if distance < minValue {
				minValue = distance
				nearest = e
			}

		}
	}
	return nearest

}

func DemoElevator() {

	b := &Building{

		Floors: 10,
		Elevators: []*Elevator{
			&Elevator{
				ID:           01,
				CurrentFloor: 10,
				State:        0,
				Capacity:     5,
				CurrentLoad:  2,
				DoorOpen:     false,
			},
			&Elevator{
				ID:           02,
				CurrentFloor: 10,
				State:        1,
				Capacity:     5,
				CurrentLoad:  2,
				DoorOpen:     true,
			},
		},
		Requests: []*Request{},
	}

	r := &Request{Floor: 3, Direction: Up}
	b.Requests = append(b.Requests, r)
	elevator := b.RequestElevator(r)
	fmt.Println(elevator.ID)

}
