[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

# boyer-moore-horspool
<strong>[EN]</strong>  
Implementation of the Boyer-Moore-Horspool algorithm or Horspool's algorithm.

Boyer-Moore-Horspool is a substring search algorithm.

<strong>[FR]</strong>  
Implémentation de l'algorithme de Boyer-Moore-Horspool.

Boyer-Moore-Horspool est un algorithme de recherche de sous-chaîne.

## Explanation
 - [English](documentation/explanation.en.md)
 - [Français](documentation/explanation.fr.md)

## Technology
* [go](https://golang.org/)
* [ginkgo](https://github.com/onsi/ginkgo)

## Usage
CLI:
```bash
go test -v ./test
```

CODE:
```go
import "github.com/coursecomputer/boyer-moore-horspool/source"

var index, length int
var buffer = []byte("abc abcdab abcdabcdabde")
var pattern = []byte("abcdabd") // 7
var pattern1 = []byte("abcdabdr") // 8

index = bmh.Search(buffer, pattern) // 15
index = bmh.Search(buffer, pattern1) // -1
```

## Links
https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm  
http://www.mathcs.emory.edu/~cheung/Courses/323/Syllabus/Text/Matching-Boyer-Moore2.html  