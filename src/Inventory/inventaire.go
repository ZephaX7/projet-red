package inventory

import "fmt"

func accessInventory() {
	inventory := []string{}

	fmt.Println("Voici votre inventaire :")
	if len(inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
	} else {
		for i, item := range inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}
	addInventory()
}

func addInventory() {
	var newItem string
	