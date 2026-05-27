package main

import "strings"

func cleanWords(s string) string {
	banned := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}

	words := strings.Split(s, " ")
	for i, w := range words {
		if _, found := banned[strings.ToLower(w)]; found {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
