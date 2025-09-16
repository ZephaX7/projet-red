package PersonalisationPersonnage

import (
	"fmt"
)

func ChoixRace() bool {
	var choix string
	fmt.Println("Choisissez votre race (1, 2 ou 3) : ")
	fmt.Println("1. Humain (adapté à la classe guerrier)")
	fmt.Println("2. Elfe (adapté à la classe mage)")
	fmt.Println("3. Nain (adapté à la classe assassin)")
	fmt.Scanln(&choix)

	switch choix {
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

func ChoixClasse() bool {
	var choix string
	fmt.Println("Choisissez votre classe (1, 2 ou 3) :")
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Assassin")
	fmt.Scanln(&choix)

	switch choix {
	case "1":
		fmt.Println("Vous avez choisi le Guerrier.")
		return true
	case "2":
		fmt.Println("Vous avez choisi le Mage.")
		return true
	case "3":
		fmt.Println("Vous avez choisi l'Assassin.")
		return true
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
		return false
	}
}

func ChoixSexe() bool {
	var choix string
	fmt.Println("Choisissez le sexe de votre personnage (1 ou 2 ou 3) :")
	fmt.Println("1. Masculin")
	fmt.Println("2. Féminin")
	fmt.Println("3. Autre")
	fmt.Scanln(&choix)

	switch choix {
	case "1":
		fmt.Println("Vous avez choisi le sexe Masculin.")
		return true
	case "2":
		fmt.Println("Vous avez choisi le sexe Féminin.")
		return true
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1 ou 2.")
		return false
	}
}

func Choisir_Nom() string {
	var nom string
	fmt.Println("Entrez le nom de votre personnage :")
	fmt.Scanln(&nom)
	fmt.Print("Votre personnage s'appelle : ", nom, "\n")
	fmt.Println("Etes vous sur? (oui/non)")
	var reponse string
	fmt.Scanln(&reponse)
	if reponse == "oui" {
		fmt.Println("Bienvenue dans le monde de Aerthar (", nom, ").")
		return nom
	} else {
		fmt.Println("Veuillez choisir un autre nom.")
		return Choisir_Nom()
	}
}
func PointdeVie() {
	race := 0
	switch race {
	case "1":
		fmt.Println("Vous avez 50 points de vie actuellement, et vous pouvez en avoir jusqu'à 100 étant humain.")
	case "2":
		fmt.Println("Vous avez 40 points de vie actuellement, et vous pouvez en avoir jusqu'à 80 étant elfe.")
	case "3":
		fmt.Println("Vous avez 60 points de vie actuellement, et vous pouvez en avoir jusqu'à 120 étant nain.")

	}
}
func Choisir_Nom() string {
	var nom string
	fmt.Println("Entrez le nom de votre personnage :")
	fmt.Scanln(&nom)
	fmt.Print("Votre personnage s'appelle : ", nom, "\n")
	fmt.Println("Etes vous sur? (oui/non)")
	var reponse string
	fmt.Scanln(&reponse)
	if reponse == "oui" {
		fmt.Println("Bienvenue dans le monde de Aerthar (", nom, ").")
		return nom
	} else {
		fmt.Println("Veuillez choisir un autre nom.")
		return Choisir_Nom()
	}
}
