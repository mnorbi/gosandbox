package wordfreq

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
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

func NewWordFreq() *WordFreq {
	return &WordFreq{make(map[string]int)}
}

func (wf WordFreq) AddWords(words string) {
	var buf bytes.Buffer
	for i := 0; i < len(words); {
		r, rSize := utf8.DecodeRuneInString(words[i:])
		if unicode.IsLetter(r) {
			buf.WriteRune(unicode.ToLower(r))
		} else {
			if buf.Len() > 0 {
				wf.merge(buf)
			}
			buf.Reset()
		}
		i += rSize
	}
	if buf.Len() > 0 {
		wf.merge(buf)
	}
}

func (wf WordFreq) merge(buf bytes.Buffer) {
	word := buf.String()
	wf.freq[word] = wf.freq[word] + 1
}

func (wf WordFreq) GetWordCount(word string) int {
	return wf.freq[strings.ToLower(word)]
}
