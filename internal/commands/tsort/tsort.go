package tsort

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"sort"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Tsort struct{}

func New() *Tsort {
	return &Tsort{}
}

func (t *Tsort) Name() string {
	return "tsort"
}

func (t *Tsort) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("tsort", pflag.ContinueOnError)
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "tsort: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	var r io.Reader
	if len(remaining) == 0 || remaining[0] == "-" {
		r = env.Stdin
	} else {
		f, err := env.FS.Open(remaining[0])
		if err != nil {
			fmt.Fprintf(env.Stderr, "tsort: %s: %v\n", remaining[0], err)
			return 1
		}
		defer f.Close()
		r = f
	}

	return t.process(env, r)
}

func (t *Tsort) process(env *commands.Environment, r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if len(words)%2 != 0 {
		fmt.Fprintf(env.Stderr, "tsort: input contains an odd number of tokens\n")
		return 1
	}

	adj := make(map[string][]string)
	inDegree := make(map[string]int)
	nodes := make(map[string]bool)

	for i := 0; i < len(words); i += 2 {
		u, v := words[i], words[i+1]
		nodes[u] = true
		nodes[v] = true
		if u != v {
			adj[u] = append(adj[u], v)
			inDegree[v]++
		}
	}

	var queue []string
	for node := range nodes {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	// For deterministic output, sort the initial queue
	sort.Strings(queue)

	var result []string
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)

		for _, v := range adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				// To keep it somewhat stable/deterministic, we could sort the queue after each add,
				// but Kahn's typically adds to the end.
				queue = append(queue, v)
				sort.Strings(queue)
			}
		}
	}

	if len(result) < len(nodes) {
		fmt.Fprintf(env.Stderr, "tsort: cycle in data\n")
		// Continue to print what we can? GNU tsort prints all nodes even if there is a cycle.
	}

	for _, node := range result {
		fmt.Fprintln(env.Stdout, node)
	}

	return 0
}
