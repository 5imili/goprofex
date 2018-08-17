package main

func leftpad(s string, length int, char rune) string {
	for len(s) < length {
		s = string(char) + s
	}
	/*
		buf := bytes.Buffer{}
		for i := 0; i < length-len(s); i++ {
			buf.WriteRune(char)
		}
		buf.WriteString(s)
		return buf.String()
	*/
	/*
		if len(s) < length {
			return strings.Repeat(string(char), length-len(s)) + s
		}
	*/
	return s
}
