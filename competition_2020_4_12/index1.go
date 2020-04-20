package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := []string{"dyk", "rgdykdu", "e", "fjrgdykduf", "nbtuep", "mxrgdykdul", "hxwang", "omxrgdykdula", "zhfkwzfbjzvpfeh", "bfjrgdykdufl", "cupitoye", "rbzhfkwzfbjzvpfehrc", "empelypoealbb", "nvnbtuepar", "bzmqfbzssvpj", "qfcupitoye", "mvvvekjw", "qfcupitoyenb", "evwttioleugq", "mdykb", "vfdhehqjenr", "qfcupitoyenba", "rknmkxrbreftew", "dddyk", "uykehqi", "ldmdykbx", "aey", "wmbzmqfbzssvpj", "qnas", "rknmkxrbreftewu", "ifyhwhpktbiixb", "vaeyxx", "ltnfyituk", "juykehqiso", "kj", "ulmxrgdykduldq", "kkpnhivvrlptukt", "mjuykehqisoma"}
	s := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(stringMatching(s))
}

func stringMatching(words []string) []string {
	subString := []string{}
	if len(words) == 0 {
		return subString
	}

	recordWord := []string{} //记录遍历的字符串

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			value1 := ""
			value2 := ""
			if len(words[i]) < len(words[j]) {
				value1, value2 = words[j], words[i]
			} else {
				value1, value2 = words[i], words[j]
			}
			if i == j || record(recordWord, value1+"*"+value2) {
				continue
			}

			recordWord = append(recordWord, value1+"*"+value2)

			if strings.Contains(value1, value2) {

				if !record(subString, value2) {
					subString = append(subString, value2)
				}
			}
			continue

		}
	}
	return subString
}

func record(words []string, word string) bool {
	flag := false
	for _, v := range words {
		if v == word {
			return true
		}

	}
	return flag
}
