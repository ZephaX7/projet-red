package customcharacter

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"projet-red/src/model"
)

type Personnalisateur struct {
	reader *bufio.Reader
}

func NewPersonnalisateur() *Personnalisateur {
	return &Personnalisateur{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (pz *Personnalisateur) lireLigne(prompt string) string {
	fmt.Print(prompt)
	s, _ := pz.reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func (pz *Personnalisateur) lireChoixInt(prompt string, min, max int) int {
	for {
		s := pz.lireLigne(prompt)
		n, err := strconv.Atoi(s)
		if err == nil && n >= min && n <= max {
			return n
		}
		fmt.Printf("Choix invalide (%d–%d).\n", min, max)
	}
}

func (pz *Personnalisateur) ChoixRace() model.Race {
	fmt.Println("1. Humain")
	fmt.Println("2. Elfe")
	fmt.Println("3. Nain")
	n := pz.lireChoixInt("Choisissez votre race : ", 1, 3)
	return model.Race(n)
}

func (pz *Personnalisateur) ChoixClasse() model.Classe {
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Assassin")
	n := pz.lireChoixInt("Choisissez votre classe : ", 1, 3)
	return model.Classe(n)
}

func (pz *Personnalisateur) ChoixSexe() model.Sexe {
	fmt.Println("1. Masculin")
	fmt.Println("2. Féminin")
	fmt.Println("3. Autre")
	n := pz.lireChoixInt("Choisissez votre sexe : ", 1, 3)
	return model.Sexe(n)
}

func (pz *Personnalisateur) ChoisirNom() string {
	for {
		nom := pz.lireLigne("Entrez le nom de votre personnage : ")
		if nom != "" {
			return nom
		}
		fmt.Println("Le nom ne peut pas être vide.")
	}
}

func (pz *Personnalisateur) CreerPersonnageInteractif() *model.Personnage {
	nom := pz.ChoisirNom()
	r := pz.ChoixRace()
	c := pz.ChoixClasse()
	s := pz.ChoixSexe()

	p := &model.Personnage{
		Nom:    nom,
		Race:   r,
		Classe: c,
		Sexe:   s,
	}

	// Init PV selon race
	switch p.Race {
	case model.Humain:
		p.PVMax, p.PVActuels = 100, 50
	case model.Elfe:
		p.PVMax, p.PVActuels = 80, 40
	case model.Nain:
		p.PVMax, p.PVActuels = 120, 60
	}

	return p
}

// Fonction publique pour main.go
func StartFlow() *model.Personnage {
	pz := NewPersonnalisateur()
	return pz.CreerPersonnageInteractif()
}
