package iowordfreq

import "github.com/mnorbi/gosandbox/wordfreq"

type IoWordFreq struct {
	wordfreq.WordFreq
}

func (wf *IoWordFreq) Write(p []byte) (n int, err error) {
	wf.AddWords(string(p))
	return len(p), nil
}
