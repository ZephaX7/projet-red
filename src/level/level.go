package statspersonnage

import (
	"fmt"
	"math"
	"strings"
)

// XPCurve calcule l'XP requise pour passer du niveau L à L+1 (L>=1).
type XPCurve func(level int) int

// DefaultXPCurve : croissance exponentielle douce.
// L=1 -> 100, L=2 -> 125, L=3 -> 156, etc.
func DefaultXPCurve(level int) int {
	if level < 1 {
		level = 1
	}
	base := 100.0
	growth := 1.25
	return int(math.Round(base * math.Pow(growth, float64(level-1))))
}

type Progression struct {
	Level    int
	XP       int // XP actuelle dans le niveau courant
	MaxLevel int
	curve    XPCurve
}

func NouvelleProgression(level, maxLevel int, curve XPCurve) *Progression {
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

func (p *Progression) XPToNext() int {
	if p.Level >= p.MaxLevel {
		return 0
	}
	return p.curve(p.Level)
}

// Noms demandés : "Expérience actuelle" / "Expérience max"
func (p *Progression) ExpActuelle() int    { return p.XP }
func (p *Progression) ExpMaxActuelle() int { return p.XPToNext() }
func (p *Progression) NiveauActuel() int   { return p.Level }
func (p *Progression) NiveauMax() int      { return p.MaxLevel }
func (p *Progression) EstAuCap() bool      { return p.Level >= p.MaxLevel }
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

// GainXP ajoute amount, gère les multi level-up, retourne (niveaux gagnés, cap atteint).
// L'excès d'XP est conservé pour le prochain niveau (overflow) tant que cap non atteint.
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
		// au cap, on remet l'XP à 0 pour éviter confusion
		p.XP = 0
		return levels, true
	}
	return levels, false
}

// Barre d'XP ASCII
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
		p.XP, need, p.Percent(), p.Level, p.MaxLevel)
}

// Version avec couleurs ANSI (optionnel)
func (p *Progression) BarColored(width int, fill, empty rune) string {
	const (
		green = "\x1b[32m"
		reset = "\x1b[0m"
	)
	if width < 3 {
		width = 3
	}
	if p.Level >= p.MaxLevel {
		return fmt.Sprintf("[%s%s%s] MAX  Niv %d/%d",
			green, strings.Repeat(string(fill), width), reset, p.Level, p.MaxLevel)
	}
	filled := int(math.Round(p.ProgressFraction() * float64(width)))
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	need := p.XPToNext()
	return fmt.Sprintf("[%s%s%s%s] %d/%d (%d%%)  Niv %d/%d",
		green, strings.Repeat(string(fill), filled), reset,
		strings.Repeat(string(empty), width-filled),
		p.XP, need, p.Percent(), p.Level, p.MaxLevel)
}
