# bloomy: simple bloom filter in pure go.

Bloomy is a simple bloom filter written in go. A bloom filter is space-efficient data-structure which tests whether an element is a part of a given list of elements. It is important to note that false positives are possible, but false negatives are not.

## Usage

Install the project as a dependency:

```
go get -u github.com/nireo/bloomy
```

```go
package main

import (
        "fmt"
        "github.com/nireo/bloomy"
)

const itemCount = 20000
const falsePositiveRate = 0.01

func main() {
    bf := bloomy.New(itemCount, falsePositiveRate) 
    bf.Insert([]byte("hello world"))

    if bf.Contains([]byte("hello world")) {
        fmt.Println("found string")
    } else {
        fmt.Println("something went wrong")
    }
}
```
