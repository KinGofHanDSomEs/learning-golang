package main

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
	GetType() string
}

type Car struct {
	Type     string
	Speed    float64
	FuelType string
}

type Motorcycle struct {
	Type     string
	Speed    float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	if c.Speed <= 0 {
		return 0
	}
	return distance / c.Speed
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	if m.Speed <= 0 {
		return 0
	}
	return distance / m.Speed
}

func (c Car) GetType() string {
	return "main.Car"
}

func (m Motorcycle) GetType() string {
	return "main.Motorcycle"
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	res := make(map[string]float64)
	for _, vehicle := range vehicles {
		res[vehicle.GetType()] = vehicle.CalculateTravelTime(distance)
	}
	return res
}
