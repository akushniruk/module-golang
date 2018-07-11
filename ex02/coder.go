package cipher

import (
	"regexp"
	s "strings"
)

type Shift struct {
	num int
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(num int) Cipher {
	if (num > 0 && num < 26) || (num < 0 && num > -26) {
		var number Cipher = &Shift{num}
		return number
	} else {
		return nil
	}
}

func convert(str string) string {
	var remove = regexp.MustCompile(`[[:punct:]]|[[:space:]]|[[:digit:]]`)
	strNew := remove.ReplaceAllString(str, "")
	strNew = s.ToLower(strNew)
	return strNew
}

func checkAlphabet(check byte, str string) byte {
	if check < 'a' {
		return ('z' - ('a' - check)) + 1
	} else if check > 'z' {
		return ('a' - ('z' - check)) - 1
	} else {
		return check
	}
}

func (cipher Shift) Encode(str string) string {
	strNew := convert(str)
	strtemp := make([]byte, len(strNew))

	for count := range strtemp {
		check := strNew[count] + byte(cipher.num)
		strtemp[count] = checkAlphabet(check, strNew)
	}
	return string(strtemp)
}

func (cipher Shift) Decode(str string) string {
	strtemp := make([]byte, len(str))

	for count := range strtemp {
		check := str[count] - byte(cipher.num)
		strtemp[count] = checkAlphabet(check, str)
	}

	return string(strtemp)
}
