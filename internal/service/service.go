package service

import (
	"regexp"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

const pattern = `^[-. /]+$`

// ConvertMessage определяет является ли текст в файле азбукой Морзе или обычным текстом.
// Конвертирует файл либо в Морзе либо в обычный текст в зависимости от содержания.
func ConvertMessage(s string) (string, error) {
	checkMorse, err := regexp.MatchString(pattern, s)
	if err != nil {
		return "", err
	}
	if checkMorse {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}
