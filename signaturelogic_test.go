package signaturelogic_test

import (
	"../signaturelogic"
	"github.com/handshakejs/handshakejserrors"
	"github.com/joho/godotenv"
	"github.com/orchestrate-io/gorc"
	"log"
	"os"
	"testing"
	//"time"
)

const (
	URL          = "http://mot.la/assets/resume.pdf"
	X            = "20"
	Y            = "20"
	DATA_URL     = "data:image/gif;base64,R0lGODlhRAIEAaIAAOLi1v7+5enp2ubm2Pf34e7u3QAAAAAAACH5BAAHAP8ALAAAAABEAgQBAAP/GLrc/jDKSau9OOvNu/9gKI5kaZ5oqq5s675wLM90bd94ru987//AoHBILBqPyKRyyWw6n9CodEqtWq/YrHbL7Xq/4LB4TC6bz+i0es1uu9/wuHxOr9vv+Lx+z+/7/4CBgoOEhYaHiImKi4yNjo+QkZKTlJWWl5iZmpucnZ6foKGio6SlpqeoqaqrrK2ur7CxsrO0tba3uLm6u7y9vr/AwcLDxMXGx8jJysvMzc7P0NHS09TV1tfY2drb3N3e3+Dh4uPk5ebn6Onq6+zt7u/w8fLz9PX29/j5+vv8/f7/AAMKHEiwoMGDCBMqXMiwocOHECNKnEixosWLGDNq3Mix/6PHjyBDihxJsqTJkyhTqlzJsqXLlzBjypxJs6bNmzhz6tzJs6fPn0CDCh1KtKjRo0iTKl3KtKnTp1CjSp1KtarVq1izat3KtavXr2DDih1LtqzZs2jTql3Ltq3bt3Djyp1Lt67du3jz6t3Lt6/fv4ADCx5MuLDhw4gTK17MuLHjx5AjS55MubLly5gza97MubPnz6BDix5NurTp06hTq17NurXr17Bjy55Nu7bt27hz697Nu7fv38CDCx9OvLjx48iTK1/OvLnz59CjS59Ovbr169g5CADAnUCE7QAEZE9DgDuAARAKmB+vZoB57w3Ud2dP3rx4BuXn009jHgD8AP/5AVDAfmrIF94C5g1IoBr9eQfefQumYWABBkbIRn/vWbgGeBlqqEaAAnq4BogKingGiNyZiAaG+qk4xoMBoueiGPLJ2OCMYBgIn4EQ4rhFgP8FcKOPWgRYogITEqlFgg/0pyQWD6bHZAMsYuhAlVZSieV6Wm4JwJVeftnllmB6WSaZY2J5ppppVrmmm22y+KaccWbJQJhi3hnmnHYiuGedTgLKpZ5mCpqioXn6WSihaDLKpqNwQkrnC1FGEKiklyraqKaPchqpp5OC2qcCePKZKal/YnqqkKmKumqpiJo6qKuzoroorYeqWiurt9q6qa+dAvupsKESOyqvvyIbrLKKwzJbLAsERDtBtNIaKmuuuCZq7KutbrsrrLpi6624zh4LbrbXalsut72u+2237pJ77rjqzhtvvfDaq2++/LZr75MAByzwwAQXbPDBCCes8MIMN+zwwxBHLPHEFFds8cUYZ6zxxhx37PHHIIcs8sgkl2zyySinrPLKLLfs8sswxyzzzDTXbPPNbiUAADs="
	PAGE_NUMBER  = "1"
	CONTENT      = "Some Content"
	DOCUMENT_URL = "https://signature-api.herokuapp.com/api/v0/documents/ef7ba0c7-dab7-425a-b849-d8157c40cd83.json"
	SIGNING_ID   = "67890"
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

//func TestDocumentsUpdate(t *testing.T) {
//	setup(t)
//	tearDown(t)
//
//	result, logic_error := createDocument(t)
//	if logic_error != nil {
//		t.Errorf("Error", logic_error)
//	}
//
//	id := result["id"].(string)
//	log.Println(id)
//	pages := []interface{}{}
//	page := map[string]interface{}{"number": 1, "url": "https://carvedevelopment.s3.amazonaws.com/87911158-edbc-488b-6e60-960d67809107/1.png"}
//	pages = append(pages, page)
//	document := map[string]interface{}{"id": id, "pages": pages, "status": "processed"}
//
//	time.Sleep(3000 * time.Millisecond)
//	log.Println(document)
//
//	result, logic_error = signaturelogic.DocumentsUpdate(document)
//	if logic_error != nil {
//		t.Errorf("Error", logic_error)
//	}
//
//	if result["status"].(string) != "processed" {
//		t.Errorf("status did not equal processed")
//	}
//}

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

func TestSigningsCreate(t *testing.T) {
	setup(t)
	tearDown(t)

	signing := map[string]interface{}{"document_url": DOCUMENT_URL}

	signaturelogic.Setup(os.Getenv("ORCHESTRATE_API_KEY"))
	result, logic_error := signaturelogic.SigningsCreate(signing)
	if logic_error != nil {
		t.Errorf("Error", logic_error)
	}
	if result["document_url"] != DOCUMENT_URL {
		t.Errorf("Incorrect document_url " + result["document_url"].(string))
	}
	if result["id"] == nil {
		t.Errorf("Missing ID")
	}
	if result["id"] == "" {
		t.Errorf("Missing ID")
	}
	if result["status"] != "signing" {
		t.Errorf("Incorrect status " + result["status"].(string))
	}
}

func TestSignatureElementsCreate(t *testing.T) {
	setup(t)
	tearDown(t)

	signature_element := map[string]interface{}{"x": X, "y": Y, "url": DATA_URL, "page_number": PAGE_NUMBER, "signing_id": SIGNING_ID}

	signaturelogic.Setup(os.Getenv("ORCHESTRATE_API_KEY"))
	result, logic_error := signaturelogic.SignatureElementsCreate(signature_element)
	if logic_error != nil {
		t.Errorf("Error", logic_error)
	}
	if result["signing_id"] != SIGNING_ID {
		t.Errorf("Incorrect signing_id " + result["signing_id"].(string))
	}
	if result["url"] != DATA_URL {
		t.Errorf("Incorrect url " + result["url"].(string))
	}
	if result["x"] != X {
		t.Errorf("Incorrect x " + result["x"].(string))
	}
	if result["y"] != Y {
		t.Errorf("Incorrect y " + result["y"].(string))
	}
	if result["page_number"] != PAGE_NUMBER {
		t.Errorf("Incorrect page_number " + result["page_number"].(string))
	}
	if result["id"] == nil {
		t.Errorf("Missing ID")
	}
	if result["id"] == "" {
		t.Errorf("Missing ID")
	}
}

func TestTextElementsCreate(t *testing.T) {
	setup(t)
	tearDown(t)

	text_element := map[string]interface{}{"x": X, "y": Y, "content": CONTENT, "page_number": PAGE_NUMBER, "signing_id": SIGNING_ID}

	signaturelogic.Setup(os.Getenv("ORCHESTRATE_API_KEY"))
	result, logic_error := signaturelogic.TextElementsCreate(text_element)
	if logic_error != nil {
		t.Errorf("Error", logic_error)
	}
	if result["signing_id"] != SIGNING_ID {
		t.Errorf("Incorrect signing_id " + result["signing_id"].(string))
	}
	if result["content"] != CONTENT {
		t.Errorf("Incorrect content " + result["content"].(string))
	}
	if result["x"] != X {
		t.Errorf("Incorrect x " + result["x"].(string))
	}
	if result["y"] != Y {
		t.Errorf("Incorrect y " + result["y"].(string))
	}
	if result["page_number"] != PAGE_NUMBER {
		t.Errorf("Incorrect page_number " + result["page_number"].(string))
	}
	if result["id"] == nil {
		t.Errorf("Missing ID")
	}
	if result["id"] == "" {
		t.Errorf("Missing ID")
	}
}

func TestSigningsShow(t *testing.T) {
	setup(t)
	tearDown(t)
	result, logic_error := createSigning(t)
	if logic_error != nil {
		t.Errorf("createSigning failed.")
	}

	id := result["id"].(string)
	result, logic_error = signaturelogic.SigningsShow(id)
	if logic_error != nil {
		t.Errorf("SigningsShow failed", logic_error)
	}
	if result["document_url"].(string) != DOCUMENT_URL {
		t.Errorf("returned document_url was incorrect")
	}
}

func TestSigningsMarkSigned(t *testing.T) {
	setup(t)
	tearDown(t)
	result, logic_error := createSigning(t)
	if logic_error != nil {
		t.Errorf("createSigning failed.")
	}

	id := result["id"].(string)
	result, logic_error = signaturelogic.SigningsMarkSigned(id)
	if logic_error != nil {
		t.Errorf("SigningsMarkSigned failed", logic_error)
	}

	if result["status"].(string) != "signed" {
		t.Errorf("returned status was incorrect")
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

func createSigning(t *testing.T) (map[string]interface{}, *handshakejserrors.LogicError) {
	signing := map[string]interface{}{"document_url": DOCUMENT_URL}

	signaturelogic.Setup(os.Getenv("ORCHESTRATE_API_KEY"))
	result, logic_error := signaturelogic.SigningsCreate(signing)
	if logic_error != nil {
		return nil, logic_error
	}

	return result, nil
}
func tearDown(t *testing.T) {
	orchestrate_api_key := os.Getenv("ORCHESTRATE_API_KEY")
	o := gorc.NewClient(orchestrate_api_key)
	o.DeleteCollection("documents")
	o.DeleteCollection("signings")
	o.DeleteCollection("signature_elements")
	o.DeleteCollection("text_elements")
}

func setup(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
