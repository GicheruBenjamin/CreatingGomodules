package utils

import(
    "errors"
	"strings"
)

func ReverseWords(s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty string")
	}
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " "), nil
}


func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}