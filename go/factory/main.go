package main

import "fmt"

// interface of vehicle type
type iVehicle interface {
	SetBrand(brand string)
	SetModel(model string)
	SetCantOfWheels(cant int)
	GetBrand() string
	GetModel() string
	GetCantOfWheels() int
}

// Vehicle concrete with its methods
type Vehicle struct {
	Brand        string
	Model        string
	cantOfWheels int
}

func (v *Vehicle) SetBrand(brand string)    { v.Brand = brand }
func (v *Vehicle) SetModel(model string)    { v.Model = model }
func (v *Vehicle) SetCantOfWheels(cant int) { v.cantOfWheels = cant }

func (v *Vehicle) GetBrand() string     { return v.Brand }
func (v *Vehicle) GetModel() string     { return v.Model }
func (v *Vehicle) GetCantOfWheels() int { return v.cantOfWheels }

// An Specific vehicle creation (an object)
type teslaModelY struct {
	Vehicle
}

func newTeslaModelY() iVehicle {
	return &teslaModelY{Vehicle{
		Brand:        "Tesla",
		Model:        "Y",
		cantOfWheels: 4,
	},
	}
}

type teslaTruck struct {
	Vehicle
}

func newTeslaTruck() iVehicle {
	return &teslaTruck{Vehicle{
		Brand:        "Tesla",
		Model:        "Truck",
		cantOfWheels: 6,
	}}
}

// Factory method
func getVehicle(brand, model string) (iVehicle, error) {
	key := fmt.Sprintf("%s-%s", brand, model)
	switch key {
	case "tesla-y":
		return newTeslaModelY(), nil
	case "tesla-truck":
		return newTeslaTruck(), nil
	default:
		return nil, fmt.Errorf("brand and model doesnt exist")
	}
}

// main function to test this.
func main() {
	y, _ := getVehicle("tesla", "y")
	t, _ := getVehicle("tesla", "truck")

	printDetails(y)
	printDetails(t)
}

func printDetails(v iVehicle) {
	fmt.Printf("Brand: %s", v.GetBrand())
	fmt.Println()
	fmt.Printf("Model: %s", v.GetModel())
	fmt.Println()
	fmt.Printf("Quantity of wheels: %v", v.GetCantOfWheels())
	fmt.Println()
}
