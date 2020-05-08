package gowsay

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"io"
	"log"
	"text/template"

	"github.com/mattn/go-runewidth"
	"github.com/mitchellh/go-wordwrap"
	flag "github.com/ogier/pflag"
)

type Face struct {
	Eyes     string
	Tongue   string
	Thoughts string
	Cowfile  string
}

type Mooptions struct {
	Borg     bool
	Dead     bool
	Greedy   bool
	Paranoid bool
	Stoned   bool
	Tired    bool
	Wired    bool
	Young    bool
	Think    bool
	Columns  int32
	Cowfile  string
}

func newFace(options Mooptions) Face {
	f := Face{
		Eyes:    "oo",
		Tongue:  "  ",
		Cowfile: options.Cowfile,
	}

	if options.Borg {
		f.Eyes = "=="
	}
	if options.Dead {
		f.Eyes = "xx"
		f.Tongue = "U "
	}
	if options.Greedy {
		f.Eyes = "$$"
	}
	if options.Paranoid {
		f.Eyes = "@@"
	}
	if options.Stoned {
		f.Eyes = "**"
		f.Tongue = "U "
	}
	if options.Tired {
		f.Eyes = "--"
	}
	if options.Wired {
		f.Eyes = "OO"
	}
	if options.Young {
		f.Eyes = ".."
	}

	return f
}

func readInput(args []string) []string {
	var tmps []string
	if len(args) == 0 {
		s := bufio.NewScanner(os.Stdin)

		for s.Scan() {
			tmps = append(tmps, s.Text())
		}

		if s.Err() != nil {
			log.Printf("failed reading stdin: %s\n", s.Err().Error())
			os.Exit(1)
		}

		if len(tmps) == 0 {
			fmt.Println("Error: no input from stdin")
			os.Exit(1)
		}
	} else {
		tmps = args
	}

	var msgs []string
	for i := 0; i < len(tmps); i++ {
		expand := strings.Replace(tmps[i], "\t", "        ", -1)

		tmp := wordwrap.WrapString(expand, uint(*columns))
		for _, s := range strings.Split(tmp, "\n") {
			msgs = append(msgs, s)
		}
	}

	return msgs
}

func setPadding(msgs []string, width int) []string {
	var ret []string
	for _, m := range msgs {
		s := m + strings.Repeat(" ", width-runewidth.StringWidth(m))
		ret = append(ret, s)
	}

	return ret
}

func constructBallon(f *face, msgs []string, width int) string {
	var borders []string
	line := len(msgs)

	if *think {
		f.Thoughts = "o"
		borders = []string{"(", ")", "(", ")", "(", ")"}
	} else {
		f.Thoughts = "\\"
		if line == 1 {
			borders = []string{"<", ">"}
		} else {
			borders = []string{"/", "\\", "\\", "/", "|", "|"}
		}
	}

	var lines []string

	topBorder := " " + strings.Repeat("_", width+2)
	bottomBoder := " " + strings.Repeat("-", width+2)

	lines = append(lines, topBorder)
	if line == 1 {
		s := fmt.Sprintf("%s %s %s", borders[0], msgs[0], borders[1])
		lines = append(lines, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], msgs[0], borders[1])
		lines = append(lines, s)
		i := 1
		for ; i < line-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], msgs[i], borders[5])
			lines = append(lines, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], msgs[i], borders[3])
		lines = append(lines, s)
	}

	lines = append(lines, bottomBoder)
	return strings.Join(lines, "\n")
}

func maxWidth(msgs []string) int {
	max := -1
	for _, m := range msgs {
		l := runewidth.StringWidth(m)
		if l > max {
			max = l
		}
	}

	return max
}

func renderCow(f *face, w io.Writer) {
	t := template.Must(template.New("cow").Parse(cows[f.cowfile]))

	if err := t.Execute(w, f); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func MakeCow(options Mooptions) {
	inputs := readInput(flag.Args())
	width := maxWidth(inputs)
	messages := setPadding(inputs, width)

	f := newFace()
	balloon := constructBallon(f, messages, width)

	fmt.Println(balloon)
	renderCow(f, os.Stdout)
}
