package personalisationpersonnage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// --------- Utilitaires de lecture ----------

func lireLigne(prompt string) (string, error) {
	fmt.Print(prompt)
	s, err := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s, err
}

func lireChoixInt(prompt string, min, max int) int {
	for {
		s, _ := lireLigne(prompt)
		if s == "" {
			// Ligne vide -> on redemande sans bruit
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil || n < min || n > max {
			if min == max {
				fmt.Printf("Choix invalide. Veuillez choisir %d.\n", min)
			} else {
				fmt.Printf("Choix invalide. Veuillez choisir un nombre entre %d et %d.\n", min, max)
			}
			continue
		}
		return n
	}
}

// ------------- Menus -----------------------

func ChoixRace() string {
	fmt.Println("1. Humain (adapté à la classe guerrier)")
	fmt.Println("2. Elfe (adapté à la classe mage)")
	fmt.Println("3. Nain (adapté à la classe assassin)")
	fmt.Println()

	race := ""
	choix := lireChoixInt("Choisissez votre race (1, 2 ou 3) : ", 1, 3)
	switch choix {
	case 1:
		fmt.Println("Vous avez choisi la race Humain.")
		race = "Humain"
	case 2:
		fmt.Println("Vous avez choisi la race Elfe.")
		race = "Elfe"
	case 3:
		fmt.Println("Vous avez choisi la race Nain.")
		race = "Nain"
	}
	return race
}

func ChoixClasse() string {
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Assassin")
	fmt.Println()

	class := ""
	choix := lireChoixInt("Choisissez votre classe (1, 2 ou 3) : ", 1, 3)
	switch choix {
	case 1:
		fmt.Println("Vous avez choisi le Guerrier.")
		class = "Guerrier"
	case 2:
		fmt.Println("Vous avez choisi le Mage.")
		class = "Mage"
	case 3:
		fmt.Println("Vous avez choisi l'Assassin.")
		class = "Assassin"
	}
	return class
}

func ChoixSexe() string {
	fmt.Println("1. Masculin")
	fmt.Println("2. Féminin")
	fmt.Println("3. Autre")
	fmt.Println()

	sexe := ""
	choix := lireChoixInt("Choisissez le sexe de votre personnage (1, 2 ou 3) : ", 1, 3)
	switch choix {
	case 1:
		fmt.Println("Vous avez choisi le sexe Masculin.")
		sexe = "homme"
	case 2:
		fmt.Println("Vous avez choisi le sexe Féminin.")
		sexe = "femme"
	case 3:
		fmt.Println("Vous avez choisi Autre.")
		sexe = "autre"
	}
	return sexe
}

func ChoisirNom() string {
	for {
		nom, _ := lireLigne("Entrez le nom de votre personnage : ")
		if nom == "" {
			fmt.Println("Le nom ne peut pas être vide.")
			continue
		}

		fmt.Printf("Votre personnage s'appelle : %s\n", nom)
		reponse, _ := lireLigne("Êtes-vous sûr ? (oui/non) : ")
		switch strings.ToLower(strings.TrimSpace(reponse)) {
		case "oui", "o", "y", "yes":
			fmt.Printf("Bienvenue dans le monde de Aerthar (%s).\n", nom)
			// ❌ Surtout PAS d'appel à CreerPersonnageInteractif() ici !
			return nom
		case "non", "n", "no":
			fmt.Println("D'accord, recommençons.")
			continue
		default:
			fmt.Println("Réponse non reconnue. Tape 'oui' ou 'non'.")
			continue
		}
	}
}

// ----------------- PV par race -----------------

func PointdeVie(race string) {
	switch race {
	case "Humain":
		fmt.Println("Vous avez 50 points de vie actuellement, et vous pouvez en avoir jusqu'à 100 étant humain.")
	case "Elfe":
		fmt.Println("Vous avez 40 points de vie actuellement, et vous pouvez en avoir jusqu'à 80 étant elfe.")
	case "Nain":
		fmt.Println("Vous avez 60 points de vie actuellement, et vous pouvez en avoir jusqu'à 120 étant nain.")
	default:
		fmt.Println("Race inconnue")
	}
}

// ------------- Orchestrateur ------------------

// CreerPersonnageInteractif lance le flux complet et affiche les PV selon la race.
func CreerPersonnageInteractif() (nom, race, classe, sexe string) {
	race = ChoixRace()
	classe = ChoixClasse()
	sexe = ChoixSexe()
	nom = ChoisirNom()

	// Afficher les PV initiaux en fonction de la race
	PointdeVie(race)

	return nom, race, classe, sexe
}
