package hub

import (
	"fmt"
	"log"
	"os"
	"projet-red/src/charactermenu"
	"projet-red/src/fight"
	"projet-red/src/model"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func musiqueAmbiance() (beep.StreamSeekCloser, beep.Format) {
	f, err := os.Open("asset/Ambiance.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Créer un effet de volume
	volume := &effects.Volume{
		Streamer: streamer,
		Base:     2.8, // base logarithmique
		Volume:   -1,  // diminue le volume de 5 dB
		Silent:   false,
	}

	// Lancer en arrière-plan
	go func() {
		speaker.Play(volume)
	}()

	return streamer, format
}

func Hub(perso *model.Personnage) {
	streamer, _ := musiqueAmbiance()
	gameOver := false

	for !gameOver {
		fmt.Println("\n🏰 Bienvenue au Hub ! 🏰")
		fmt.Println("1 - Partir à l'aventure")
		fmt.Println("2 - Ouvrir le menu")
		fmt.Println("3 - Quitter")

		var choix int
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			ennemi := model.RandomEnnemi()
			fmt.Printf("\nUn ennemi apparaît : %s ! PV : %d/%d\n", ennemi.Nom, ennemi.PVActuels, ennemi.PVMax)

			fight.Combat(perso, ennemi)

			if perso.PVActuels <= 0 {
				if !perso.Revived {
					perso.PVActuels = perso.PVMax / 2
					perso.Revived = true
					fmt.Printf("\n💀 Vous êtes mort mais ressuscité ! PV : %d/%d\n", perso.PVActuels, perso.PVMax)
				} else {
					fmt.Println("\n💀 Vous êtes mort pour de bon ! Game Over.")
					gameOver = true // stoppe le hub
				}
			}

		case 2:
			charactermenu.AfficherMenu(perso)
		case 3:
			fmt.Println("Au revoir !")
			streamer.Close()
			gameOver = true
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
