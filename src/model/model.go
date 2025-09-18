package model

import "fmt"

type Race int
type Classe int
type Sexe int

const (
	Humain Race = iota + 1
	Elfe
	Nain
)

const (
	Guerrier Classe = iota + 1
	Mage
	Assassin
)

const (
	Masculin Sexe = iota + 1
	Feminin
	Autre
)

type Personnage struct {
	Nom       string
	Race      Race
	Classe    Classe
	Sexe      Sexe
	PVActuels int
	PVMax     int
}

// Affichage lisible
func (p Personnage) Afficher() string {
	return fmt.Sprintf(
		"Nom : %s\nRace : %s\nClasse : %s\nSexe : %s\nPV : %d/%d\n",
		p.Nom, p.Race, p.Classe, p.Sexe, p.PVActuels, p.PVMax,
	)
}

// Conversions en texte
func (r Race) String() string {
	switch r {
	case Humain:
		return "Humain"
	case Elfe:
		return "Elfe"
	case Nain:
		return "Nain"
	default:
		return "Inconnue"
	}
}

func (c Classe) String() string {
	switch c {
	case Guerrier:
		return "Guerrier"
	case Mage:
		return "Mage"
	case Assassin:
		return "Assassin"
	default:
		return "Inconnue"
	}
}

func (s Sexe) String() string {
	switch s {
	case Masculin:
		return "Masculin"
	case Feminin:
		return "Féminin"
	case Autre:
		return "Autre"
	default:
		return "Inconnu"
	}
}
