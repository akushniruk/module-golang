package letter

type Map map[rune]int

func Frequency(text string) Map {
	total := Map{}

	for _, count := range text {
		total[count]++
	}

	return total
}

func ConcurrentFrequency(text []string) Map {
	channel := make(chan Map, len(text))

	for _, list := range text {
		go func(word string) {
			channel <- Frequency(word)
		}(list)
	}

	total := Map{}

	for range text {
		for key, value := range <-channel {
			total[key] += value
		}
	}

	return total
}
