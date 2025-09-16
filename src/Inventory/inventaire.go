package Inventory

import "fmt"

type Objet struct {
	Nom      string
	Quantite int
	Type     string
}

var Inventaire [10]Objet

// Affiche l'inventaire
func AccessInventory() {
	fmt.Println("Voici votre inventaire :")
	if len(Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
	} else {
		for i, item := range Inventaire {
			fmt.Printf("%d. %s (x%d)\n", i+1, item.Nom, item.Quantite)
		}
	}
}

// Ajoute un objet à l'inventaire
func AddInventory(objet Objet) {
	for i, item := range Inventaire {
		if item.Nom == objet.Nom {
			Inventaire[i].Quantite += objet.Quantite
			fmt.Printf("Vous avez maintenant %d %s.\n", Inventaire[i].Quantite, objet.Nom)
			return
		}
	}
	Inventaire = append(Inventaire, objet)
	fmt.Printf("Vous avez ajouté %d %s à votre inventaire.\n", objet.Quantite, objet.Nom)
}

func RemoveInventory(objet Objet) {
	for i, item := range Inventaire {
		if item.Nom == objet.Nom {
			Inventaire[i].Quantite -= objet.Quantite
			fmt.Printf("Vous avez maintenant %d %s.\n", Inventaire[i].Quantite, objet.Nom)
			return
		}
	}
	Inventaire = append(Inventaire, objet)
	fmt.Printf("Vous avez ajouté %d %s à votre inventaire.\n", objet.Quantite, objet.Nom)
}
