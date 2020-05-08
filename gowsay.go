package gowsay

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/mattn/go-runewidth"
	"github.com/mitchellh/go-wordwrap"
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
	cowfile := ""
	if options.Cowfile == "" {
		cowfile = "apt"
	} else {
		cowfile = options.Cowfile
	}

	f := Face{
		Eyes:    "oo",
		Tongue:  "  ",
		Cowfile: cowfile,
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

func setPadding(msgs []string, width int) []string {
	var ret []string
	for _, m := range msgs {
		s := m + strings.Repeat(" ", width-runewidth.StringWidth(m))
		ret = append(ret, s)
	}

	return ret
}

func constructBallon(f Face, msgs []string, width int, think bool) string {
	var borders []string
	line := len(msgs)

	if think {
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

func renderCow(face Face) (string, error) {
	var output bytes.Buffer
	templateString, exists := cows[face.Cowfile]
	if !exists {
		return "", fmt.Errorf("No such template: %s", face.Cowfile)
	}
	t := template.Must(template.New("cow").Parse(templateString))
	if err := t.Execute(&output, face); err != nil {
		return "", err
	}
	return output.String(), nil
}

func MakeCow(sentence string, options Mooptions) (string, error) {
	inputs := strings.Split(wordwrap.WrapString(sentence, uint(options.Columns)), "\n")
	width := maxWidth(inputs)
	messages := setPadding(inputs, width)

	face := newFace(options)
	balloon := constructBallon(face, messages, width, options.Think)
	cow, err := renderCow(face)
	if err != nil {
		return "", err
	}
	return strings.Join([]string{balloon, cow}, "\n"), nil
}
