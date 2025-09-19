package menu

import (
	"fmt"
	"log"
	"os"
	"time"

	"projet-red/src/charactermenu"
	personalisationpersonnage "projet-red/src/customcharacter"
	"projet-red/src/hub"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Fonction qui affiche le menu (texte + ascii art)
func afficherMenu() {
	fmt.Println("Bienvenue dans les Chroniques d’Aerthar — Édition Terminal")
	fmt.Println()

	fmt.Println()
	Lore, err := os.ReadFile("asset/Lore.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(Lore))

	// Lire le fichier ASCII art
	Menu, err := os.ReadFile("asset/asciimenu.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(Menu))

	fmt.Println("Que voulez-vous faire ?")
	fmt.Println()
	fmt.Println("1 - Jouer")
	fmt.Println("2 - Quitter")
	fmt.Println()
}

// Fonction pour lancer la musique d'accueil
func musiqueAccueil() (beep.StreamSeekCloser, beep.Format) {
	f, err := os.Open("asset/Bienvenue-en-Gaule.mp3")
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
		Base:     2.8,  // base logarithmique
		Volume:   -1.5, // diminue le volume de 5 dB
		Silent:   false,
	}

	// Lancer en arrière-plan
	go func() {
		speaker.Play(volume)
	}()

	return streamer, format
}

// Fonction principale du menu
func RunMenu() {
	// Lance la musique d'accueil
	streamer, _ := musiqueAccueil()

	// Affiche le menu une première fois
	afficherMenu()

	menu := true
	for menu {
		var choix int
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			Demarage, err := os.ReadFile("asset/asciidemarage.txt")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(Demarage))
			fmt.Println()
			menu = false
			// 🔽 Création du personnage et lancement du jeu
			perso := personalisationpersonnage.StartFlow()
			charactermenu.Menu_character(perso)
			hub.Hub(perso)

		case 2:
			fmt.Println("Au revoir !")
			menu = false
			streamer.Close()

		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
