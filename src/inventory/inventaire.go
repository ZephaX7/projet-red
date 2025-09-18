package inventory

import (
	"fmt"
	"projet-red/src/items"
	"projet-red/src/model"
)

type Objet struct {
	Nom      string
	Quantite int
	Type     string
}

var Inventaire []Objet

const CapaciteMax = 10

// Affiche l'inventaire
func AccessInventory() {
	fmt.Println("Voici votre inventaire 😊")
	if len(Inventaire) == 0 {
		fmt.Println("   (vide)")
	} else {
		for i, item := range Inventaire {
			fmt.Printf("%d. %s (x%d)\n", i+1, item.Nom, item.Quantite)
		}
		fmt.Printf("➡ %d/%d places utilisées\n", len(Inventaire), CapaciteMax)
	}
}

// Ajoute un objet à l'inventaire
func AddInventory(objet Objet) {
	// Vérifier la capacité max
	if len(Inventaire) >= CapaciteMax {
		fmt.Println("Inventaire plein ! Impossible d'ajouter :", objet.Nom)
		return
	}

	// Si l'objet existe déjà → on augmente la quantité
	for i, item := range Inventaire {
		if item.Nom == objet.Nom {
			Inventaire[i].Quantite += objet.Quantite
			fmt.Printf("Vous avez maintenant %d %s.\n", Inventaire[i].Quantite, objet.Nom)
			return
		}
	}

	// Sinon, ajouter un nouvel objet
	Inventaire = append(Inventaire, objet)
	fmt.Printf("Vous avez ajouté %d %s à votre inventaire.\n", objet.Quantite, objet.Nom)
}

// Supprime un objet de l'inventaire
func RemoveInventory(objet Objet) {
	for i, item := range Inventaire {
		if item.Nom == objet.Nom {
			Inventaire[i].Quantite -= objet.Quantite
			if Inventaire[i].Quantite <= 0 {
				Inventaire = append(Inventaire[:i], Inventaire[i+1:]...)
				fmt.Printf("%s retiré de l'inventaire.\n", objet.Nom)
			} else {
				fmt.Printf("Vous avez maintenant %d %s.\n", Inventaire[i].Quantite, objet.Nom)
			}
			return
		}
	}
	fmt.Println("⚠ L'objet", objet.Nom, "n'est pas dans l'inventaire.")
}
func UtiliserObjet(nom string, perso *model.Personnage) {
	for _, item := range Inventaire {
		if item.Nom == nom {
			switch item.Nom {
			case "Potion de soin":
				items.TakePot(perso)
			case "Potion de poison":
				items.PoisonPot(perso, "ennemi")
			default:
				fmt.Println("Objet inconnu :", item.Nom)
				return
			}
			RemoveInventory(Objet{Nom: item.Nom, Quantite: 1})
			fmt.Println("Vous avez utilisé :", item.Nom)
			return
		}
	}
	fmt.Println("⚠ Vous n'avez pas cet objet dans l'inventaire :", nom)
}

// Vérifie si l'objet est présent dans l'inventaire
func HasItem(nom string) bool {
	for _, item := range Inventaire {
		if item.Nom == nom && item.Quantite > 0 {
			return true
		}
	}
	return false
}

func HasMagicBook(perso *model.Personnage) bool {
	for _, item := range Inventaire {
		if item.Nom == "Livre de Sort : Boule de feu" {
			return true
		}
	}
	return false
}
