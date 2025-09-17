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
func AfficherStatsPersonnage(nom, race, classe, sexe string, pvActuels, pvMax int) {
	// garde-fous pour un affichage correct
	if pvMax < 0 {
		pvMax = 0
	}
	if pvActuels < 0 {
		pvActuels = 0
	}
	if pvActuels > pvMax {
		pvActuels = pvMax
	}

	fmt.Println("Nom du personnage :", nom)
	fmt.Println("Race du personnage :", race)
	fmt.Println("Classe du personnage :", classe)
	fmt.Println("Sexe du personnage :", sexe)
	fmt.Printf("Points de vie : %d/%d\n", pvActuels, pvMax)
}

// ---------------------
// Gestion des PV
// ---------------------

// RecevoirDegats applique des dégâts et, si le perso "meurt" (PV <= 0),
// il revient automatiquement à 50% de ses PV max.
// Retourne : nouveauxPV, mort(bool), ressuscite(bool)
func RecevoirDegats(pvActuels, pvMax, degats int) (int, bool, bool) {
	if pvMax <= 0 {
		pvMax = 1
	}
	if degats <= 0 {
		// Pas de dégâts (pour soigner, utilise Soigner)
		return clampPV(pvActuels, pvMax), false, false
	}

	pvActuels -= degats
	if pvActuels > 0 {
		return pvActuels, false, false
	}

	// Mort -> rez à 50% des PV max (arrondi inférieur).
	// 👉 Si tu préfères arrondir au supérieur, remplace par : half := (pvMax + 1) / 2
	half := pvMax / 2
	if half < 1 {
		half = 1
	}
	return half, true, true
}

// Soigner augmente les PV sans dépasser le max
func Soigner(pvActuels, pvMax, soin int) int {
	if pvMax <= 0 {
		pvMax = 1
	}
	pvActuels += soin
	if pvActuels > pvMax {
		pvActuels = pvMax
	}
	if pvActuels < 0 {
		pvActuels = 0
	}
	return pvActuels
}

func clampPV(pvActuels, pvMax int) int {
	if pvMax <= 0 {
		pvMax = 1
	}
	if pvActuels < 0 {
		return 0
	}
	if pvActuels > pvMax {
		return pvMax
	}
	return pvActuels
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
