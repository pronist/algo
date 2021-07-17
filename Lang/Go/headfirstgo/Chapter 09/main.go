package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}
func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}
func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

//func LitersToGallons(l Liters) Gallons {
//	return Gallons(l * 0.264)
//}
//func MillilitersToGallons(m Milliliters) Gallons {
//	return Gallons(m * 0.000264)
//}
//func GallonsToLiters(g Gallons) Liters {
//	return Liters(g * 3.785)
//}
//func GallonsToMilliliters(g Gallons) Milliliters {
//	return Milliliters(g * 3785.41)
//}

func main() {
	soda := Liters(2)
	fmt.Println(soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Println(water, water.ToGallons())
}
