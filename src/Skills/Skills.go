package skills

func Skills() {

	obtainedMagicSkills := 0
	if obtainedMagicSkills == 0 {
		println("Vous n'avez pas encore débloqué de compétences magiques.")
	} else {
		println("Voici vos compétences :")
		skills := 0
		for skill := range skills {
			println(skill, "niveau")
			skill = "coup de poing"
		}
	}
	newSkills := 0
	if newSkills > 0 {
		println("Vous avez débloqué les compétences suivantes :")
		for skill := range newSkills {
			println(skill)
		}
	}
	useSpellBook := nil
	if useSpellBook {
		println("Vous utilisez le livre de sorts pour obtenir de nouvelles compétences.")
	}
	println("Voulez-vous utiliser un livre de sorts pour débloquer une nouvelle compétence ? (oui/non)")
	var response string
	Scanln(&response)
	if response == "oui" {
		useSpellBook = true
		spellName := ""
		println("Vous avez utilisé un livre de sorts pour le sort " + spellName + ".")
	} else {
		println("Vous n'avez pas utilisé de livre de sorts.")
	}
}
