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

1. Import it:

```go
import simpleISBN "github.com/GianniBYoung/simpleISBN"
```



2. Create an Instance
```go
isbn, err := simpleISBN.NewISBN(raw)
if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println(isbn)
```

3. Or Just Convert an ISBN
```go
testISBN10 := "153432593X"
testISBN13 := "9781534325937"
// 10 -> 13
isbn_13, err := simpleISBN.ConvertISBN(testISBN10, ISBN13)
// 13 -> 10
isbn_10, err := simpleISBN.ConvertISBN(testISBN13, ISBN10)

```
