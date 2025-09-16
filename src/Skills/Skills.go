package skills

import (
	"fmt"
)

func Skills() {
	obtainedMagicSkills := 0
	if obtainedMagicSkills == 0 {
		fmt.Println("Vous n'avez pas encore débloqué de compétences magiques.")
	} else {
		fmt.Println("Voici vos compétences :")
		skills := []string{"coup de poing"} // Exemple de compétences
		for i := range skills {
			fmt.Println(skills[i])
		}

		newSkills := []string{} // Liste vide pour l'instant
		if len(newSkills) > 0 {
			fmt.Println("Vous avez débloqué les compétences suivantes :")
			for _, skill := range newSkills {
				fmt.Println(skill)
			}
		}

		response := ""
		if response == "oui" {
			spellName := "boule de feu"
			fmt.Println("Vous avez utilisé un livre de sorts pour le sort " + spellName + ".")
		} else {
			fmt.Println("Vous n'avez pas utilisé de livre de sorts.")
		}

	}

}
