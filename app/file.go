package app

import (
	"encoding/json"
	"fmt"
	"os"
)

func getNotesFilePath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("ошибка при взятии пути дериктории: %s", err)
	}
	path += "\\note.json"
	return path
}

func ReadNotesFromFile() ([]Note, error) {
	filePath := getNotesFilePath()
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("create file")
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("ошибка при сoздании файла: %s\n", err)
			return nil, err
		}
		defer file.Close()
		os.WriteFile(filePath, []byte("[]"), 0644)
		return []Note{}, nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("ошбка при открытии файла(read): %s", err)
		return nil, err
	}
	defer file.Close()
	notes := []Note{}
	err = json.NewDecoder(file).Decode(&notes)
	if err != nil {
		fmt.Printf("ошибка при чтении из json файла: %s", err)
		return nil, err
	}
	return notes, nil
}

func WriteNotesToFile(notes []Note) error {
	filePath := getNotesFilePath()
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("ошибка при открытии файла(write): %s", err)
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(notes)
	if err != nil {
		fmt.Printf("ошибка при записи в файл: %s", err)
		return err
	}
	return nil
}
