package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Conver(input string) (string, error) {
	newInput := strings.TrimSpace(input)

	if isMorse(newInput) {
		return morse.ToText(newInput), nil
	}

	// Если это не код Морзе, конвертируем в код Морзе
	return morse.ToMorse(newInput), nil
}

func isMorse(input string) bool {
	for _, char := range input {
		if char != '.' && char != '-' && char != ' ' {
			return false // Если встретился недопустимый символ, возвращаем false

		}
	}
	return true // Если все символы допустимы, возвращаем true
}
