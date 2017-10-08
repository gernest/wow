package wow

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/gernest/wow/spinner"
)

const erase = "\033[2K\r"

// Wow writes beautiful spinners on the terminal.
type Wow struct {
	txt     string
	s       spinner.Spinner
	out     io.Writer
	running bool
	done    func()
	mu      sync.RWMutex
}

// New creates a new wow instance ready to start spinning.
func New(o io.Writer, s spinner.Spinner, text string) *Wow {
	return &Wow{out: o, s: s, txt: text}
}

// Start starts the spinner. The frames are written based on the spinner
// interval.
func (w *Wow) Start() {
	if !w.running {
		ctx, done := context.WithCancel(context.Background())
		t := time.NewTicker(time.Duration(w.s.Interval) * time.Millisecond)
		w.done = done
		w.running = true
		go func() {
			at := 0
			for {
				select {
				case <-ctx.Done():
					break
				case <-t.C:
					if at >= len(w.s.Frames) {
						at = 0
					}
					txt := erase + w.s.Frames[at] + w.txt
					fmt.Fprint(w.out, txt)
					at++
				}
			}
		}()
	}
}

// Stop stops the spinner
func (w *Wow) Stop() {
	if w.done != nil {
		w.done()
	}
}

// Spinner sets s to the current spinner
func (w *Wow) Spinner(s spinner.Spinner) *Wow {
	w.mu.Lock()
	w.s = s
	w.mu.Unlock()
	return w
}
func (w *Wow) Text(txt string) *Wow {
	w.mu.Lock()
	w.txt = txt
	w.mu.Unlock()
	return w
}

// Persist writes the last character of the currect spinner frames together with
// the text on stdout.
//
// A new line is added at the end to ensure the text stay that way.
func (w *Wow) Persist() {
	w.Stop()
	at := len(w.s.Frames) - 1
	txt := erase + w.s.Frames[at] + w.txt + "\n"
	fmt.Fprint(w.out, txt)
}

// PersisWith writes the last frame of s together with text with a new line
// added to make it stick.
func (w *Wow) PersisWith(s spinner.Spinner, text string) {
	w.Stop()
	var a string
	if len(s.Frames) > 0 {
		a = s.Frames[len(s.Frames)-1]
	}
	txt := erase + a + text + "\n"
	fmt.Fprint(w.out, txt)
}
