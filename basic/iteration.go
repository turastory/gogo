package basic

func Repeat(str string, repeatCnt int) (result string) {
	for i := 0; i < repeatCnt; i++ {
		result += str
	}

	return
}
