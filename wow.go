package wow

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gernest/wow/spinner"
)

type WriteFlusher interface {
	io.Writer
}

func Render(o io.Writer, s spinner.Spinner, txt string) error {
	t := time.NewTicker(time.Millisecond * time.Duration(s.Interval))
	defer t.Stop()
	for _, v := range s.Frames {
		fmt.Fprint(o, erase)
		fmt.Fprint(o, v)
		fmt.Fprint(o, txt)
		<-t.C
	}
	return nil
}

const erase = "\033[2K\r"

type component struct {
	s   spinner.Spinner
	txt string
}

type Wow struct {
	Text    string
	Spinner spinner.Spinner
	Out     io.Writer
	done    func()
}

func New(o io.Writer, s spinner.Spinner, text string) *Wow {
	return &Wow{Out: o, Spinner: s, Text: text}
}

func (w *Wow) Start() {
	ctx, done := context.WithCancel(context.Background())
	t := time.NewTicker(time.Duration(w.Spinner.Interval) * time.Millisecond)
	w.done = done
	go func() {
		at := 0
		for {
			select {
			case <-ctx.Done():
				break
			case <-t.C:
				if at >= len(w.Spinner.Frames) {
					at = 0
				}
				fmt.Fprint(w.Out, erase)
				fmt.Fprint(w.Out, w.Spinner.Frames[at])
				fmt.Fprintf(w.Out, " %s", w.Text)
				at++
			}
		}
	}()
}

func (w *Wow) Stop() {
	w.done()
}

func pad(txt string) string {
	return "  " + txt
}
