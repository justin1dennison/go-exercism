package dna

import "errors"

// Histogram is a mapping of nucleotides in a strand of DNA to their frequency
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA string

//Counts generate a histogram of valid nucleotides in the given DNA
//Returns error for an invalid DNA strand
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
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
