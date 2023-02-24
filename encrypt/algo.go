package encrypt

func Enc(s string) string {

	n := len(s)
	str := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		str = append(str, s[i]+3) //the encryption algorithm simply adds 3 to the ascii code
	}
	return string(str)
}
