# boyer-moore-horspool
Implementation of the Boyer-Moore-Horspool algorithm or Horspool's algorithm in GO

## Technology
* [go](https://golang.org/)
* [ginkgo](https://github.com/onsi/ginkgo)

## Usage
CLI:
```bash
    go test -v test/
```

CODE:
```go
import "github.com/AlgoCafe/boyer-moore-horspool/bmh"

var index, length int
var buffer = []byte("abc abcdab abcdabcdabde")
var pattern = []byte("abcdabd") // 7
var pattern1 = []byte("abcdabdr") // 8

index = bmh.Search(buffer, pattern) // 15
index = bmh.Search(buffer, pattern1) // -1
```

## BMH Explications
 - [Explications en francais](doc/FR-EXPLICATION.md)
 - [English Explications](doc/EN-EXPLICATION.md)

## Links
  