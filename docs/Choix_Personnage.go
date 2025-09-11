package main

import "fmt"

func ChoixPersonnage() bool {
	var choix string
	fmt.Println("Choisissez votre personnage (1, 2 ou 3) :")
	fmt.Println("1. Guerrier Géant")
	fmt.Println("2. Nain mage")
	fmt.Println("3. Elfe Assassin")
	fmt.Scanln(&choix)

	switch choix {
	case "1":
		fmt.Println("Vous avez choisi le Guerrier Géant.")
		return true
	case "2":
		fmt.Println("Vous avez choisi le Nain mage.")
		return true
	case "3":
		fmt.Println("Vous avez choisi l'Elfe Assassin.")
		return true
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
		return false
	}
}
