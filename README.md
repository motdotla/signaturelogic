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
options := signaturelogic.Options{}
signaturelogic.Setup("ORCHESTRATE_API_KEY", options)
```

### DocumentsCreate

```go
document := map[string]interface{}{"url": "http://mot.la/assets/resume.pdf"}
result, logic_error := signaturelogic.DocumentsCreate(app)
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

## Database Schema Details (using Redis)

signaturelogic uses a purposely simple database schema - as simple as possible. If you know a simpler approach, even better, please let me know or share as a pull request. 

signaturelogic uses Redis because of its light footprint, ephemeral nature, and lack of migrations.

* apps - collection of keys with all the app_names in there. SADD
* apps/myappname - hash with all the data in there. HSET or HMSET
* apps/theappname/identities - collection of keys with all the identities' emails in there. SADD
* apps/theappname/identities/emailaddress HSET or HMSET