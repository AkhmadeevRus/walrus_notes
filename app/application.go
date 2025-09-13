package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type App struct {
	Notes []Note
}

func NewNote(title string, description string) *Note {
	return &Note{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

func RunApp() {
	var notes = App{Notes: []Note{}}
	err := notes.ReadNotesFromFile()
	if err != nil {
		fmt.Println(err)
	}
	var funcTable = map[string]menuFunction{
		"0": menuFunction{Title: "Выход", Function: notes.StopApp},
		"1": menuFunction{Title: "Тестовая функция с ошибкой", Function: TestFuncWithError},
		"2": menuFunction{Title: "Тестовая успешная функция", Function: AccessFunc},
		"3": menuFunction{Title: "Добавить новую запись", Function: notes.AddNote},
		"4": menuFunction{Title: "вывести все записи", Function: notes.ListNotes},
		"5": menuFunction{Title: "изменить выбранную заметку", Function: notes.UpdateNote},
	}
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

func (notes *App) StopApp() error {
	err := notes.WriteNotesToFile()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\033[33mДо свидания!\033[0m")
	os.Exit(0)
	return nil
}

func (notes *App) AddNote() error {
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
	notes.Notes = append(notes.Notes, *note)
	return nil
}

func (notes *App) ListNotes() error {
	var command int
	for {
		generateListMenu(notes.Notes)
		fmt.Println("0 - закончить просмотр записей")
		fmt.Print("вы можете увидеть подробное описание заметки, введя ее номер:")
		_, err := fmt.Scan(&command)
		if err != nil || command > len(notes.Notes) {
			fmt.Println("\n\033[31mКоманда не найдена\033[0m")
			continue
		}
		if command == 0 {
			break
		}
		fmt.Printf("%d - %s - %s\n%-100s\n", command, notes.Notes[command-1].Title, notes.Notes[command-1].CreatedAt.Format("01.02.2006"), notes.Notes[command-1].Description)
	}
	return nil
}

func (notes *App) UpdateNote() error {
	scan := bufio.NewScanner(os.Stdin)
	generateListMenu(notes.Notes)
	fmt.Print("введите номер заметки, которую хотите изменить: ")
	for scan.Scan() {
		break
	}
	scan.Scan()
	command := scan.Text()
	idx, err := strconv.Atoi(command)
	if err != nil || idx > len(notes.Notes) || idx < 1 {
		return errors.New("\033[31mКоманда не найдена\033[0m")
	}
	fmt.Println("введите новый title для заметки: ")
	scan.Scan()
	title := scan.Text()
	if title == "" {
		fmt.Println("Текст заголовка не изменен!")
	} else {
		notes.Notes[idx-1].Title = title
		fmt.Println("Текст заголовка изменен!")
	}
	fmt.Println("введите новой description для заметки: ")
	scan.Scan()
	description := scan.Text()
	if description == "" {
		fmt.Println("Текст заметки не изменен!")
	} else {
		notes.Notes[idx-1].Description = description
		fmt.Println("Текст заметки изменен!")
	}
	fmt.Println("введите новою дату для заметки (в формате 01.01.2000): ")
	scan.Scan()
	date := scan.Text()
	if date != "" {
		ParseDate, err := time.Parse("01.02.2006", date)
		if err != nil {
			fmt.Println("неверный формат даты!")
		} else {
			notes.Notes[idx-1].CreatedAt = ParseDate
			fmt.Println("дата изменена")
		}
	} else {
		fmt.Println("дата не изменена")
	}
	return nil
}
