# simpleISBN
Go library for parsing, validating, and converting ISBN numbers.

## Features

- **Bidirectional conversion**
  - `ISBN-10` → `ISBN-13`
  - `ISBN-13` → `ISBN-10` (only for the `978` prefix)
- **Normalization**
  - Strips hyphens and `ISBN` prefix case‑insensitively
  - Accepts lowercase `x` or uppercase `X` in ISBN-10 check digit
- **Validation**
  - Verifies checksums for both ISBN-10 and ISBN-13
- **Convenience**
  - `String()` method for pretty-printing

## Installation

```bash
go get github.com/GianniBYoung/simpleISBN
```

## Usage

1. Import it:

```go
import "github.com/GianniBYoung/simpleISBN"
```



2. Create an Instance
```go
isbn, _ := simpleISBN.NewISBN("153432593X")

fmt.Println(isbn.ISBN13Number)
fmt.Println(isbn)
```

3. Or Just Convert an ISBN
```go
testISBN10 := "153432593X"
testISBN13 := "9781534325937"
// 10 -> 13
isbn_13, _ := simpleISBN.ConvertISBN(testISBN10, simpleISBN.ISBN13)
// 13 -> 10
isbn_10, _ := simpleISBN.ConvertISBN(testISBN13, simpleISBN.ISBN10)
fmt.Println("Here is the ISBN-13: " + isbn_13)
fmt.Println("Here is the ISBN-10: " + isbn_10)

```
