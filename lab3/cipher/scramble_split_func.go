package cipher

import "regexp"

// tokenize returns a slice of tokens for the given text string.
// In this case, a token may be a word, a number, a space, or any punctuation
// as represented in the regular expression in this function.
// Regular expressions are a bit out of scope for this course,
// so feel free to learn about them later.
func tokenize(text string) []string {
	re := regexp.MustCompile(`[A-Za-z0-9']+|[':;?().,!\\ ]`)
	return re.FindAllString(text, -1)
}
