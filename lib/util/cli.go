package util

func EnforceFlag(flag interface{}, compare interface{}, message string) {
	if flag == compare {
		Fatalf(message)
	}
}
