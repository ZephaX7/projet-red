package menu

import (
	"fmt"
	"log"
	"os"
	"time"

	personalisationpersonnage "projet-red/src/customcharacter"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Fonction qui affiche le menu (texte + ascii art)
func afficherMenu() {
	fmt.Println("Bienvenu dans les Chroniques d’Aerthar — Édition Terminal")
	fmt.Println()

	fmt.Println()
	Lore, err := os.ReadFile("src/menu/Lore.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(Lore))

	// Lire le fichier ASCII art
	Menu, err := os.ReadFile("src/menu/asciimenu.txt")
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
	f, err := os.Open("src/menu/Bienvenue-en-Gaule.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Lancer en arrière-plan
	go func() {
		speaker.Play(streamer)
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
			Demarage, err := os.ReadFile("src/menu/asciidemarage.txt")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(Demarage))
			fmt.Println()
			menu = false
			personalisationpersonnage.ChoixRace()

		case 2:
			fmt.Println("Au revoir !")
			menu = false
			streamer.Close()

		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
