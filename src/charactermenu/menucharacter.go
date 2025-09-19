package charactermenu

import (
	"fmt"
	"log"
	"os"
	combat "projet-red/src/fight"
	"projet-red/src/forgeron"
	"projet-red/src/inventory"
	"projet-red/src/items"
	"projet-red/src/model"
	"projet-red/src/shop"
	"projet-red/src/stats"
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

func Menu_character(perso *model.Personnage) {
	streamer, _ := musiqueAmbiance()
	for {
		fmt.Println("1- Statistiques")
		fmt.Println("2- Inventaire")
		fmt.Println("3- Boutique")
		fmt.Println("4- Forge")
		fmt.Println("5- Retour")
		fmt.Println()

		var Menu int
		fmt.Println("Entrez votre choix :")
		fmt.Scanln(&Menu)

		switch Menu {
		case 1:
			stats.AfficherStats(perso)
		case 2:
			inventory.AccessInventory(perso, false, nil, items.UtiliserObjet)
		case 3:
			shop.Shop(perso)
		case 4:
			forgeron.Shop(perso)
		case 5:
			fmt.Println("En avant vers l'aventure !")
			ennemi := model.RandomEnnemi() // ✅ maintenant r est défini
			combat.Combat(perso, ennemi)
			streamer.Close()
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func AfficherMenu(p *model.Personnage) {
	Menu_character(p)
}
