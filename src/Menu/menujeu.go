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

// Fonction qui affiche le menu (texte + ascii art)
func afficherMenu() {
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
}

// Fonction pour lancer la musique d'accueil
func musiqueAccueil() (beep.StreamSeekCloser, beep.Format) {
	f, err := os.Open("Bienvenue-en-Gaule.mp3")
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

<<<<<<< HEAD:src/Menu/menujeu.go
	Menu2, err := os.ReadFile("asciimenu.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(Menu2))
=======
// Fonction principale du menu
func RunMenu() {
	// Lance la musique d'accueil
	streamer, _ := musiqueAccueil()
>>>>>>> e74c50a6a1bb9fec111fd8a611e78bbb0881c226:src/Menu/menujeu

	// Affiche le menu une première fois
	afficherMenu()

	menu := true
	for menu {
		var choix int
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
<<<<<<< HEAD:src/Menu/menujeu.go
			Demarage, err := os.ReadFile("asciidemarage.txt")
			//police small shadow
=======
			Demarage, err := os.ReadFile("ascii_Demarage.txt")
>>>>>>> e74c50a6a1bb9fec111fd8a611e78bbb0881c226:src/Menu/menujeu
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

			// 🔥 Arrêter la musique d'accueil
			streamer.Close()

			// Charger et jouer la musique du Lore
			f2, err := os.Open("GilraensMemorial.mp3")
			if err != nil {
				log.Fatal(err)
			}
			streamer2, _, err := mp3.Decode(f2)
			if err != nil {
				log.Fatal(err)
			}

			go func() {
				speaker.Play(streamer2)
			}()

			afficherMenu()

		case 3:
			fmt.Println("Au revoir !")
			menu = false
			streamer.Close()

		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
