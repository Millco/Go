package main

import "fmt"

var x, y int

func kartenInReihe(x int) float64 {
	var p float64 = 1
	for i := 1; i <= x; i++ {
		p = p * (float64(1) / (52 - (float64(i) - 1)))
	}
	return p
}

func eineAusXKarten(x int) float64 {
	var p float64
	for i := 1; i <= x; i++ {
		p = p + (float64(i) / (52 - (float64(i) - 1)))
	}
	return p
}

func xKartenAusY(x, y int) float64 {
	var p float64 = 1
	for i := 1; i <= y; i++ {
		p = p * (float64(y) / (52 / float64(float64(y)-(float64(i)-1))))
	}
	return p
}

func main() {
	fmt.Println("Bitte geben Sie ein wieviele Karten gezogen werden sollen.")
	fmt.Scan(&x)
	fmt.Println("Wie viele Richtige werden benötigt?")
	fmt.Scan(&y)
	fmt.Printf("Die Wahrscheinlichkeit %v Karten in der richtigen Reihenfolge zu ziehen ist %e Prozent\n", x, kartenInReihe(x)*100)
	fmt.Printf("Die Wahrscheinlichkeit die richtige Karte in %v Versuchen zu ziehen ist %f Prozent\n", x, eineAusXKarten(x)*100)
	fmt.Printf("Die Wahrscheinlichkeit %v richtige Karten in %v Versuchen zu ziehen ist %e Prozent\n", y, x, xKartenAusY(x, y)*100)
}
