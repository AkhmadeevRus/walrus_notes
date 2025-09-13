package app

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type menuFunction struct {
	Title    string
	Function func() error
}

func generateMenu(commands map[string]menuFunction) string {
	menu := "Меню:\n"
	var idx string
	for i := 0; i < len(commands); i++ {
		idx = strconv.Itoa(i)
		menu += fmt.Sprintf("%s - %s\n", idx, commands[idx].Title)
	}
	menu += "Введите номер команды: "
	return menu
}

func generateListMenu(notes []Note) {
	for i, note := range notes {
		if utf8.RuneCountInString(note.Title) > 19 {
			titleRune := []rune(note.Title)
			fmt.Printf("%d - %-20s - %s\n", i+1, string(titleRune[:17])+"...", note.CreatedAt.Format("01.02.2006"))
		} else {
			fmt.Printf("%d - %-20s - %s\n", i+1, note.Title, note.CreatedAt.Format("01.02.2006"))
		}
	}
}
