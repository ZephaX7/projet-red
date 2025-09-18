package forgeron

import (
	"fmt"
	"projet-red/src/inventory"
	"projet-red/src/model"
)

type Material struct {
	Name     string
	Quantity int
}

type ItemCost struct {
	Gold      int
	Materials []Material
}

func Shop(perso *model.Personnage) {
	items := []string{
		"Chapeau de l'aventurier",
		"Tunique de l'aventurier",
		"Bottes de l'aventurier",
	}

	costs := []ItemCost{
		{Gold: 5, Materials: []Material{{"Plume de Corbeau", 1}, {"Cuir de Sanglier", 1}}},
		{Gold: 5, Materials: []Material{{"Fourrure de loup", 2}, {"Peau de Troll", 1}}},
		{Gold: 5, Materials: []Material{{"Fourrure de loup", 1}, {"Cuir de Sanglier", 1}}},
	}

	bought := []string{}

	for {
		fmt.Printf("\nVous avez %d pièces d'or.\n", perso.Gold)
		for i, item := range items {
			fmt.Printf("%d. %s - %d pièces d'or + matériaux : ", i+1, item, costs[i].Gold)
			for _, m := range costs[i].Materials {
				fmt.Printf("%d %s ", m.Quantity, m.Name)
			}
			fmt.Println()
		}
		fmt.Println("Entrez le numéro de l'article à fabriquer (ou 0 pour quitter) :")

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}
		if choice < 1 || choice > len(items) {
			fmt.Println("Choix invalide.")
			continue
		}

		cost := costs[choice-1]

		// Vérifier l'or
		if perso.Gold < cost.Gold {
			fmt.Println("Vous n'avez pas assez d'or.")
			continue
		}

		// Vérifier les matériaux
		canBuy := true
		for _, m := range cost.Materials {
			found := false
			for _, obj := range inventory.Inventaire {
				if obj.Nom == m.Name && obj.Quantite >= m.Quantity {
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("Il vous manque %d %s.\n", m.Quantity, m.Name)
				canBuy = false
			}
		}
		if !canBuy {
			continue
		}

		// Vérifier l'inventaire
		if len(inventory.Inventaire) >= inventory.CapaciteMax {
			fmt.Println("Votre inventaire est plein, vous ne pouvez rien fabriquer de plus.")
			continue
		}

		// Retirer l'or
		perso.Gold -= cost.Gold

		// Retirer les matériaux
		for _, m := range cost.Materials {
			inventory.RemoveInventory(inventory.Objet{
				Nom:      m.Name,
				Quantite: m.Quantity,
			})
		}

		// Ajouter l'objet fabriqué
		inventory.AddInventory(inventory.Objet{
			Nom:      items[choice-1],
			Quantite: 1,
			Type:     "Équipement",
		})

		bought = append(bought, items[choice-1])
		fmt.Println("Vous avez fabriqué", items[choice-1])
	}
}
