package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработка запроса к /")
	http.ServeFile(w, r, "../index.html")
}

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработка запроса к форме")
	// получаем файл из формы
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Ошибка при получении файла")
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
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
	log.Printf("Имя файла: %s\n", handler.Filename)
	log.Printf("Содержимое файла:\n%s\n", content)

	// Получаем переконвертируемую строку из функции service
	convertedString, err := service.Conver(string(content))
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

	w.Write([]byte("Входные данные: " + string(content) + "\nРезультат конвертации: " + string(convertedString)))
}
