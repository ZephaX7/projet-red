package statspersonnage

import (
	"fmt"
	personalisationpersonnage "projet-red/src/customcharacter"
)

// =============================
// Constantes d'équipement
// =============================

const (
	CasqueDeFer    = "Casque de fer"
	CasqueMagique  = "Casque magique"
	ArmureLegere   = "Armure légère"
	ArmureLourde   = "Armure lourde"
	BottesUsees    = "Bottes usées"
	BottesMagiques = "Bottes magiques"
)

// =============================
// Modèles
// =============================
type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Personnage struct {
	Nom       string
	Race      string
	Classe    string
	Sexe      string
	PvActuels int
	PvMax     int
	BasePvMax int // base avant bonus d'équipement (ex: 100)
	Equip     Equipment
}

// =============================
// Constructeurs
// =============================

// Nouveau crée un personnage "POO" avec PV cohérents.
func Nouveau(nom, race, classe, sexe string, pvActuels, pvMax int) *Personnage {
	base := pvMax
	if base <= 0 {
		base = 100
	}
	if pvActuels < 0 {
		pvActuels = 0
	}
	if pvActuels > base {
		pvActuels = base
	}
	return &Personnage{
		Nom:       nom,
		Race:      race,
		Classe:    classe,
		Sexe:      sexe,
		PvActuels: pvActuels,
		PvMax:     base,
		BasePvMax: base,
		Equip:     Equipment{},
	}
}

// NouveauDepuisCustom initialise à partir du package customcharacter
// en supposant qu'il expose une variable exportée `Personnage`.
func NouveauDepuisCustom() *Personnage {
	src := personalisationpersonnage.Personnage

	base := src.PvMax
	if base <= 0 {
		base = 100
	}
	actuels := src.PvActuels
	if actuels < 0 {
		actuels = 0
	}
	if actuels > base {
		actuels = base
	}

	return &Personnage{
		Nom:       src.Nom,
		Race:      src.Race,
		Classe:    src.Classe,
		Sexe:      src.Sexe,
		PvActuels: actuels,
		PvMax:     base,
		BasePvMax: base,
		Equip:     Equipment{},
	}
}

// Si ton package `customcharacter` n'expose PAS une variable mais une
// fonction (ex: GetPersonnage()), dis-le moi et je te fais l'adaptateur.

// =============================
// Méthodes "POO"
// =============================

// Soigner augmente les PV actuels sans dépasser PvMax.
func (p *Personnage) Soigner(soin int) {
	if soin <= 0 {
		return
	}
	p.PvActuels += soin
	if p.PvActuels > p.PvMax {
		p.PvActuels = p.PvMax
	}
}

// RecevoirDegats diminue les PV actuels sans passer sous 0.
func (p *Personnage) RecevoirDegats(degats int) {
	if degats <= 0 {
		return
	}
	p.PvActuels -= degats
	if p.PvActuels < 0 {
		p.PvActuels = 0
	}
}

// RecalculerPvMax recalcule PvMax à partir de BasePvMax + bonus équipement.
func (p *Personnage) RecalculerPvMax() {
	bonus := 0

	switch p.Equip.Head {
	case CasqueDeFer:
		bonus += 20
	case CasqueMagique:
		bonus += 50
	}

	switch p.Equip.Torso {
	case ArmureLegere:
		bonus += 30
	case ArmureLourde:
		bonus += 60
	}

	switch p.Equip.Feet {
	case BottesUsees:
		bonus += 5
	case BottesMagiques:
		bonus += 15
	}

	p.PvMax = p.BasePvMax + bonus
	if p.PvMax < 1 {
		p.PvMax = 1
	}
	if p.PvActuels > p.PvMax {
		p.PvActuels = p.PvMax
	}
}

// Equiper remplace tout l'équipement et recalcule les PV.
func (p *Personnage) Equiper(head, torso, feet string) {
	p.Equip.Head = head
	p.Equip.Torso = torso
	p.Equip.Feet = feet
	p.RecalculerPvMax()
}

// EquiperTete remplace uniquement la tête et recalcule les PV.
func (p *Personnage) EquiperTete(head string) {
	p.Equip.Head = head
	p.RecalculerPvMax()
}

// EquiperTorse remplace uniquement le torse et recalcule les PV.
func (p *Personnage) EquiperTorse(torso string) {
	p.Equip.Torso = torso
	p.RecalculerPvMax()
}

// EquiperPieds remplace uniquement les pieds et recalcule les PV.
func (p *Personnage) EquiperPieds(feet string) {
	p.Equip.Feet = feet
	p.RecalculerPvMax()
}

// AfficherStats affiche joliment les infos du perso.
func (p *Personnage) AfficherStats() {
	fmt.Println("Nom :", p.Nom)
	fmt.Println("Race :", p.Race)
	fmt.Println("Classe :", p.Classe)
	fmt.Println("Sexe :", p.Sexe)
	fmt.Println()
	fmt.Printf("PV : %d/%d\n", p.PvActuels, p.PvMax)
	fmt.Println()

	fmt.Println("Équipement :")
	if p.Equip.Head == "" {
		fmt.Println(" - Tête : rien")
	} else {
		fmt.Println(" - Tête :", p.Equip.Head)
	}

	if p.Equip.Torso == "" {
		fmt.Println(" - Torse : rien")
	} else {
		fmt.Println(" - Torse :", p.Equip.Torso)
	}

	if p.Equip.Feet == "" {
		fmt.Println(" - Pieds : rien")
	} else {
		fmt.Println(" - Pieds :", p.Equip.Feet)
	}
}
