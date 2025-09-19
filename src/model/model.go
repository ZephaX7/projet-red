package model

import "fmt"

// Types de race, classe et sexe
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

// Personnage représente le joueur
type Personnage struct {
	Nom       string
	Race      Race
	Classe    Classe
	Sexe      Sexe
	PVMax     int
	PVActuels int
	Gold      int
	Xp        int
	Revived   bool
	Skills    []string        // compétences
	Equip     map[string]bool // objets équipés
}

// Affichage lisible du personnage
func (p Personnage) Afficher() string {
	equip := ""
	for item, ok := range p.Equip {
		if ok {
			if equip != "" {
				equip += ", "
			}
			equip += item
		}
	}
	if equip == "" {
		equip = "(aucun)"
	}

	return fmt.Sprintf(
		"Nom : %s\nRace : %s\nClasse : %s\nSexe : %s\nPV : %d/%d\nGold : %d\nXp : %d\nSkills : %v\nÉquipement : %s",
		p.Nom, p.Race.String(), p.Classe.String(), p.Sexe.String(),
		p.PVActuels, p.PVMax, p.Gold, p.Xp, p.Skills, equip,
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

// InitCharacter crée un personnage avec PV selon race et Coup de poing comme compétence de base
func InitCharacter(nom string, race Race, classe Classe, sexe Sexe) *Personnage {
	var pvMax int
	switch race {
	case Humain:
		pvMax = 100
	case Elfe:
		pvMax = 80
	case Nain:
		pvMax = 120
	default:
		pvMax = 100
	}

	return &Personnage{
		Nom:       nom,
		Race:      race,
		Classe:    classe,
		Sexe:      sexe,
		PVMax:     pvMax,
		PVActuels: pvMax,
		Gold:      100,
		Xp:        0,
		Revived:   false,
		Skills:    []string{"Coup de poing"}, // compétence de base
		Equip:     make(map[string]bool),
	}
}

// EquipItem équipe un objet et applique ses bonus
func (p *Personnage) EquipItem(itemName string) {
	if p.Equip == nil {
		p.Equip = make(map[string]bool)
	}

	if p.Equip[itemName] {
		fmt.Println("⚠ Cet objet est déjà équipé :", itemName)
		return
	}

	switch itemName {
	case "Chapeau de l'aventurier":
		p.PVMax += 10
	case "Tunique de l'aventurier":
		p.PVMax += 25
	case "Bottes de l'aventurier":
		p.PVMax += 15
	default:
		fmt.Println("⚠ Cet objet ne peut pas être équipé :", itemName)
		return
	}

	p.Equip[itemName] = true
	p.PVActuels = p.PVMax // remplir les PV après équipement (optionnel)
	fmt.Println("✅ Vous avez équipé :", itemName)
	fmt.Printf("PV max actuel : %d\n", p.PVMax)
}
