package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

func NewNote(title string, description string) *Note {
	return &Note{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now().Format("02.01.2006"),
	}
}

var funcTable = map[string]menuFunction{
	"0": menuFunction{Title: "Выход", Function: StopApp},
	"1": menuFunction{Title: "Тестовая функция с ошибкой", Function: TestFuncWithError},
	"2": menuFunction{Title: "Тестовая успешная функция", Function: AccessFunc},
	"3": menuFunction{Title: "Добавить новую запись", Function: AddNote},
	"4": menuFunction{Title: "вывести все записи", Function: ListNotes},
	"5": menuFunction{Title: "изменить выбранную заметку", Function: UpdateNote},
}

func RunApp() {
	fmt.Println("\033[33mДобро пожаловать \"Записки Ластоногих\"\033[0m")
	var command string
	for {
		fmt.Println("----------------------------")
		fmt.Print(generateMenu(funcTable))
		fmt.Scan(&command)
		fmt.Println("----------------------------")
		targetF, ok := funcTable[command]
		if !ok {
			fmt.Println("\033[31mКоманда не найдена\033[0m")
			continue
		}
		err := targetF.Function()
		if err != nil {
			fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())

			continue
		}
		fmt.Println("\033[32mУспешно!\033[0m")
	}
}

func AccessFunc() error {
	return nil
}

func TestFuncWithError() error {
	return errors.New("тестовая ошибка")
}

func StopApp() error {
	fmt.Println("\033[33mДо свидания!\033[0m")
	os.Exit(0)
	return nil
}

func AddNote() error {
	var title, description string
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		break // просто очищаем буфер
	}
	fmt.Print("дайте название своей заметке: ")
	scan.Scan()
	title = scan.Text()
	fmt.Print("описание: ")
	scan.Scan()
	description = scan.Text()
	note := NewNote(title, description)
	notes, err := ReadNotesFromFile()
	if err != nil {
		fmt.Println(err)
	}
	notes = append(notes, *note)
	err = WriteNotesToFile(notes)
	if err != nil {
		fmt.Print(err)
	}
	return nil
}

func ListNotes() error {
	notes, err := ReadNotesFromFile()
	if err != nil {
		fmt.Printf("Ошибка при чтении из json файла %s", err)
	}
	var command int
	for {
		fmt.Println("вы можете увидеть подробное описание заметки, введя ее номер:")
		generateListMenu(notes)
		fmt.Println("0 - закончить просмотр записей")
		_, err := fmt.Scan(&command)
		if err != nil || command > len(notes) {
			fmt.Println("\033[31mКоманда не найдена\033[0m")
		}
		if command == 0 {
			break
		}
		fmt.Printf("%d - %s - %s\n%-100s\n", command, notes[command-1].Title, notes[command-1].CreatedAt, notes[command-1].Description)
	}
	return nil
}

func UpdateNote() error {
	notes, err := ReadNotesFromFile()
	if err != nil {
		fmt.Printf("Ошибка при чтении из json файла %s", err)
	}
	var command int
	for {
		generateListMenu(notes)
		fmt.Println("0 - закончить просмотр записей")
		fmt.Print("введите номер заметки, которую хотите изменить: ")
		_, err := fmt.Scan(&command)
		if err != nil || command > len(notes) {
			fmt.Println("\033[31mКоманда не найдена\033[0m")
			return nil
		}
		if command == 0 {
			break
		}
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			break
		}
		fmt.Println("введите новый title для заметки: ")
		scan.Scan()
		title := scan.Text()
		if title == "" {
			fmt.Println("Текст заголовка не изменен!")
		} else {
			notes[command-1].Title = title
			fmt.Println("Текст заголовка изменен!")
		}
		fmt.Println("введите новой description для заметки: ")
		scan.Scan()
		description := scan.Text()
		if description == "" {
			fmt.Println("Текст заметки не изменен!")
		} else {
			notes[command-1].Description = description
			fmt.Println("Текст заметки изменен!")
		}
		fmt.Println("введите новою дату для заметки (в формате 01.01.2000): ")
		scan.Scan()
		date := scan.Text()
		if date != "" {
			_, err := time.Parse("01.02.2006", date)
			if err != nil {
				fmt.Println("неверный формат даты!")
			} else {
				notes[command-1].CreatedAt = date
				fmt.Println("дата изменена")
			}
		} else {
			fmt.Println("дата не изменена")
		}
	}
	WriteNotesToFile(notes)
	return nil
}
