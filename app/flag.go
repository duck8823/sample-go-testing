package app

import (
	"flag"
	"os"
)

func ParseFlag(args ...string) *string {
	flg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	optC := flg.String("c", "default value", "flag usage")
	flg.Parse(args)

	return optC
}
