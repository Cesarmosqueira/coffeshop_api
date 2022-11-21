package util

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func StringToTitle(str string) string {
	return cases.Title(language.English, cases.Compact).String(str)
}

func ValidEmail(str string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(str)
}

func ToSnakeCase(str string) string {
	firstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	newStr := firstCap.ReplaceAllString(str, "${1}_${2}")

	allCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	newStr = allCap.ReplaceAllString(newStr, "${1}_${2}")

	return strings.ToLower(newStr)
}
