package lld

import "fmt"

type VehicleType int

const (
	Motorcycle VehicleType = iota // iota=0, Motorcycle=0
	Car                           // iota=1, Car=1
	Truck                         // iota=2, Truck=2
)

type SpotSize int

const (
	Small  SpotSize = iota // 0
	Medium                 // 1
	Large                  // 2
)

type ParkingLot struct {
	Floors []*Floor // slice because many floors
}

type Spot struct {
	Size       SpotSize
	IsOccupied bool
	Vehicle    Vehicle // interface - no * needed
}

type Floor struct {
	FloorNumber int
	Spots       []*Spot // slice of struct pointers - * needed
}

type Ticket struct {
	Plate       string // which vehicle's plate?
	Spot        *Spot  // which spot?
	FloorNumber int    // which floor?
}

type Vehicle interface {
	GetType() VehicleType
	GetPlate() string
}

type MotorcycleVehicle struct {
	Plate string
}

type CarVehicle struct {
	Plate string
}

type TruckVehicle struct {
	Plate string
}

func (c *CarVehicle) GetPlate() string {
	return c.Plate

}

func (t *TruckVehicle) GetPlate() string {
	return t.Plate

}

func (m *MotorcycleVehicle) GetPlate() string {
	return m.Plate

}

func (m *MotorcycleVehicle) GetType() VehicleType {
	return Motorcycle
}

func (c *CarVehicle) GetType() VehicleType {
	return Car
}

func (t *TruckVehicle) GetType() VehicleType {
	return Truck
}

func (s *Spot) IsAvailable() bool {
	if !s.IsOccupied {
		return true
	}
	return false
}

func (s *Spot) Assign(v Vehicle) bool {
	// check if available
	if s.IsAvailable() {
		s.Vehicle = v
		s.IsOccupied = true
		return true
	}
	return false
}

func (s *Spot) Free() {

	s.Vehicle = nil
	s.IsOccupied = false

}

func (f *Floor) GetAvailableSpot(size SpotSize) *Spot {
	// loop through f.Spots
	// find a spot that matches size AND is available
	// return it if found
	// return nil if nothing found

	for _, spot := range f.Spots {
		if spot.Size == size && spot.IsAvailable() {
			return spot
		}
	}
	return nil

}

func getSpotSize(vType VehicleType) SpotSize {
	switch vType {
	case Motorcycle:
		return Small
	case Car:
		return Medium
	case Truck:
		return Large
	}
	return Medium
}

func (pl *ParkingLot) Park(v Vehicle) *Ticket {
	// 1. get spot size from vehicle type
	// 2. loop through all floors
	// 3. find available spot on each floor
	// 4. if found, assign vehicle to spot
	// 5. return ticket
	// 6. if nothing found, return nil

	spotSize := getSpotSize(v.GetType())
	for _, floor := range pl.Floors {
		spot := floor.GetAvailableSpot(spotSize)
		if spot != nil {
			spot.Assign(v)
			return &Ticket{
				Plate:       v.GetPlate(),
				Spot:        spot,
				FloorNumber: floor.FloorNumber,
			}
		}
	}
	return nil
}

func (pl *ParkingLot) Unpark(ticket *Ticket) {

	ticket.Spot.Free()
}

func (pl *ParkingLot) GetAvailableCount() int {
	// loop through all floors
	// loop through all spots on each floor
	// count spots that are available
	// return count
	count := 0
	for _, floor := range pl.Floors {
		for _, spot := range floor.Spots {
			if spot.IsAvailable() {
				count++
			}
		}
	}
	return count
}

func RunDemo() {
	// create parking lot with 2 floors
	lot := &ParkingLot{
		Floors: []*Floor{
			{
				FloorNumber: 1,
				Spots: []*Spot{
					{Size: Small},
					{Size: Medium},
					{Size: Large},
				},
			},
			{
				FloorNumber: 2,
				Spots: []*Spot{
					{Size: Small},
					{Size: Medium},
					{Size: Large},
				},
			},
		},
	}

	// park a car
	// check availability
	// unpark
	// check availability again

	car := &CarVehicle{Plate: "KA01AB1234"}
	ticket := lot.Park(car)
	fmt.Printf("Car parked on floor %d\n", ticket.FloorNumber)
	fmt.Printf("Available spots: %d\n", lot.GetAvailableCount())

	lot.Unpark(ticket)
	fmt.Printf("After unpark, available spots: %d\n", lot.GetAvailableCount())

}
