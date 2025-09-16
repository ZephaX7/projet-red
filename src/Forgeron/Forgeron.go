package Forgeron

import (
	"fmt"

	"github.com/ZephaX7/projet-red/src/Inventory"
)

type Material struct {
	Name     string
	Quantity int
}

type ItemCost struct {
	Gold      int
	Materials []Material
}

func Shop() {
	items := []string{
		"Chapeau de l'aventurier",
		"Tunique de l'aventurier",
		"Bottes de l'aventurier",
	}

	costs := []ItemCost{
		{Gold: 5, Materials: []Material{{"Tissu", 2}}},
		{Gold: 5, Materials: []Material{{"Tissu", 3}}},
		{Gold: 5, Materials: []Material{{"Cuir", 2}, {"Tissu", 1}}},
	}

	Gold := 100
	bought := []string{}

	for {
		fmt.Printf("\nVous avez %d pièces d'or.\n", Gold)
		for i, item := range items {
			fmt.Printf("%d. %s - %d pièces d'or + matériaux : ", i+1, item, costs[i].Gold)
			for _, m := range costs[i].Materials {
				fmt.Printf("%d %s ", m.Quantity, m.Name)
			}
			fmt.Println()
		}
		fmt.Println("Entrez le numéro de l'article à acheter (ou 0 pour quitter) :")

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
		if Gold < cost.Gold {
			fmt.Println("Vous n'avez pas assez d'or.")
			continue
		}

		// Vérifier les matériaux
		canBuy := true
		for _, m := range cost.Materials {
			found := false
			for _, obj := range Inventory.Inventaire {
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

		// Vérifier l'inventaire pour l'objet final
		if len(Inventory.Inventaire) >= Inventory.CapaciteMax {
			fmt.Println("Votre inventaire est plein, vous ne pouvez rien acheter de plus.")
			continue
		}

		// Retirer l'or
		Gold -= cost.Gold

		// Retirer les matériaux
		for _, m := range cost.Materials {
			Inventory.RemoveInventory(Inventory.Objet{
				Nom:      m.Name,
				Quantite: m.Quantity,
			})
		}

		// Ajouter l'objet fabriqué
		Inventory.AddInventory(Inventory.Objet{
			Nom:      items[choice-1],
			Quantite: 1,
			Type:     "Équipement",
		})

		bought = append(bought, items[choice-1])
		fmt.Println("Vous avez fabriqué", items[choice-1])
	}
}
