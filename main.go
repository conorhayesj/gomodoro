package main

import (
	"fmt"
	"time"
	"flag"
	"github.com/nsf/termbox-go"
)

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func draw(i int, s string) {
	defer termbox.Flush()
	w, h := termbox.Size()
	tbPrint((w/2)-(len(s)/2), (h+i)/2, termbox.ColorRed, termbox.ColorDefault, s)
}

func main() {

	minutePtr := flag.Int("m", 25, "Default 25 mins")
	taskPtr := flag.String("t", "No task", "Task to be completed")

	flag.Parse()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	s := *minutePtr * 60
	for i:=0; i<=s; i++ {
		f := s - i
		q := fmt.Sprintf("%v", f)
		draw(0, q)
		draw(2,*taskPtr)
		time.Sleep(1 * time.Second)
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Close()

}
