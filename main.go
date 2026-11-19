package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/1pajaco/vremenske_postaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "vremenskaPostaja",
		Usage: "Beleži meritve vremenskih postaj in izpiše poročilo",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura meritve (°C)",
				Value: -50,
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura meritve (°C)",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			vreme.NastaviVhodne(
				cmd.Float("minTemp"),
				cmd.Float("maxTemp"),
				int(cmd.Int("stMeritev")),
			)
			zagon()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func zagon() {
	vreme.DodajPostajo("KP01", "Koper Luka", "Koper")
	vreme.DodajPostajo("KR01", "Krederica", "Triglav")
	vreme.DodajPostajo("LJ01", "Ljubljana Bežigrad", "Ljubljana")
	vreme.DodajPostajo("MB01", "Maribor Tabor", "Maribor")

	vreme.DodajMeritev("w", 24.5)

	vreme.DodajMeritev("KP01", -50)

	vreme.DodajMeritev("KP01", 50)

	vreme.DodajMeritev("KP01", 24.5)
	vreme.DodajMeritev("KP01", 27)
	vreme.DodajMeritev("KP01", 26.1)
	vreme.DodajMeritev("KP01", 28.4)

	vreme.DodajMeritev("KR01", 2)
	vreme.DodajMeritev("KR01", -1.5)
	vreme.DodajMeritev("KR01", 3.5)

	vreme.DodajMeritev("LJ01", 18.2)
	vreme.DodajMeritev("LJ01", 21.5)
	vreme.DodajMeritev("LJ01", 19.8)
	vreme.DodajMeritev("LJ01", 16.5)

	vreme.DodajMeritev("MB01", 12.4)

	vreme.IzpisMeritev()
	fmt.Println("---------------------------------------------------------------------------")
	vreme.IzpisPorocila()
	fmt.Println("---------------------------------------------------------------------------")

}
