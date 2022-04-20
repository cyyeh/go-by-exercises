package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var freqMap FreqMap
	freqMapChannel := make(chan FreqMap, len(l))
	freqMapResult := FreqMap{}

	for _, words := range l {
		go func(words string) {
			freqMapChannel <- Frequency(words)
		}(words)
	}

	for i := 0; i < len(l); i++ {
		freqMap = <-freqMapChannel
		for key, value := range freqMap {
			_, ok := freqMapResult[key]
			if ok {
				freqMapResult[key] += value
			} else {
				freqMapResult[key] = value
			}
		}
	}

	return freqMapResult
}
