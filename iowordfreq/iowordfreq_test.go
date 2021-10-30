package iowordfreq

import (
	"fmt"
	"testing"

	"github.com/mnorbi/gosandbox/wordfreq"
)

type IoWordFreq struct {
	wordfreq.WordFreq
}

func NewIoWordFreq() *IoWordFreq {
	return &IoWordFreq{*wordfreq.NewWordFreq()}
}

func (wf IoWordFreq) Write(p []byte) (n int, err error) {
	wf.AddWords(string(p))
	return len(p), nil
}

func ExampleIoWordFreq() {
	iowf := *NewIoWordFreq()
	iowf.AddWords("hello friend")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again")
	fmt.Println(iowf.GetWordCount("hello"))

	fmt.Fprintln(&iowf, "and hello again again")
	fmt.Println(iowf.GetWordCount("hello"))
	// Output:
	// 1
	// 2
	// 3
}

func TestMapRelocation(t *testing.T) {
	iowf := *NewIoWordFreq()
	fmt.Fprintln(&iowf, "one")
	if iowf.GetWordCount("one") != 1 {
		t.Errorf("one problem")
	}
}
