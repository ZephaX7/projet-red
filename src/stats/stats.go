package statspersonnage

import "fmt"

//équipement

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

// Stats du personnage

type Personnage struct {
	Nom       string
	Race      string
	Classe    string
	Sexe      string
	PvActuels int
	PvMax     int
	Equip     Equipment
}

// Joueur

var Joueur = Personnage{
	Nom:       Personnage.Nom,
	Race:      Personnage.Race,
	Classe:    Personnage.Classe,
	Sexe:      Personnage.Sexe,
	PvActuels: Personnage.PvActuels,
	PvMax:     Personnage.PvMax,
	Equip: Equipment{
		Head:  "",
		Torso: "",
		Feet:  "",
	},
}

// Gestion des PV

func Soigner(soin int) {
	Joueur.PvActuels += soin
	if Joueur.PvActuels > Joueur.PvMax {
		Joueur.PvActuels = Joueur.PvMax
	}
}

func RecevoirDegats(degats int) {
	Joueur.PvActuels -= degats
	if Joueur.PvActuels < 0 {
		Joueur.PvActuels = 0
	}
}

// Application des bonus/malus d'équipement

func AppliquerEffetsEquipement() {
	basePvMax := 100

	// Vérifie ce que le joueur porte
	switch Joueur.Equip.Head {
	case "Casque de fer":
		basePvMax += 20
	case "Casque magique":
		basePvMax += 50
	}

	switch Joueur.Equip.Torso {
	case "Armure légère":
		basePvMax += 30
	case "Armure lourde":
		basePvMax += 60
	}

	switch Joueur.Equip.Feet {
	case "Bottes usées":
		basePvMax += 5
	case "Bottes magiques":
		basePvMax += 15
	}

	// Met à jour les PV max
	Joueur.PvMax = basePvMax

	// Ajuste les PV actuels si besoin
	if Joueur.PvActuels > Joueur.PvMax {
		Joueur.PvActuels = Joueur.PvMax
	}
}

// Affichage

func AfficherStats() {
	fmt.Println("Nom :", Joueur.Nom)

	fmt.Println("Race :", Joueur.Race)

	fmt.Println("Classe :", Joueur.Classe)

	fmt.Println("Sexe :", Joueur.Sexe)
	fmt.Println()
	fmt.Printf("PV : %d/%d\n", Joueur.PvActuels, Joueur.PvMax)
	fmt.Println()

	fmt.Println("Équipement :")
	if Joueur.Equip.Head == "" {
		fmt.Println(" - Tête : rien")

	} else {
		fmt.Println(" - Tête :", Joueur.Equip.Head)

	}

	if Joueur.Equip.Torso == "" {
		fmt.Println(" - Torse : rien")

	} else {
		fmt.Println(" - Torse :", Joueur.Equip.Torso)

	}

	if Joueur.Equip.Feet == "" {
		fmt.Println(" - Pieds : rien")
	} else {
		fmt.Println(" - Pieds :", Joueur.Equip.Feet)
	}

}
