package statspersonnage

import "fmt"

// ---------------------
// Modèle d'équipement
// ---------------------
// (Supprime cette struct si tu l'as déjà ailleurs dans le package)
type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

// ---------------------
// Affichage des stats
// ---------------------

// Affiche les statistiques principales du personnage
// pvActuels : points de vie actuels
// pvMax     : points de vie maximum
func AfficherStatsPersonnage(nom, race, classe, sexe string, PvActuels, PvMax int) {
	// garde-fous pour un affichage correct
	if PvMax < 0 {
		PvMax = 0
	}
	if PvActuels < 0 {
		PvActuels = 0
	}
	if PvActuels > PvMax {
		PvActuels = PvMax
	}

	fmt.Println("Nom du personnage :", nom)
	fmt.Println("Race du personnage :", race)
	fmt.Println("Classe du personnage :", classe)
	fmt.Println("Sexe du personnage :", sexe)
	fmt.Printf("Points de vie : %d/%d\n", PvActuels, PvMax)
}

// ---------------------
// Gestion des PV
// ---------------------

// RecevoirDegats applique des dégâts et, si le perso "meurt" (PV <= 0),
// il revient automatiquement à 50% de ses PV max.
// Retourne : nouveauxPV, mort(bool), ressuscite(bool)
func RecevoirDegats(PvActuels, PvMax, degats int) (int, bool, bool) {
	if PvMax <= 0 {
		PvMax = 1
	}
	if degats <= 0 {
		// Pas de dégâts (pour soigner, utilise Soigner)
		return clampPV(PvActuels, PvMax), false, false
	}

	PvActuels -= degats
	if PvActuels > 0 {
		return PvActuels, false, false
	}

	// Mort -> rez à 50% des PV max (arrondi inférieur).
	// 👉 Si tu préfères arrondir au supérieur, remplace par : half := (pvMax + 1) / 2
	half := PvMax / 2
	if half < 1 {
		half = 1
	}
	return half, true, true
}

// Soigner augmente les PV sans dépasser le max
func Soigner(PvActuels, PvMax, soin int) int {
	if PvMax <= 0 {
		PvMax = 1
	}
	PvActuels += soin
	if PvActuels > PvMax {
		PvActuels = PvMax
	}
	if PvActuels < 0 {
		PvActuels = 0
	}
	return PvActuels
}

func clampPV(PvActuels, PvMax int) int {
	if PvMax <= 0 {
		PvMax = 1
	}
	if PvActuels < 0 {
		return 0
	}
	if PvActuels > PvMax {
		return PvMax
	}
	return PvActuels
}

// ---------------------
// Affichage équipement
// ---------------------

func AfficherEquipementTete(equipement string) {
	fmt.Println("Équipement de tête :", equipement)
}

func AfficherEquipementTorse(equipement string) {
	fmt.Println("Équipement du torse :", equipement)
}

func AfficherEquipementPieds(equipement string) {
	fmt.Println("Équipement des pieds :", equipement)
}

// Affiche l'équipement complet du personnage à partir de la structure Equipment
func AfficherEquipementPersonnage(equipement Equipment) {
	fmt.Println("Équipement du personnage :")
	AfficherEquipementTete(equipement.Head)
	AfficherEquipementTorse(equipement.Torso)
	AfficherEquipementPieds(equipement.Feet)
}
