package cipher

type Vigenere struct {
	key string
}

func NewVigenere(key string) Cipher {
	if !invalidKeys(key) {
		return nil
	} else {
		var vigener Cipher = &Vigenere{key}
		return vigener
	}
}

func (cipher Vigenere) Encode(str string) string {
	strNew := convert(str)
	strtemp := ""
	var char string
	for count := range strNew {
		n := NewShift(int(cipher.key[count%len(cipher.key)] - 'a'))
		if n != nil {
			char = n.Encode(string(strNew[count]))
		} else {
			char = string(strNew[count])
		}
		strtemp += char
	}
	return strtemp
}

func (cipher Vigenere) Decode(str string) string {
	strtemp := ""
	var char string
	for count := range str {
		n := NewShift(int(cipher.key[count%len(cipher.key)] - 'a'))
		if n != nil {
			char = n.Decode(string(str[count]))
		} else {
			char = string(str[count])
		}
		strtemp += char
	}
	return strtemp
}

func invalidKeys(str string) bool {
	check := false

	if len(str) < 1 {
		return false
	}

	for count := range str {
		if str[count] == ' ' || (str[count] > 40 && str[count] < 91) {
			return false
		}

		if str[count] != 'a' {
			check = true
		}
	}
	return check
}
