package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"path"
	"runtime"
)

func CodeFilename(skip int) string {
	_, filename, line, ok := runtime.Caller(skip)
	if ok {
		return fmt.Sprintf(`%s:%d`, path.Base(filename), line)
	} else {
		return ""
	}
}

func LogCode(me ...interface{}) {
	var mein = []interface{}{
		CodeFilename(2),
	}
	mein = append(mein, me...)

	log.Println(mein...)
}

func streamToByte(r io.Reader) []byte {
	var buf bytes.Buffer
	buf.ReadFrom(r)

	return buf.Bytes()
}
