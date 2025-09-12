package Boutique

import "fmt"

func boutique() {
	var shop string
	fmt.Println("Bienvenue,")
	fmt.Println()
	fmt.Println("Que désirez-vous ?")
	fmt.Println()
	fmt.Println("Magasin (choisir par le chiffre ou le numéro):")
	fmt.Println()
	fmt.Println("0  - Potion de vie gratuite(C'est pour toi bg,cadeau de la maison. Par contre si tu l'uses pour rien, je vais te retrouver donc fais gaffe à toi)")
	fmt.Println("1  - Potion de vie")
	fmt.Println("2  - Potion de poison")
	fmt.Scanln(&shop)

	switch shop {
	case "0":
		fmt.Println("Voulez vraiment le récupérer")
		return true
	case "1":
		fmt.Println("Vous avez choisi la race Elfe.")
		return true
	case "2":
		fmt.Println("Vous avez choisi la race Nain.")
		return true
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
		return false
	}
}
