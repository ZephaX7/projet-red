package Boutique

import "fmt"

func boutique() {
	var shop string
	fmt.Println("Choisissez votre race (1, 2 ou 3) : ") //MODIFIER TOUS LES TRUCS DE PERSOS EN ACHAT
	fmt.Println("1. Humain (adapté à la classe guerrier)")
	fmt.Println("2. Elfe (adapté à la classe mage)")
	fmt.Println("3. Nain (adapté à la classe assassin)")
	fmt.Scanln(&shop)

	switch shop {
	case "1":
		fmt.Println("Vous avez choisi la race Humain.")
		return true
	case "2":
		fmt.Println("Vous avez choisi la race Elfe.")
		return true
	case "3":
		fmt.Println("Vous avez choisi la race Nain.")
		return true
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
		return false
	}
}
