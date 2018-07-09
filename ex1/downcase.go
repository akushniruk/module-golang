package downcase

func Downcase(str string) (string, error) {
	strtemp := make([]byte, len(str))

	for count := range strtemp {
		check := str[count]

		if check >= 'A' && check <= 'Z' {
			check += 'a' - 'A'
		}

		strtemp[count] = check
	}

	return string(strtemp), nil
}
