package inventory

import (
	"fmt"
	"projet-red/src/model"
)

type Objet struct {
	Nom      string
	Quantite int
	Type     string // "Objet" ou "Équipement"
}

// Inventaire global accessible depuis d'autres packages
var Inventaire []Objet

// var CapaciteMax = 10   // au lieu de const
var CapaciteMax = 10

// Ajouter un objet à l'inventaire
func AddInventory(obj Objet) {
	for i := range Inventaire {
		if Inventaire[i].Nom == obj.Nom {
			Inventaire[i].Quantite += obj.Quantite
			return
		}
	}
	// Ajouter un nouvel objet si inventaire pas plein
	if len(Inventaire) < CapaciteMax {
		Inventaire = append(Inventaire, obj)
	} else {
		fmt.Println("⚠ Inventaire plein, impossible d'ajouter", obj.Nom)
	}
}

// Retirer un objet de l'inventaire
func RemoveInventory(obj Objet) {
	for i := range Inventaire {
		if Inventaire[i].Nom == obj.Nom {
			if Inventaire[i].Quantite > obj.Quantite {
				Inventaire[i].Quantite -= obj.Quantite
			} else {
				Inventaire = append(Inventaire[:i], Inventaire[i+1:]...)
			}
			return
		}
	}
	fmt.Println("⚠ Objet non trouvé dans l'inventaire :", obj.Nom)
}

// Afficher l'inventaire et permettre de choisir un objet à utiliser ou équiper
func AccessInventory(perso *model.Personnage, inCombat bool, enemy *model.Personnage, utiliserObjet func(string, *model.Personnage, *[]Objet, *model.Personnage, bool)) {
	fmt.Println("=== Inventaire ===")
	for i, item := range Inventaire {
		fmt.Printf("%d. %s (x%d)\n", i+1, item.Nom, item.Quantite)
	}
	fmt.Printf("➡ %d/%d places utilisées\n", len(Inventaire), CapaciteMax)

	fmt.Println("\nQue voulez-vous faire ?")
	fmt.Println("0 - Revenir")
	fmt.Println("1 - Utiliser un objet")
	fmt.Println("2 - Équiper un équipement")
	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 0:
		return
	case 1:
		fmt.Println("Entrez le numéro de l'objet à utiliser :")
		var objIndex int
		fmt.Scan(&objIndex)
		if objIndex < 1 || objIndex > len(Inventaire) {
			fmt.Println("⚠ Choix invalide.")
			return
		}
		obj := Inventaire[objIndex-1]
		utiliserObjet(obj.Nom, perso, &Inventaire, enemy, inCombat)
	case 2:
		fmt.Println("Entrez le numéro de l'équipement à équiper :")
		var eqIndex int
		fmt.Scan(&eqIndex)
		if eqIndex < 1 || eqIndex > len(Inventaire) {
			fmt.Println("⚠ Choix invalide.")
			return
		}
		obj := Inventaire[eqIndex-1]
		if obj.Type != "Équipement" {
			fmt.Println("⚠ Cet objet ne peut pas être équipé.")
			return
		}
		perso.EquipItem(obj.Nom)
	default:
		fmt.Println("⚠ Choix invalide.")
	}
}
