package basic

const repeatCnt = 5

func Repeat(str string) (result string) {
	for i := 0; i < repeatCnt; i++ {
		result += str
	}

	return
}
