package main

import (
	"bytes"
	"testing"
)

func TestFuture(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	errbuf := bytes.NewBuffer([]byte{})
	opt := &option{ root: "t", }
	runWithGoroutine(buf,errbuf, opt)
	t.Log(buf, errbuf)
}