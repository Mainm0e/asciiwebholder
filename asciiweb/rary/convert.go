package rary

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Output(s, t string) string {
	String := s
	Theme := t
	FileName := "hello.txt"
	var NewString string

	splitString := strings.Split(String, "\\n")

	for i := 0; i < len(splitString); i++ {
		if len(splitString) > 0 {
			NewString = NewString + makeArt(splitString[i], Theme)
		} else {
			NewString = NewString + makeArt(splitString[i], Theme)
		}
	}

	output := []byte(NewString)
	MakeFile := os.WriteFile(FileName, output, 0644)
	Check(MakeFile)
	return NewString
}
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func makeArt(s, t string) string {
	Theme := t
	String := s
	var AsciiArt string
	var intLetter int
	file, err := os.Open(Theme)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	var l1, l2, l3, l4, l5, l6, l7, l8 string

	for s := 0; s < len(String); s++ {
		intLetter = ((int(String[s]) - 32) * 9) + 1

		for numb, eachline := range txtlines {
			if numb == intLetter {
				l1 = l1 + eachline
			}
			if numb == intLetter+1 {
				l2 = l2 + eachline
			}
			if numb == intLetter+2 {
				l3 = l3 + eachline
			}
			if numb == intLetter+3 {
				l4 = l4 + eachline
			}
			if numb == intLetter+4 {
				l5 = l5 + eachline
			}
			if numb == intLetter+5 {
				l6 = l6 + eachline
			}
			if numb == intLetter+6 {
				l7 = l7 + eachline
			}
			if numb == intLetter+7 {
				l8 = l8 + eachline
			}
		}
	}

	AsciiArt = l1 + "\n" + l2 + "\n" + l3 + "\n" + l4 + "\n" + l5 + "\n" + l6 + "\n" + l7 + "\n" + l8
	return AsciiArt
}
