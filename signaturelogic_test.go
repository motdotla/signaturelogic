package signaturelogic_test

import (
	"../signaturelogic"
	"github.com/handshakejs/handshakejserrors"
	"github.com/joho/godotenv"
	"github.com/orchestrate-io/gorc"
	"log"
	"os"
	"testing"
)

const (
	URL = "http://mot.la/assets/resume.pdf"
)

func TestDocumentsCreate(t *testing.T) {
	setup(t)
	tearDown(t)

	document := map[string]interface{}{"url": URL}

	signaturelogic.Setup(os.Getenv("ORCHESTRATE_API_KEY"))
	result, logic_error := signaturelogic.DocumentsCreate(document)
	if logic_error != nil {
		t.Errorf("Error", logic_error)
	}
	if result["url"] != URL {
		t.Errorf("Incorrect url " + result["url"].(string))
	}
	if result["status"] != "processing" {
		t.Errorf("Incorrect status " + result["status"].(string))
	}
	if result["id"] == nil {
		t.Errorf("Missing ID")
	}
	if result["id"] == "" {
		t.Errorf("Missing ID")
	}
}

func TestDocumentsCreateNilUrl(t *testing.T) {
	setup(t)
	tearDown(t)

	document := map[string]interface{}{}

	signaturelogic.Setup(os.Getenv("ORCHESTRATE_API_KEY"))
	_, logic_error := signaturelogic.DocumentsCreate(document)
	if logic_error.Code != "required" {
		t.Errorf("Error", "Logic error should have been 'required'")
	}
}

func TestDocumentsCreateBlankUrl(t *testing.T) {
	setup(t)
	tearDown(t)

	document := map[string]interface{}{"url": ""}

	signaturelogic.Setup(os.Getenv("ORCHESTRATE_API_KEY"))
	_, logic_error := signaturelogic.DocumentsCreate(document)
	if logic_error.Code != "required" {
		t.Errorf("Error", "Logic error should have been 'required'")
	}
}

func TestDocumentsCreateInvalidOrchestrateApiKey(t *testing.T) {
	setup(t)
	tearDown(t)

	document := map[string]interface{}{"url": URL}

	signaturelogic.Setup("invalid-orchestrate-api-key")
	_, logic_error := signaturelogic.DocumentsCreate(document)
	if logic_error.Code != "unknown" {
		t.Errorf("Error code was not 'unknown'")
	}
}

func TestDocumentsShow(t *testing.T) {
	setup(t)
	tearDown(t)
	result, logic_error := createDocument(t)
	if logic_error != nil {
		t.Errorf("createDocument failed.")
	}

	id := result["id"].(string)
	result, _ = signaturelogic.DocumentsShow(id)
	if result["url"].(string) != URL {
		t.Errorf("returned url was incorrect")
	}
}

func TestDocumentsShowWrongId(t *testing.T) {
	setup(t)
	tearDown(t)
	_, logic_error := createDocument(t)
	if logic_error != nil {
		t.Errorf("createDocument failed.")
	}

	_, logic_error = signaturelogic.DocumentsShow("wrong-id")
	if logic_error == nil {
		t.Errorf("logic error should have been raised")
	}
}

func TestDocumentsUpdate(t *testing.T) {
	setup(t)
	tearDown(t)

	result, logic_error := createDocument(t)
	if logic_error != nil {
		t.Errorf("Error", logic_error)
	}

	id := result["id"].(string)
	pages := []interface{}{}
	page := map[string]interface{}{"number": 1, "url": "https://carvedevelopment.s3.amazonaws.com/87911158-edbc-488b-6e60-960d67809107/1.png"}
	pages = append(pages, page)
	document := map[string]interface{}{"id": id, "pages": pages, "status": "processed"}

	result, logic_error = signaturelogic.DocumentsUpdate(document)
	if logic_error != nil {
		t.Errorf("Error", logic_error)
	}

	if result["status"].(string) != "processed" {
		t.Errorf("status did not equal processed")
	}
}

func TestDocumentsUpdateIdDoesNotExist(t *testing.T) {
	setup(t)
	tearDown(t)

	id := "does-not-exist"
	pages := []interface{}{}
	page := map[string]interface{}{"number": 1, "url": "https://carvedevelopment.s3.amazonaws.com/87911158-edbc-488b-6e60-960d67809107/1.png"}
	pages = append(pages, page)
	document := map[string]interface{}{"id": id, "pages": pages, "status": "processed"}

	_, logic_error := signaturelogic.DocumentsUpdate(document)
	if logic_error.Code != "unknown" {
		t.Errorf("Error", "Logic error should have been 'unknown'")
	}
}

func createDocument(t *testing.T) (map[string]interface{}, *handshakejserrors.LogicError) {
	document := map[string]interface{}{"url": URL}

	signaturelogic.Setup(os.Getenv("ORCHESTRATE_API_KEY"))
	result, logic_error := signaturelogic.DocumentsCreate(document)
	if logic_error != nil {
		return nil, logic_error
	}

	return result, nil
}

func tearDown(t *testing.T) {
	orchestrate_api_key := os.Getenv("ORCHESTRATE_API_KEY")
	o := gorc.NewClient(orchestrate_api_key)
	o.DeleteCollection("documents")
}

func setup(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
