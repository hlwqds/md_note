package main

import (
	"bytes"
	"fmt"
)

const wordBits = 64

//IntSet is a set containing non negative integers
//zero value represents an empty set
//we cant visit its member directly
type IntSet struct {
	words  []uint64
	number int //not safe
}

//Len return the number of the set's member
func (s *IntSet) Len() int {
	return s.number
}

//Has the return value of "Has" method indicates whether there is a non negative integers
//every member of set can save 64 integers, they are identified by bits
func (s *IntSet) Has(x int) bool {
	word, bit := x/wordBits, x%wordBits
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//Add Add will add x into set, when the space is not enough, we should extend the set
//so that it has just enough space
func (s *IntSet) Add(x int) {
	word, bit := x/wordBits, x%wordBits
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
	s.number++
}

//Remove the return value: true indicates that this set have member x
func (s *IntSet) Remove(x int) bool {
	word, bit := x/wordBits, x%wordBits
	if word > len(s.words) {
		return false
	} else if s.words[word]&(1<<bit) != 0 {
		s.words[word] &= ^(1 << bit)
		return true
	} else {
		return false
	}
}

//Clear	clear all the bits of s
func (s *IntSet) Clear() {
	var zero uint64
	for i, _ := range s.words {
		s.words[i] &= zero
	}
	s.number = 0
}

//UnionWith Merge s and t and save the result in s
func (s *IntSet) UnionWith(t *IntSet) {
	var slen int
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}

		//
		for j := 0; j < wordBits; j++ {
			if s.words[i]&(1<<j) != 0 {
				slen++
			}
		}
	}

	s.number = slen
}

//reach the intersection of s and t, save result in s
func (s *IntSet) Intersection(t *IntSet) {
	var slen int
	var wordTmp []uint64
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
			wordTmp = append(wordTmp, s.words[i])
		} else {
			break
		}

		//
		for j := 0; j < wordBits; j++ {
			if s.words[i]&(1<<j) != 0 {
				slen++
			}
		}
	}
	s.words = wordTmp
	s.number = slen
}

//return the copy of s
func (s *IntSet) Copy() *IntSet {
	//take care of the pointer in composite type
	wordCpy := []uint64{}
	copy(wordCpy, s.words)
	c := IntSet{
		words:  wordCpy,
		number: s.number,
	}
	return &c
}

//String return string like {1,2,3}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteRune('{')
	for i, sword := range s.words {
		for j := 0; j < wordBits; j++ {
			if sword&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteRune(' ')
				}
				fmt.Fprintf(&buf, "%d", wordBits*i+j)
			}
		}
	}

	buf.WriteRune('}')
	return buf.String()
}

func main() {
	var set1, set2 IntSet
	set1.Add(2)
	fmt.Println(set1.Len())

	set1.Add(122)
	fmt.Println(set1.Len())

	set2.Add(0)
	set2.Add(999)
	fmt.Println(set1.String())
	fmt.Println(set2.String())
	set1.UnionWith(&set2)
	fmt.Println(set1.Len())

	fmt.Println(set1.String())
	fmt.Println(set1.Has(2))
	fmt.Println(set1.Has(10))

	fmt.Println(set1.Remove(9999))
	fmt.Println(set1.Remove(2))
	fmt.Println(set1.Remove(13))
	fmt.Println(set1.String())

	set1.Clear()
	fmt.Println(set1.String())

	set3 := set1.Copy()
	set3.Add(10)
	fmt.Println(set1.String())
	fmt.Println(set3.String())
	fmt.Println(set1.Len())
	fmt.Println(set3.Len())
	set1.Add(10)
	set1.Intersection(set3)

	fmt.Println(set1.String())
	fmt.Println(set1.Len())
	set1.Intersection(&set2)
	fmt.Println(set1.Len())
}
