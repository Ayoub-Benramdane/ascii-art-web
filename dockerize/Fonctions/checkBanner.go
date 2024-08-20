package Fonctions

import (
	"os"
	"strings"
)

func AsciiArt(Banner, str string) string {
	file, err := os.ReadFile("banner/" + Banner)
	if err != nil || !CheckBanner(file, str) {
		return ""
	}
	fileFinal := strings.ReplaceAll(string(file), "\r", "")
	return PrintAsciiArt(string(fileFinal), str)
}

func CheckBanner(Banner []byte, str string) bool {
	var maxAscii rune
	for _, c := range str {
		if maxAscii == 0 && c >= 32 && c <= 126 {
			maxAscii = c
		} else if c > maxAscii && c >= 32 && c <= 126 {
			maxAscii = c
		}
	}
	Req_Len := int((maxAscii - 31) * 9)
	Len_File := len(strings.Split(string(Banner), "\n"))
	return Req_Len < Len_File
}
