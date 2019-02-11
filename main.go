package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/umaumax/cgrep"
)

var (
	maxWidth int
	tabWidth int
)

func init() {
	flag.IntVar(&maxWidth, "max", 80, "max width")
	flag.IntVar(&tabWidth, "tab", 8, "tab width")
}

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	cnt := 0
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Replace(text, "\t", strings.Repeat(" ", tabWidth), -1)
		ansiText := cgrep.ANSITextParse(text)
		tail := "..."
		truncatedPlaintext, truncated := Truncate(ansiText.Plaintext, maxWidth, tail)
		truncatedTextLen := len([]rune(truncatedPlaintext))
		if truncated {
			truncatedTextLen -= len([]rune(tail))
		}
		truncatedText := ansiText.TrancateString(truncatedTextLen)
		if truncated {
			truncatedText += tail
		}
		if cnt > 0 {
			fmt.Println()
		}
		fmt.Print(truncatedText)
		cnt++
	}
	if err := scanner.Err(); err != nil {
		log.Printf("stdin read err:%s\n", err)
	}
}

// FYI: github.com/mattn/go-runewidth/runewidth.go
func Truncate(s string, w int, tail string) (string, bool) {
	c := runewidth.DefaultCondition
	if c.StringWidth(s) <= w {
		return s, false
	}
	r := []rune(s)
	tw := c.StringWidth(tail)
	w -= tw
	width := 0
	i := 0
	for ; i < len(r); i++ {
		cw := c.RuneWidth(r[i])
		if width+cw > w {
			break
		}
		width += cw
	}
	return string(r[0:i]) + tail, true
}
