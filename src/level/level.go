package xp

import (
	"fmt"
	"math"
	"strings"
)

type XPCurve func(level int) int

func DefaultXPCurve(level int) int {
	if level < 1 {
		level = 1
	}
	base := 100.0  // XP de base pour passer 1 -> 2
	growth := 1.25 // facteur de croissance par niveau
	return int(math.Round(base * math.Pow(growth, float64(level-1))))
}

// Progression gère le niveau, l'XP courante, le cap et la courbe d'XP.
// (POO : méthodes sur *Progression)
type Progression struct {
	Level    int     // Niveau actuel (>=1)
	XP       int     // XP accumulée dans le niveau courant (>=0)
	MaxLevel int     // Niveau maximum atteignable (cap)
	curve    XPCurve // Courbe XP (par niveau)
}

// New crée une progression valide.
func New(level, maxLevel int, curve XPCurve) *Progression {
	if level < 1 {
		level = 1
	}
	if maxLevel < level {
		maxLevel = level
	}
	if curve == nil {
		curve = DefaultXPCurve
	}
	return &Progression{
		Level:    level,
		XP:       0,
		MaxLevel: maxLevel,
		curve:    curve,
	}
}

// ---- Getters pratiques (lisibles côté affichage) ----
func (p *Progression) XPToNext() int {
	if p.Level >= p.MaxLevel {
		return 0
	}
	return p.curve(p.Level)
}
func (p *Progression) ExpActuelle() int    { return p.XP }
func (p *Progression) ExpMaxActuelle() int { return p.XPToNext() }
func (p *Progression) NiveauActuel() int   { return p.Level }
func (p *Progression) NiveauMax() int      { return p.MaxLevel }
func (p *Progression) EstAuCap() bool      { return p.Level >= p.MaxLevel }

// ProgressFraction : progression [0..1] dans le niveau courant (1 si cap).
func (p *Progression) ProgressFraction() float64 {
	need := p.XPToNext()
	if need <= 0 {
		return 1.0
	}
	f := float64(p.XP) / float64(need)
	if f < 0 {
		return 0
	}
	if f > 1 {
		return 1
	}
	return f
}
func (p *Progression) Percent() int {
	return int(math.Round(p.ProgressFraction() * 100))
}

// GainXP ajoute amount, gère les multi level-up, retourne (niveaux_gagnés, cap_atteint).
// L'excès d'XP est conservé pour le prochain niveau tant que le cap n'est pas atteint.
func (p *Progression) GainXP(amount int) (int, bool) {
	if amount <= 0 || p.Level >= p.MaxLevel {
		return 0, p.Level >= p.MaxLevel
	}
	levels := 0
	p.XP += amount

	for p.Level < p.MaxLevel {
		need := p.curve(p.Level)
		if p.XP < need {
			break
		}
		p.XP -= need
		p.Level++
		levels++
	}
	if p.Level >= p.MaxLevel {
		// Au cap, on remet l'XP à 0 (affichage plus clair)
		p.XP = 0
		return levels, true
	}
	return levels, false
}

// Bar ascii ici
func (p *Progression) Bar(width int, fill, empty rune) string {
	if width < 3 {
		width = 3
	}
	if p.Level >= p.MaxLevel {
		return fmt.Sprintf("[%s] MAX  Niv %d/%d",
			strings.Repeat(string(fill), width), p.Level, p.MaxLevel)
	}
	filled := int(math.Round(p.ProgressFraction() * float64(width)))
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	need := p.XPToNext()
	return fmt.Sprintf("[%s%s] %d/%d (%d%%)  Niv %d/%d",
		strings.Repeat(string(fill), filled),
		strings.Repeat(string(empty), width-filled),
		p.XP, need, p.Percent(), p.Level, p.MaxLevel,
	)
}
