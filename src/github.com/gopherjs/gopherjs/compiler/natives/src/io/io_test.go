// +build js

package io

import (
	"testing"
)

func TestMultiWriter_WriteStringSingleAlloc(t *testing.T) {
	t.Skip()
}

func TestMultiReaderFlatten(t *testing.T) {
	t.Skip()
}

func TestMultiReaderFreesExhaustedReaders(t *testing.T) {
	t.Skip("test relies on runtime.SetFinalizer, which GopherJS does not implement")
}
