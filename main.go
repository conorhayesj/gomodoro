package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	pterm.Info.Println("The time is: ")
	pterm.Print("\n\n")

	area, _ := pterm.DefaultArea.WithCenter().Start()
	for i := 0; i < 10; i++ {
		str, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(time.Now().Format("15:04:05"))).Srender()
		area.Update(str)
		time.Sleep(time.Second)
	}
	area.Stop()
}
