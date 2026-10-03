package service

import (
	"strings"
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Translater(text string) (string, error) {
	alphabed := "АаБбВвГгДдЕеЁёЖжЗзИиЙйКкЛлМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщЪъЫыЬьЭэЮюЯя"

	if strings.ContainsAny(text, alphabed) {
		return morse.ToMorse(text), nil
	} else if strings.ContainsAny(text, "-."){
		return morse.ToText(text), nil
	}
	return "", fmt.Errorf("Text translate error")
}
