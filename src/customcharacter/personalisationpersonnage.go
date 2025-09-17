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
    // Même en cas d'err, si on a lu qqch, on le traite quand même
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

func ChoixRace() int {
    fmt.Println("1. Humain (adapté à la classe guerrier)")
    fmt.Println("2. Elfe (adapté à la classe mage)")
    fmt.Println("3. Nain (adapté à la classe assassin)")
    fmt.Println()
}
    choix := lireChoixInt("Choisissez votre race (1, 2 ou 3) : ", 1, 3)
    switch choix {
    case 1:
        fmt.Println("Vous avez choisi la race Humain.")
    case 2:
        fmt.Println("Vous avez choisi la race Elfe.")
    case 3:
	}   

func ChoixClasse() int {
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Assassin")
	fmt.Println()

	var choix int
	fmt.Print("Choisissez votre classe (1, 2 ou 3) :")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		fmt.Println("Vous avez choisi le Guerrier.")
		ChoixSexe()
	case 2:
		fmt.Println("Vous avez choisi le Mage.")
		ChoixSexe()
	case 3:
		fmt.Println("Vous avez choisi l'Assassin.")
		ChoixSexe()
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
		return 0
	}
	return choix
}

func ChoixSexe() int {
	fmt.Println("1. Masculin")
	fmt.Println("2. Féminin")
	fmt.Println("3. Autre")
	fmt.Println()

	var choix int
	fmt.Print("Choisissez le sexe de votre personnage (1, 2 ou 3) :")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		fmt.Println("Vous avez choisi le sexe Masculin.")
		ChoisirNom()
	case 2:
		fmt.Println("Vous avez choisi le sexe Féminin.")
		ChoisirNom()
	case 3:
		fmt.Println("Vous avez choisi Autre.")
		ChoisirNom()
	default:
		fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
		return 0
	}
	return choix
}

func ChoisirNom() string {
	var nom string
	fmt.Print("Entrez le nom de votre personnage :")
	fmt.Scanln(&nom)

	var reponse string
	fmt.Print("Votre personnage s'appelle : ", nom, "\n")
	fmt.Println("Êtes-vous sûr ? (oui/non)")

	switch reponse {
	case "oui":
		fmt.Println("Bienvenue dans le monde de Aerthar (", nom, ").")
		if ChoixRace() == 1 {
			PointdeVie(1)
		}
		if ChoixRace() == 2 {
			PointdeVie(2)
		}
		if ChoixRace() == 3 {
			PointdeVie(3)
		}
		return nom

	default:
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
