# signaturelogic

<img src="https://raw.githubusercontent.com/motdotla/signaturelogic/master/signaturelogic.gif" alt="signaturelogic" align="right" width="190" />

[![BuildStatus](https://travis-ci.org/motdotla/signaturelogic.png?branch=master)](https://travis-ci.org/motdotla/signaturelogic)

Logic for saving [signature-api](https://github.com/motdotla/signature-api) data to the database.

This library is part of the larger [signature-api ecosystem](https://github.com/motdotla/signature-api).

## Usage

```go
package main

import (
  "log"
  "github.com/motdotla/signaturelogic"
)

func main() {
  signaturelogic.Setup("ORCHESTRATE_API_KEY")

  document := map[string]interface{}{"url": "http://mot.la/assets/resume.pdf"}
  result, logic_error := signaturelogic.DocumentsCreate(document)
  if logic_error != nil {
    log.Println(logic_error)
  }
  log.Println(result)
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

### DocumentsShow

```go
result, logic_error := signaturelogic.DocumentsShow("ID-OF-DOCUMENT")
```

### DocumentsUpdate

```go
id := "existing-id"
pages := []interface{}{}
page := map[string]interface{}{"number": 1, "url": "https://carvedevelopment.s3.amazonaws.com/87911158-edbc-488b-6e60-960d67809107/1.png"}
pages = append(pages, page)
document := map[string]interface{}{"pages": pages, "status": "processed", "id": id}

result, logic_error := signaturelogic.DocumentsUpdate(document)
```

### SigningsCreate

```go
signing := map[string]interface{}{"document_url": "https://signature-api.herokuapp.com/api/v0/documents/ef7ba0c7-dab7-425a-b849-d8157c40cd83.json"}
result, logic_error := signaturelogic.SigningsCreate(signing)
```

### SigningsShow

```go
result, logic_error := signaturelogic.SigningsShow("ID-OF-SIGNING")
```

### SigningsMarkSigned

```go
result, logic_error := signaturelogic.SigningsMarkSigned("ID-OF-SIGNING")
```

### SignatureElementsCreate

```go
signature_element := map[string]interface{}{"x": "20", "y": "20", "url": "data:image/gif;base64,R0lGODlhRAIEAaIAAOLi1v7..", "page_number": "1", "signing_id": "123456"}
result, logic_error := signaturelogic.SignatureElementsCreate(signature_element)
```

### SignatureElementsUpdate

```go
signature_element := map[string]interface{}{"x": "50", "y": "52", "id": "some-id"}

result, logic_error := signaturelogic.SignatureElementsUpdate(signature_element)
```

### SignatureElementsDelete

```go
result, logic_error := signaturelogic.SignatureElementsDelete(id)
```

### TextElementsCreate

```go
text_element := map[string]interface{}{"x": "20", "y": "20", "content": "Some content", "page_number": "1", "signing_id": "123456"}
result, logic_error := signaturelogic.TextElementsCreate(text_element)
```

### TextElementsUpdate

```go
text_element := map[string]interface{}{"x": "50", "y": "52", "id": "some-id"}

result, logic_error := signaturelogic.TextElementsUpdate(text_element)
```

### TextElementsDelete

```go
result, logic_error := signaturelogic.TextElementsDelete(id)
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
