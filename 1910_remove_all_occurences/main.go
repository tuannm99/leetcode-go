package main

func removeOccurrences(s string, part string) string {
    for i := 0; i <= len(s)-len(part); i++ {
		if s[i:i+len(part)] == part {
			s = s[0:i] + s[i+len(part):]
			i = -1
		}
	}
    return s
}
