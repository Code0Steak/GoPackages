package decrypt

func Dec(s string) string {
	n := len(s)
	str := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		str = append(str, s[i]-3)
	}

	return string(str)

}
