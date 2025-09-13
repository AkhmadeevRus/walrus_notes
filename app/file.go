package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func (notes *App) ReadNotesFromFile() error {
	filePath := "note.json"
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("create file")
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		os.WriteFile(filePath, []byte("[]"), 0644)
		return nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&notes.Notes)
	if err != nil {
		return err
	}
	return nil
}

func (notes *App) WriteNotesToFile() error {
	filePath := "note.json"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.New("ошибка при записи в файл, ты потерял все данные!!!:(")
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(notes.Notes)
	if err != nil {
		return errors.New("ошибка при записи в файл, ты потерял все данные!!!:(")
	}
	return nil
}
