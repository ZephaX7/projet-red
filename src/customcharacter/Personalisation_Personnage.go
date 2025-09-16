package PersonalisationPersonnage

import (
	"fmt"
)

func ChoixRace() int {
	var choix int
	fmt.Println("Choisissez votre race (1, 2 ou 3) : ")
	fmt.Println("1. Humain (adapté à la classe guerrier)")
	fmt.Println("2. Elfe (adapté à la classe mage)")
	fmt.Println("3. Nain (adapté à la classe assassin)")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		fmt.Println("Vous avez choisi la race Humain.")
	case 2:
		fmt.Println("Vous avez choisi la race Elfe.")
	case 3:
		fmt.Println("Vous avez choisi la race Nain.")
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
		return 0
	}
	return choix
}

func ChoixClasse() int {
	var choix int
	fmt.Println("Choisissez votre classe (1, 2 ou 3) :")
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Assassin")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		fmt.Println("Vous avez choisi le Guerrier.")
	case 2:
		fmt.Println("Vous avez choisi le Mage.")
	case 3:
		fmt.Println("Vous avez choisi l'Assassin.")
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
		return 0
	}
	return choix
}

func ChoixSexe() int {
	var choix int
	fmt.Println("Choisissez le sexe de votre personnage (1, 2 ou 3) :")
	fmt.Println("1. Masculin")
	fmt.Println("2. Féminin")
	fmt.Println("3. Autre")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		fmt.Println("Vous avez choisi le sexe Masculin.")
	case 2:
		fmt.Println("Vous avez choisi le sexe Féminin.")
	case 3:
		fmt.Println("Vous avez choisi Autre.")
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
		return 0
	}
	return choix
}

func ChoisirNom() string {
	var nom string
	fmt.Println("Entrez le nom de votre personnage :")
	fmt.Scanln(&nom)
	fmt.Print("Votre personnage s'appelle : ", nom, "\n")
	fmt.Println("Êtes-vous sûr ? (oui/non)")
	var reponse string
	fmt.Scanln(&reponse)
	if reponse == "oui" {
		fmt.Println("Bienvenue dans le monde de Aerthar (", nom, ").")
		return nom
	} else {
		fmt.Println("Veuillez choisir un autre nom.")
		return ChoisirNom()
	}
}

func PointdeVie(race int) {
	switch race {
	case 1:
		fmt.Println("Vous avez 50 points de vie actuellement, et vous pouvez en avoir jusqu'à 100 étant humain.")
	case 2:
		fmt.Println("Vous avez 40 points de vie actuellement, et vous pouvez en avoir jusqu'à 80 étant elfe.")
	case 3:
		fmt.Println("Vous avez 60 points de vie actuellement, et vous pouvez en avoir jusqu'à 120 étant nain.")
	default:
		fmt.Println("Race inconnue")
	}
}
