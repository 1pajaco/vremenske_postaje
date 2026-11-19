# vremenske_postaje

Paket `vreme` omogoča beleženje meritev temperature na vremenskih postajah in izpis poročil o njih.

Postaje se dodajo s funkcijo `DodajPostajo`, meritve s funkcijo `DodajMeritev`. Funkciji `IzpisMeritev` in `IzpisPorocila` izpišeta vse meritve oziroma poročilo z najnižjo, najvišjo in povprečno temperaturo ter oceno (Mrzlo, Hladno, Toplo, Vroče).

Primer uporabe:

```go
vreme.NastaviVhodne(-50, 50, 3)
vreme.DodajPostajo("LJ01", "Ljubljana Bežigrad", "Ljubljana")
vreme.DodajMeritev("LJ01", 18.2)
vreme.DodajMeritev("LJ01", 21.5)
vreme.DodajMeritev("LJ01", 19.8)

vreme.IzpisMeritev()
vreme.IzpisPorocila()
```

Namestitev:

```
go get github.com/name/vremenske_postaje@v1.0.0
```

Aplikacija ima tri stikala:

* `minTemp`: najnižja dovoljena temperatura meritve (privzeto -50 °C)
* `maxTemp`: najvišja dovoljena temperatura meritve (privzeto 50 °C)
* `stMeritev`: najmanjše število meritev za izračun statistike (privzeto 3)

Zagon:

```
go run . --minTemp -10 --maxTemp 40 --stMeritev 2
```