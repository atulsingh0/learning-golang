package main

import (
	"fmt"
	"os"
	"text/tabwriter" // https://pkg.go.dev/text/tabwriter#pkg-overview
)

func main() {
	const padding = 8
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned\t")
	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
	fmt.Fprintln(w, "aaaa\tbbbb\tcc")    // no trailing tab
	fmt.Fprintln(w, "aaaa\tbbbb\tccccc") // no trailing tab
	w.Flush()

	w = tabwriter.NewWriter(os.Stdout, 0, 0, padding, '.', 0) // default is left Aligned
	fmt.Println()
	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned\t")
	w.Flush()

	w = tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0) // default is left Aligned
	fmt.Println()
	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned\t")

	w.Flush()
}
