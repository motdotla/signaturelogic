package signaturelogic_test

import (
	"../signaturelogic"
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
