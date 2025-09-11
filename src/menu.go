<<<<<<< HEAD
package main

import "fmt"

// Fonction qui affiche le menu
func menu() {
	fmt.Println("Bienvenu dans les Chroniques d’Aerthar — Édition Terminal")

	fmt.Println()

	fmt.Println("Que voulez-vous faire ?")

	fmt.Println()

	fmt.Println("1 - Jouer")
	fmt.Println("2 - Paramètre")
	fmt.Println("3 - Quitter")

	fmt.Println()

}

func main() {
	continuer := true

	for continuer {
		menu()

		var choix int
		fmt.Print("Entrez votre choix : ")
		fmt.Println()
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("Démarrage du jeu...")
			fmt.Println()
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
=======
package menu

import "fmt"

// Fonction qui affiche le menu
func menu() {
	fmt.Println("Bienvenu dans Chroniques d’Aerthar — Édition Terminal")
	fmt.Println("Que voulez-vous faire ?")
	fmt.Println("1 - Jouer")
	fmt.Println("2 - Paramètres")
	fmt.Println("3 - Quitter")
}

func main() {
	continuer := true

	for continuer {
		menu()

		var choix int
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("Démarrage du jeu...")
			// Ici tu pourras lancer la fonction du jeu
		case 2:
			fmt.Println("Ouverture des paramètres...")
			// Ici tu pourras gérer les paramètres
		case 3:
			fmt.Println("Au revoir !")
			continuer = false
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
>>>>>>> 175a497f77f5fdad4a9c05944b7ebad9fdf6a24f
