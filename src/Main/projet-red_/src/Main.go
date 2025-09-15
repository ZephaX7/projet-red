package main

import (
	"fmt"
	//"src/Custom_Character/Personalisation_Personnage"
	"os"
)

// Fonction qui affiche le menu
func menu() {

	fmt.Println("Bienvenu dans les Chroniques d’Aerthar — Édition Terminal")

	fmt.Println()

	// Lire le fichier ASCII art
	data, err := os.ReadFile("ascii_menu.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Afficher le contenu
	fmt.Println(string(data))

	fmt.Println("Que voulez-vous faire ?")

	fmt.Println()

	fmt.Println("1 - Jouer")
	fmt.Println("2 - Paramètre")
	fmt.Println("3 - Quitter")

	fmt.Println()

}

func main() {
	continuer := true
	//Open_Menu := false

	for continuer {
		menu()

		var Menu string
		fmt.Print("Ecrivez : Paramètre , a tout moment pour ouvrir le Menu !")
		fmt.Println()
		fmt.Scan(&Menu)

		var choix int
		fmt.Print("Entrez votre choix : ")
		fmt.Println()
		fmt.Scan(&choix)

		switch Menu {
		case "Paramètre":
			fmt.Println("ouverture des paramètres...")
			fmt.Println()
			//Open_Menu = true
			//Menu_Character.Menu_character(Open_Menu)

		default:

		}

		switch choix {
		case 1:
			fmt.Println("Démarrage du jeu...")
			fmt.Println()

			//Personalisation_Personnage.ChoixRace()

			// Ici tu pourras lancer la fonction du jeu
		case 2:
			fmt.Println("Ouverture des paramètres...")
			fmt.Println()
			// Ici tu pourras gérer les paramètres
		case 3:
			fmt.Println("Au revoir !")
			continuer = false
		default:
			fmt.Println("Choix invalide, réessayez.")

		}
	}
}
