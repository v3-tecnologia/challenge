package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func NormalizeText(s string) string {
	// Normaliza para decompor acentos
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ = transform.String(t, s)

	// Transforma para minúsculas
	s = strings.ToLower(s)

	// Substitui tudo que não é letra, número ou espaço por espaço
	reg, _ := regexp.Compile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "_")

	// Remove underscores duplicados
	s = regexp.MustCompile(`_+`).ReplaceAllString(s, "_")

	// Remove "_" do início e do fim
	s = strings.Trim(s, "_")

	return s
}
