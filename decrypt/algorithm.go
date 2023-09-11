package decrypt

func Nimbus(str string) string {
	decyptStr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		decyptStr += character
	}
	return decyptStr
}
