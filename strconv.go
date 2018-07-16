package mantis

//全角转换半角s
func DBCtoSBC(s string) string {
	res := ""
	for _, i := range s {
		insideCode := i
		if insideCode == 12288 {
			insideCode = 32
		} else {
			insideCode -= 65248
		}
		if insideCode < 32 || insideCode > 126 {
			res += string(i)
		} else {
			res += string(insideCode)
		}
	}
	return res
}
