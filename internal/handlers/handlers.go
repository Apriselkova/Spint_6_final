package handlers

import (
	// "fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./index.html")
}

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// получаем файл из формы
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Ошибка при получении файла")
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	// закрываем файл
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Println("Ошибка при чтении файла")
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}
	// Преобразуем байтовый массив в строку
	contentStr := string(content)

	// Проверяем, является ли строка пустой или содержит только пробелы
	if len(contentStr) == 0 {
		http.Error(w, "Содержимое файла пустое", http.StatusBadRequest)
		return
	}

	// Получаем переконвертируемую строку из функции service
	convertedString, err := service.Conver(contentStr)
	if err != nil {
		log.Println("Ошибка при конвертации строки")
		http.Error(w, "ошибка при конвертации строки", http.StatusInternalServerError)
		return
	}

	// Создаем локальный файл
	fileName := time.Now().UTC().Format("20060102150405") + filepath.Ext(handler.Filename)
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Println("Ошибка при создании файла")
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(convertedString)
	if err != nil {
		log.Println("Ошибка при записи в файл")
		http.Error(w, "ошибка при записи в файл", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Входные данные: " + contentStr + "\nРезультат конвертации: " + convertedString))
}
