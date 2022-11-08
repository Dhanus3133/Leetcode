func isPalindrome(s string) bool {
	s = strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, ""))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}