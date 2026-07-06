package panels

import "fmt"

// Leveling System Menu
func NewStatusPanel(level, expPoints, manaPoints, strength, intelligence int) Panel {
	return Panel{
		Kind:  Status,
		Title: "STATUS",
		Lines: []string{
			fmt.Sprintf("LEVEL  %d", level),
			fmt.Sprintf("EXP    %d", expPoints),
			fmt.Sprintf("MP     %d", manaPoints),
			fmt.Sprintf("STR    %d", strength),
			fmt.Sprintf("INT    %d", intelligence),
		},
	}
}
