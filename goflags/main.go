package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type HostList []string

func (hl *HostList) String() string {
	return strings.Join(*hl, ",")
}

func (hl *HostList) Set(list string) error {
	for _, host := range strings.Split(list, ",") {
		*hl = append(*hl, host)
	}
	return nil
}

var (
	// Example1: using flag.Type() return a pointer to underlying data
	ptrFlag = flag.Bool("boolptr", false, "a pointer var point to bool")

	// Example2: using flag.TypeVar() store data in bound value
	valFlag string

	// Example3: custom type implement flag.Value interface
	customTypeFlag HostList

	// Example4: subcommand
	commonFlag int
	fooSubCmd  = flag.NewFlagSet("foo", flag.ExitOnError)
	fooPtrFlag = fooSubCmd.String("fooptr", "foosub", "a pointer var point to string")
	barSubCmd  = flag.NewFlagSet("bar", flag.ExitOnError)
	barPtrFlag = barSubCmd.String("barptr", "barsub", "a pointer var point to string")
)

func init() {
	flag.StringVar(&valFlag, "strval", "String Value", "a value var store string")
	flag.Var(&customTypeFlag, "customtype", "a custom type store comma seperated string")
	setupCommonFlag()
}

// setupCommonFlag register common flags for subcommand
func setupCommonFlag() {
	for _, subcmd := range []*flag.FlagSet{fooSubCmd, barSubCmd} {
		subcmd.IntVar(&commonFlag, "commonflag", 10000, "subcommand common flag")
	}
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "foo":
			fooSubCmd.Parse(os.Args[2:])
		case "bar":
			barSubCmd.Parse(os.Args[2:])
		default:
			flag.Parse()
		}
	}
	fmt.Printf(`ptrFlag = %v
valFlag = %v
customTypeFlag = %v
commonFlag = %v
fooPtrFlag = %v
barPtrFlag = %v
`, *ptrFlag, valFlag, customTypeFlag, commonFlag, *fooPtrFlag, *barPtrFlag)
}
