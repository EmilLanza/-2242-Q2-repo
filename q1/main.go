package main

import "fmt"

type shape interface{
	area() float64
	perimeter() float64
}

type square struct{
	length float64
}

type circle struct{
	radius float64
}

func(s square) area() float64{
	return s.length * s.length
}

func(s square) perimeter() float64{
	return s.length * s.length * 2
}

func(c circle) area() float64{
	return 3.14 * c.radius * c.radius
}

func(c circle) perimeter() float64{
	return 2 * 3.14 * c.radius
}

func printShape(s shape){
	fmt.Println("area", s.area())
	fmt.Println("perimeter", s.area())
}

func main(){
	a := square{5}
	b := circle{5}
	printShape(a)
	printShape(b)
}