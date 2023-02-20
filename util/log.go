package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	Trace = iota + 1
	Debug
	Info
	Warn
	Error
	Fatal
)

var Logging bool

func PrintOrLog(s string, level int) {
	switch level {
	case Debug:
		s = "Debug - " + s
	case Info:
		s = "Info - " + s
	case Warn:
		s = "Warning - " + s
	case Error:
		s = "Error - " + s
	case Fatal:
		s = "Fatal - " + s
	default:
	}

	if level < Fatal {
		// only print out the message level less then Fatal
		fmt.Println(s)
	}

	if Logging {
		// write to the log file
		switch level {
		case Warn:
			writeLog(s, "Warn.log")
		case Error:
			writeLog(s, "Error.log")
		case Fatal:
			writeLog(s, "Fatal.log")
		default:
		}
	}
}

func writeLog(msg string, fileName string) {
	filePath := "log/" + fileName
	if _, err := os.Stat(filePath); err == nil {
		// File exist
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		check(err)
		defer file.Close()
		_, err = file.WriteString(msg + "\n")
		check(err)
	} else {
		// File does not exist
		file, err := os.Create(filePath)
		check(err)
		defer file.Close()
		_, err = file.WriteString(msg + "\n")
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CaptureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()

	os.Stdout = writer
	os.Stderr = writer

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	wg.Wait()
	f()

	writer.Close()
	return <-out
}
