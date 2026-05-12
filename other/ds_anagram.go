package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// Поиск анаграмм (работа с мапами и строками)
// Условие:
// Дан слайс строк words []string. Нужно сгруппировать слова, которые являются анаграммами друг друга (игнорируя регистр).
// Вернуть map[string][]string, где ключ — первое встретившееся слово в группе (в исходном регистре),
// значение — все анаграммы (включая ключ, отсортированные по алфавиту). Пустые строки игнорировать.

func groupAnagrams(words []string) map[string][]string {
	res := make(map[string][]string, len(words))
	for i, s := range words {
		if s == "" {
			continue
		}
		sn := normalize(s)
		for j, v := range words {
			if i != j && sn == normalize(v) {
				res[s] = append(res[s], v)
				words[j] = ""
			}
		}
	}
	for _, v := range res {
		sort.Strings(v)
	}
	return res
}

func normalize(str string) string {
	str = strings.ToLower(str)
	runes := []rune(str)
	slices.Sort(runes)
	return string(runes)
}

func main() {
	words := []string{"кот", "ток", "кто", "листок", "слиток", "столик", "Кот"}
	fmt.Println(groupAnagrams(words))
}
