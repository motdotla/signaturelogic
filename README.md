# signaturelogic

<img src="https://raw.githubusercontent.com/motdotla/signaturelogic/master/signaturelogic.gif" alt="signaturelogic" align="right" width="190" />

[![BuildStatus](https://travis-ci.org/motdotla/signaturelogic.png?branch=master)](https://travis-ci.org/motdotla/signaturelogic)

Logic for saving [signature-api](https://github.com/motdotla/signature-api) data to the database.

This library is part of the larger [signature-api ecosystem](https://github.com/motdotla/signature-api).

## Usage

```go
package main

import (
  "fmt"
  "github.com/motdotla/signaturelogic"
)

func main() {
  signaturelogic.Setup("ORCHESTRATE_API_KEY")

  document := map[string]interface{}{"url": "http://mot.la/assets/resume.pdf"}
  result, logic_error := signaturelogic.DocumentsCreate(document)
  if logic_error != nil {
    fmt.Println(logic_error)
  }
  fmt.Println(result)
}
```

### Setup

Connects to Orchestrate.

```go
signaturelogic.Setup("ORCHESTRATE_API_KEY")
```

### DocumentsCreate

```go
document := map[string]interface{}{"url": "http://mot.la/assets/resume.pdf"}
result, logic_error := signaturelogic.DocumentsCreate(document)
```

## Installation

```
go get github.com/motdotla/signaturelogic
```

## Running Tests

```
cp .env.example .env
```

Edit the contents of `.env`.

```
go test -v
```
