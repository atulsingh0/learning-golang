package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {

	// one way of initialization
	var toronto location
	toronto.lat = 43.67028
	toronto.long = -79.38667

	// toronto.display("toronto")
	fmt.Println("Toronto:", toronto)

	// another way
	delhi := location{lat: 28.644800, long: 77.216721}
	delhi.display("delhi")
	// fmt.Println("Delhi:", delhi)

	// another way
	montreal := location{45.50889, -73.56167} // positional mapping
	fmt.Printf("Montreal: %v\n", montreal)
	fmt.Printf("Montreal: %v %v\n", montreal.lat, montreal.long)
	fmt.Printf("Montreal: %+v\n", montreal)

	// changing a struct
	addSomething(delhi) // No change, Copy by value
	fmt.Printf("Delhi: %+v\n", delhi)

}

func (lc *location) display(city string) {
	fmt.Printf("%v Latitude: %v\n", city, lc.lat)
	fmt.Printf("%v Longitude: %v\n", city, lc.long)
}

func addSomething(lc location) {
	lc.lat += 10
}
