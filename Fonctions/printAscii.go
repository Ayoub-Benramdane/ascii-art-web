package Fonctions

import (
	"strings"
)

func PrintAsciiArt(banner, str string) string {
	var result string
	table := AsciiTable(strings.Split(banner[1:], "\n\n"))
	Lines := strings.Split(str, "\r\n")
	c := 0
	for i, str := range Lines {
		if str == "" {
			if c == 0 && i == len(Lines)-1 {
				break
			}
			result += "\n"
		} else if str != "" {
			c = 1
			result += AddLine(table, str)
		}
	}
	return result
}

func AsciiTable(split_File []string) [][]string {
	var table [][]string
	for i := 0; i < len(split_File); i++ {
		table = append(table, strings.Split(split_File[i], "\n"))
	}
	return table
}

func AddLine(tableau [][]string, str string) string {
	result := "\n"
	for k := 0; k < 8*len(str); k++ {
		if int(str[k%len(str)]) >= 32 && int(str[k%len(str)]) <= 126 {
			result += (tableau[int(str[k%len(str)])-32][k/len(str)%8])
		}
		if (k+1)%len(str) == 0 {
			result += "\n"
		}
	}
	return result
}
