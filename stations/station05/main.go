package main

func main() {
	sampleFile := "./sample.txt"
	CheckFileExist(sampleFile)

	notExistFile := "./abc.txt"
	CheckFileExist(notExistFile)
}

func CheckFileExist(path string) {}
