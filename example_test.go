package unidecode_test

import (
	"fmt"

	"github.com/tisorlawan/go-unidecode"
)

func ExampleUnidecode() {
	s := "北京kožušček"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing kozuscek
}

func ExampleUnidecode_aSCII() {
	s := "abc"
	fmt.Println(unidecode.Unidecode(s))
	// Output: abc
}
