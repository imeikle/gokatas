// Package logger uses channels to implement non-blocking logging. Adapted from
// https://youtu.be/zDCKZn4-dck.
//
// Level: advanced
// Topics: design, buffered channels, os/signal
package logger

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	ch chan string // logs
	wg sync.WaitGroup
}

func New(w io.Writer, cap int) *Logger {
	// New is sometimes called a factory function. It's useful when you need
	// to initialize one or more fields of a type.

	l := Logger{
		ch: make(chan string, cap),
	}

	l.wg.Add(1)
	go func() {
		for v := range l.ch {
			fmt.Fprintln(w, v)
		}
		l.wg.Done()
	}()

	return &l
}

func (l *Logger) Stop() {
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Println(s string) {
	select {
	case l.ch <- s:
	default:
		fmt.Println("WARN: dropping logs")
	}
}
