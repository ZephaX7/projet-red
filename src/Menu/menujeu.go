package menu

import (
	"fmt"
	"log"
	"os"
	"time"

	PersonalisationPersonnage "github.com/ZephaX7/projet-red/src/customcharacter"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func menu() {

	menu := true

	fmt.Println("Bienvenu dans les Chroniques d’Aerthar — Édition Terminal")

	fmt.Println()

	// Lire le fichier ASCII art
	Menu, err := os.ReadFile("asciimenu.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(Menu))

	fmt.Println("Que voulez-vous faire ?")

	fmt.Println()

	fmt.Println("1 - Jouer")
	fmt.Println("2 - Afficher le Lore")
	fmt.Println("3 - Quitter")

	fmt.Println()

	// --- Lecture musique ---
	f, err := os.Open("Bienvenue-en-Gaule.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Musique en arrière-plan
	go func() {
		speaker.Play(streamer)
	}()

	// --- Intro affichée UNE seule fois ---
	fmt.Println("Bienvenu dans les Chroniques d’Aerthar — Édition Terminal")
	fmt.Println()

	Menu2, err := os.ReadFile("asciimenu.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(Menu2))

	// --- Boucle principale ---
	if menu {

		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1 - Jouer")
		fmt.Println("2 - Afficher le Lore")
		fmt.Println("3 - Quitter")
		fmt.Println()

		var choix int
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			Demarage, err := os.ReadFile("asciidemarage.txt")
			//police small shadow
			if err != nil {
				panic(err)
			}
			fmt.Println(string(Demarage))
			fmt.Println()
			menu = false
			PersonalisationPersonnage.ChoixRace()

		case 2:
			fmt.Println("Ouverture du Lore...")
			fmt.Println()
			Lore, err := os.ReadFile("Lore.txt")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(Lore))

			f, err := os.Open("Sauron.mp3")
			if err != nil {
				log.Fatal(err)
			}
			streamer, format, err := mp3.Decode(f)
			if err != nil {
				log.Fatal(err)
			}

			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

			done := make(chan bool)
			speaker.Play(beep.Seq(streamer, beep.Callback(func() {
				done <- true
			})))

		case 3:
			fmt.Println("Au revoir !")
			menu = false
			streamer.Close()
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}

}
