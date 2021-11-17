package internal

func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}
	if s[0] >= 'A' && s[0] <= 'Z' {
		return s
	}
	return string(append([]byte{s[0] + 'A' - 'a'}, s[1:]...))
}
