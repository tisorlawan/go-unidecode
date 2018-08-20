go-unidecode
==============

[![Build Status](https://travis-ci.org/tisorlawan/go-unidecode.svg?branch=master)](https://travis-ci.org/tisorlawan/go-unidecode)
[![Coverage Status](https://coveralls.io/repos/tisorlawan/go-unidecode/badge.svg?branch=master)](https://coveralls.io/r/tisorlawan/go-unidecode?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/tisorlawan/go-unidecode)](https://goreportcard.com/report/github.com/tisorlawan/go-unidecode)
[![GoDoc](https://godoc.org/github.com/tisorlawan/go-unidecode?status.svg)](https://godoc.org/github.com/tisorlawan/go-unidecode)

ASCII transliterations of Unicode text.


Installation
------------

```
go get -u github.com/tisorlawan/go-unidecode
```

Install CLI tool:

```
go get -u github.com/tisorlawan/go-unidecode/unidecode
$ unidecode 北京
Bei Jing 
```


Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/tisorlawan/go-unidecode


Usage
------

```go
package main

import (
	"fmt"
	"github.com/tisorlawan/go-unidecode"
)

func main() {
	s := "abc"
	fmt.Println(unidecode.Unidecode(s))
	// Output: abc

	s = "北京"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing

	s = "kožušček"
	fmt.Println(unidecode.Unidecode(s))
	// Output: kozuscek
}
```
