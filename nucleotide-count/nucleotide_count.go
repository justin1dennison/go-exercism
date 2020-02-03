package dna

import "errors"

type Histogram map[rune]int

func NewHistogram() Histogram {
	h := Histogram(make(map[rune]int))
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0
	return h
}

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := NewHistogram()
	for _, r := range d {
		switch r {
		case 'A', 'C', 'G', 'T':
			h[rune(r)]++
		default:
			return nil, errors.New("Invalid Strand")
		}
	}
	return h, nil
}
