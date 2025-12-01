package main

import "fmt"

func computeGrade(points float32, maxPoints float32) float32{
	grade := (points/maxPoints)*5+1
	return grade
}

func main() {
	grade := computeGrade(50, 50)
	fmt.Printf("Grade: " + "%.2f\n", grade)
}