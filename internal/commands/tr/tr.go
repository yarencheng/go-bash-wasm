package tr

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Tr struct{}

func New() *Tr {
	return &Tr{}
}

func (t *Tr) Name() string {
	return "tr"
}

func (t *Tr) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("tr", pflag.ContinueOnError)
	deleteFlag := flags.BoolP("delete", "d", false, "delete characters in SET1, do not translate")
	squeezeFlag := flags.BoolP("squeeze-repeats", "s", false, "replace each sequence of a repeated character that is listed in the last specified SET, with a single occurrence of that character")
	complementFlag := flags.BoolP("complement", "c", false, "use the complement of SET1")
	truncateFlag := flags.BoolP("truncate-set1", "t", false, "first truncate SET1 to length of SET2")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "tr: %v\n", err)
		return 1
	}

	remaining := flags.Args()

	data, err := io.ReadAll(env.Stdin)
	if err != nil {
		fmt.Fprintf(env.Stderr, "tr: %v\n", err)
		return 1
	}

	if *deleteFlag {
		if len(remaining) < 1 {
			fmt.Fprintf(env.Stderr, "tr: missing operand after '-d'\n")
			return 1
		}
		set1 := t.expandSet(remaining[0])
		if *complementFlag {
			set1 = t.complement(set1)
		}
		result := t.deleteChars(string(data), set1)
		if *squeezeFlag && len(remaining) > 1 {
			set2 := t.expandSet(remaining[1])
			result = t.squeeze(result, set2)
		}
		fmt.Fprint(env.Stdout, result)
		return 0
	}

	if len(remaining) < 2 {
		fmt.Fprintf(env.Stderr, "tr: missing operand\n")
		return 1
	}

	set1 := t.expandSet(remaining[0])
	set2 := t.expandSet(remaining[1])

	if *truncateFlag && len(set1) > len(set2) {
		set1 = set1[:len(set2)]
	}

	if *complementFlag {
		// Complement translation is slightly more complex.
		// GNU tr -c SET1 SET2 translates characters not in SET1 to the last char of SET2.
		result := t.translateComplement(string(data), set1, set2)
		if *squeezeFlag {
			result = t.squeeze(result, set2)
		}
		fmt.Fprint(env.Stdout, result)
		return 0
	}

	result := t.translate(string(data), set1, set2)
	
	if *squeezeFlag {
		result = t.squeeze(result, set2)
	}

	fmt.Fprint(env.Stdout, result)
	return 0
}


func (t *Tr) expandSet(s string) string {
	// Very basic expansion of a-z, A-Z, 0-9
	if s == "a-z" {
		return "abcdefghijklmnopqrstuvwxyz"
	}
	if s == "A-Z" {
		return "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if s == "0-9" {
		return "0123456789"
	}
	return s
}

func (t *Tr) deleteChars(input, set string) string {
	var sb strings.Builder
	setMap := make(map[rune]bool)
	for _, r := range set {
		setMap[r] = true
	}

	for _, r := range input {
		if !setMap[r] {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func (t *Tr) translate(input, set1, set2 string) string {
	if len(set2) == 0 {
		return input
	}

	transMap := make(map[rune]rune)
	runes1 := []rune(set1)
	runes2 := []rune(set2)

	for i, r1 := range runes1 {
		r2 := runes2[len(runes2)-1]
		if i < len(runes2) {
			r2 = runes2[i]
		}
		transMap[r1] = r2
	}

	var sb strings.Builder
	for _, r := range input {
		if tr, ok := transMap[r]; ok {
			sb.WriteRune(tr)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func (t *Tr) squeeze(input, set string) string {
	if len(input) == 0 {
		return input
	}

	setMap := make(map[rune]bool)
	for _, r := range set {
		setMap[r] = true
	}

	var sb strings.Builder
	runes := []rune(input)
	sb.WriteRune(runes[0])
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] && setMap[runes[i]] {
			continue
		}
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

func (t *Tr) complement(set string) string {
	setMap := make(map[rune]bool)
	for _, r := range set {
		setMap[r] = true
	}
	var sb strings.Builder
	for i := 0; i < 256; i++ {
		r := rune(i)
		if !setMap[r] {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func (t *Tr) translateComplement(input, set1, set2 string) string {
	if len(set2) == 0 {
		return input
	}
	set1Map := make(map[rune]bool)
	for _, r := range set1 {
		set1Map[r] = true
	}
	lastChar := []rune(set2)[len([]rune(set2))-1]
	var sb strings.Builder
	for _, r := range input {
		if !set1Map[r] {
			sb.WriteRune(lastChar)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
