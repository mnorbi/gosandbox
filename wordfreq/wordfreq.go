package wordfreq

import (
	"bytes"
	"strings"
	"unicode"
)

// Make:
// type WordFreq struct {...}
//
// And make the test compile, then pass.
//
// To split a string into pieces you can iterate over:
// buf := bytes.Buffer{}
// for over string
//   if unicode.IsLetter(r)
//      buf.WriteRune(unicode.ToLower(r))
//   else
//      we have one token in buf.String
//      buf.Reset

type WordFreq struct {
	freq map[string]int
}

//func NewWordFreq() *WordFreq {
//	return &WordFreq{make(map[string]int)}
//}

func (wf *WordFreq) AddWords(words string) {
	if wf.freq == nil {
		wf.freq = map[string]int{}
	}
	var buf bytes.Buffer
	for _, r := range words {
		if unicode.IsLetter(r) {
			buf.WriteRune(unicode.ToLower(r))
		} else {
			if buf.Len() > 0 {
				wf.freq[buf.String()]++
			}
			buf.Reset()
		}
	}
	if buf.Len() > 0 {
		wf.freq[buf.String()]++
	}
}

func (wf *WordFreq) GetWordCount(word string) int {
	return wf.freq[strings.ToLower(word)]
}
