package letter

import "sync"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	m := FreqMap{}
	c := make(chan FreqMap, len(texts))
	wg.Add(len(texts))
	for _, t := range texts {
		go func(results chan FreqMap, t string) {
			results <- Frequency(t)
			wg.Done()
		}(c, t)
	}
	wg.Wait()
	close(c)
	for r := range c {
		for k, v := range r {
			m[k] += v
		}
	}
	return m

}
