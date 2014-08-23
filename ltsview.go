package ltsview

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/ymotongpoo/goltsv"
)

type LTSView struct {
	Reader io.Reader
	Writer io.Writer

	Keys  map[string]struct{}
	Ikeys map[string]struct{}
}

func (v *LTSView) Start() {
	reader := goltsv.NewReader(v.Reader)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		PrintSep(v.Writer)
		sortedKeys := SortKeys(record)
		for _, k := range sortedKeys {
			if _, ignore := v.Ikeys[k]; len(v.Ikeys) > 0 && ignore {
				continue
			}
			if _, show := v.Keys[k]; len(v.Keys) > 0 && !show {
				continue
			}
			PrintLine(v.Writer, k, record[k])
		}
	}
}

func SortKeys(record map[string]string) []string {
	keys := []string{}
	for k := range record {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func PrintLine(w io.Writer, k, v string) {
	fmt.Fprintln(w, k+":", v)
}

func PrintSep(w io.Writer) {
	fmt.Fprintln(w, "---")
}

type ColorableWriter struct {
	io.Writer
}

func (w *ColorableWriter) Write(p []byte) (int, error) {
	f := color.New(color.FgMagenta).SprintFunc()
	for i := range p {
		str := string(p[i])
		if p[i] == ':' {
			fmt.Fprint(w.Writer, str)
			f = color.New(color.FgGreen).SprintFunc()
			continue
		}
		fmt.Fprint(w.Writer, f(str))
	}
	return len(p), nil
}

func ParseKeysByFlag(flag string) map[string]struct{} {
	keys := map[string]struct{}{}
	for _, k := range strings.Split(flag, ",") {
		if k == "" {
			continue
		}
		keys[k] = struct{}{}
	}
	return keys
}
