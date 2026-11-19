package vreme

import (
	"fmt"
)

// vhodne vrednosti spremenljivk
var (
	minTemp   = -50.0
	maxTemp   = 50.0
	stMeritev = 3
)

// Postaja predstavlja eno vremensko postajo, za razlicne senzorje.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

var postaje = map[string]Postaja{}

// NastaviVhodne nastavi vrednosti vhodnih spremenljivk.
func NastaviVhodne(min, max float64, st int) {
	minTemp, maxTemp, stMeritev = min, max, st
}

func DodajPostajo(oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{ime, kraj, []float64{}}
}

// DodajMeritev omogoca dodajanje ustrezne meritve za vremensko postajo.
func DodajMeritev(oznaka string, temperatura float64) {
	if _, ok := postaje[oznaka]; !ok {
		fmt.Println("Izbrana oznaka ne obstaja, \nMeritev ni bila dodana!")
		return
	}

	if !(temperatura > minTemp && temperatura < maxTemp) {
		fmt.Printf("\nTemperatura je izven meje (%f in %f ~ neinkluzivno), \nMeritev ni bila dodana!", minTemp, maxTemp)
		return
	}

	p := postaje[oznaka]
	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
}

// IzpisMeritev na stdout izpise vse trenutne meritve, ki so jih vremenske postaje izvedele.
func IzpisMeritev() {
	fmt.Println("MERITVE:")
	for key, val := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", key, val.ime, val.kraj, val.meritve)
	}
}

// IzpisPorocila na stdout izpise porocilo glede na meritve posamezne vremenske postaje.
func IzpisPorocila() {

	for key, val := range postaje {
		min, max, avg, ok := statistika(key)
		temp := ""

		if ok {
			switch {
			case avg >= 25:
				temp = "Vroče"
			case avg >= 15 && avg < 25:
				temp = "Toplo"
			case avg >= 5 && avg < 15:
				temp = "Hladno"
			case avg < 5:
				temp = "Mrzlo"
			}

			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C, -> %v\n", val.ime, min, max, avg, temp)
		} else {
			fmt.Printf("%s: premalo meritev\n", val.ime)
		}
	}
}

// statistka je helper function za izracun min, max in avg vrednosti za IzpisPorocila
func statistika(oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	p := postaje[oznaka].meritve

	if _, ok = postaje[oznaka]; !ok {
		return 0, 0, 0, ok
	}

	if len(p) == 0 || len(p) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja, najvisja, povprecna = p[0], p[0], 0

	for _, val := range p {
		if najnizja > val {
			najnizja = val
		}
		if najvisja < val {
			najvisja = val
		}
		povprecna += val
	}

	povprecna /= float64(len(p))

	return najnizja, najvisja, povprecna, ok
}
