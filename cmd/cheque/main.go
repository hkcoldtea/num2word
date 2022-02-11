package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hkcoldtea/num2word"
)

var (
	BUILD string
)

func main() {
	var amount_in_words string
	var bUpper bool
	flag.BoolVar(&bUpper, "u", false, "Presenting output in all uppercase")
	flag.Usage = func() {
		w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily
		if len(BUILD) > 0 {
			fmt.Fprintf(w, "Build: %s\n", BUILD)
		}
		fmt.Fprintf(w, "Usage of %s: [amount]\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(w, "\n")
	}
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		return
	}
	for _, amount := range flag.Args() {
		amount_in_words = num2word.Translate_full_amount(amount)
		if bUpper {
			amount_in_words = strings.ToUpper(amount_in_words)
		}
		fmt.Println(amount_in_words)
	}
}
