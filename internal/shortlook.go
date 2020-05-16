package internal

func shortlook(s string, n int) string {
	words := 0
	index := 0
	for _, v := range s {
		if v == ' ' {
			words++
		}
		if words == n {
			break
		}
		index++
	}
	if words < n {
		return s
	} else {
		return s[:index+1] + "..."
	}
}
