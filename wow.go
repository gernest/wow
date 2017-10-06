package wow

import (
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

const erase = "\033[1K\r"

type component struct {
	s   spinner.Spinner
	txt string
}

type Wow struct {
	c []component
}

func (w *Wow) AddLine(s spinner.Spinner, txt string) {
	w.c = append(w.c, component{s: s, txt: txt})
}

func (w *Wow) RenderTo(o io.Writer) {
	for _, c := range w.c {
		Render(o, c.s, c.txt)
	}
}
