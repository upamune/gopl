package ex03

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Echo1(args []string) {
	s, sep := "", ""

	for _, arg := range args {
		s += sep + arg
		sep = " "
	}

	fmt.Fprintln(ioutil.Discard, s)
}

func Echo2(args []string) {
	fmt.Fprintln(ioutil.Discard, strings.Join(args, " "))
}
