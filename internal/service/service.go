package service

import (
	"strings"
	"unicode"

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

// isMorse проверяет, состоит ли строка только из допустимых символов Морзе.
func isMorse(input string) bool {
	// Допустимые символы Морзе: точки и тире
	isMorseSymbol := func(r rune) bool {
		return r == '.' || r == '-' || unicode.IsSpace(r)
	}

	// Проверяем, есть ли недопустимые символы
	if strings.ContainsFunc(input, func(r rune) bool {
		return !isMorseSymbol(r)
	}) {
		return false
	}
	return true
}
