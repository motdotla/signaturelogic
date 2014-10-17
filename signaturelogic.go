package signaturelogic

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/handshakejs/handshakejserrors"
	"github.com/orchestrate-io/gorc"
	"strings"
)

const (
	DOCUMENTS          = "documents"
	SIGNATURE_ELEMENTS = "signature_elements"
	TEXT_ELEMENTS      = "text_elements"
	SIGNINGS           = "signings"
	PROCESSING         = "processing"
)

type SignatureElement struct {
	X          string
	Y          string
	Url        string
	PageNumber string
	SigningId  string
}

var (
	ORCHESTRATE_API_KEY string
	client              *gorc.Client
)

func Setup(orchestrate_api_key string) {
	ORCHESTRATE_API_KEY = orchestrate_api_key
}

func DocumentsShow(id string) (map[string]interface{}, *handshakejserrors.LogicError) {
	conn := Conn()
	result, err := conn.Get(DOCUMENTS, id)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	document := make(map[string]interface{})
	err = result.Value(&document)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return document, nil
}

func DocumentsUpdate(new_document map[string]interface{}) (map[string]interface{}, *handshakejserrors.LogicError) {

	document, logic_error := DocumentsShow(new_document["id"].(string))
	if logic_error != nil {
		return nil, logic_error
	}
	document["pages"] = new_document["pages"]
	document["status"] = new_document["status"]

	conn := Conn()
	_, err := conn.Put(DOCUMENTS, document["id"].(string), document)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return document, nil
}

func DocumentsCreate(document map[string]interface{}) (map[string]interface{}, *handshakejserrors.LogicError) {
	var url string
	if str, ok := document["url"].(string); ok {
		url = strings.Replace(str, " ", "", -1)
	} else {
		url = ""
	}
	if url == "" {
		logic_error := &handshakejserrors.LogicError{"required", "url", "url cannot be blank"}
		return document, logic_error
	}
	pages := []string{}
	document["pages"] = pages
	document["status"] = PROCESSING
	document["url"] = url
	key := uuid.New()
	document["id"] = key

	conn := Conn()
	_, err := conn.Put(DOCUMENTS, document["id"].(string), document)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return document, nil
}

func SigningsShow(id string) (map[string]interface{}, *handshakejserrors.LogicError) {
	conn := Conn()
	result, err := conn.Get(SIGNINGS, id)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	signing := make(map[string]interface{})
	err = result.Value(&signing)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	conn = Conn()
	results, err := conn.Search(SIGNATURE_ELEMENTS, "signing_id:"+id, 1000, 0)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	// here is where I need to get all the signature_elements
	signature_elements := make([]SignatureElement, len(results.Results))
	for i, item := range results.Results {
		item.Value(&signature_elements[i])
	}
	signing["signature_elements"] = signature_elements
	// here is where I need to get all the text_elements

	return signing, nil
}

func SigningsCreate(signing map[string]interface{}) (map[string]interface{}, *handshakejserrors.LogicError) {
	var document_url string
	if str, ok := signing["document_url"].(string); ok {
		document_url = strings.Replace(str, " ", "", -1)
	} else {
		document_url = ""
	}
	if document_url == "" {
		logic_error := &handshakejserrors.LogicError{"required", "document_url", "document_url cannot be blank"}
		return signing, logic_error
	}

	signature_elements := []string{}
	text_elements := []string{}
	signing["signature_elements"] = signature_elements
	signing["text_elements"] = text_elements
	signing["document_url"] = document_url
	key := uuid.New()
	signing["id"] = key

	conn := Conn()
	_, err := conn.Put(SIGNINGS, signing["id"].(string), signing)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return signing, nil
}

func SignatureElementsCreate(signature_element map[string]interface{}) (map[string]interface{}, *handshakejserrors.LogicError) {
	var x string
	var y string
	var url string
	var page_number string
	var signing_id string
	if str, ok := signature_element["x"].(string); ok {
		x = strings.Replace(str, " ", "", -1)
	} else {
		x = ""
	}
	if x == "" {
		logic_error := &handshakejserrors.LogicError{"required", "x", "x cannot be blank"}
		return signature_element, logic_error
	}

	if str, ok := signature_element["y"].(string); ok {
		y = strings.Replace(str, " ", "", -1)
	} else {
		y = ""
	}
	if y == "" {
		logic_error := &handshakejserrors.LogicError{"required", "y", "y cannot be blank"}
		return signature_element, logic_error
	}

	if str, ok := signature_element["url"].(string); ok {
		url = strings.Replace(str, " ", "", -1)
	} else {
		url = ""
	}
	if url == "" {
		logic_error := &handshakejserrors.LogicError{"required", "url", "url cannot be blank"}
		return signature_element, logic_error
	}

	if str, ok := signature_element["page_number"].(string); ok {
		page_number = strings.Replace(str, " ", "", -1)
	} else {
		page_number = ""
	}
	if page_number == "" {
		logic_error := &handshakejserrors.LogicError{"required", "page_number", "page_number cannot be blank"}
		return signature_element, logic_error
	}

	if str, ok := signature_element["signing_id"].(string); ok {
		signing_id = strings.Replace(str, " ", "", -1)
	} else {
		signing_id = ""
	}
	if signing_id == "" {
		logic_error := &handshakejserrors.LogicError{"required", "signing_id", "signing_id cannot be blank"}
		return signature_element, logic_error
	}
	signature_element["x"] = x
	signature_element["y"] = y
	signature_element["url"] = url
	signature_element["page_number"] = page_number
	signature_element["signing_id"] = signing_id
	key := uuid.New()
	signature_element["id"] = key

	conn := Conn()
	_, err := conn.Put(SIGNATURE_ELEMENTS, signature_element["id"].(string), signature_element)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return signature_element, nil
}

func SignatureElementsShow(id string) (map[string]interface{}, *handshakejserrors.LogicError) {
	conn := Conn()
	result, err := conn.Get(SIGNATURE_ELEMENTS, id)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	signature_element := make(map[string]interface{})
	err = result.Value(&signature_element)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return signature_element, nil
}

func SignatureElementsUpdate(new_signature_element map[string]interface{}) (map[string]interface{}, *handshakejserrors.LogicError) {
	signature_element, logic_error := SignatureElementsShow(new_signature_element["id"].(string))
	if logic_error != nil {
		return nil, logic_error
	}

	var x string
	var y string
	if str, ok := new_signature_element["x"].(string); ok {
		x = strings.Replace(str, " ", "", -1)
	} else {
		x = ""
	}
	if x == "" {
		logic_error := &handshakejserrors.LogicError{"required", "x", "x cannot be blank"}
		return new_signature_element, logic_error
	}

	if str, ok := new_signature_element["y"].(string); ok {
		y = strings.Replace(str, " ", "", -1)
	} else {
		y = ""
	}
	if y == "" {
		logic_error := &handshakejserrors.LogicError{"required", "y", "y cannot be blank"}
		return new_signature_element, logic_error
	}

	signature_element["x"] = x
	signature_element["y"] = y

	conn := Conn()
	_, err := conn.Put(SIGNATURE_ELEMENTS, signature_element["id"].(string), signature_element)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return signature_element, nil
}

func TextElementsCreate(text_element map[string]interface{}) (map[string]interface{}, *handshakejserrors.LogicError) {
	var x string
	var y string
	var content string
	var page_number string
	var signing_id string
	if str, ok := text_element["x"].(string); ok {
		x = strings.Replace(str, " ", "", -1)
	} else {
		x = ""
	}
	if x == "" {
		logic_error := &handshakejserrors.LogicError{"required", "x", "x cannot be blank"}
		return text_element, logic_error
	}

	if str, ok := text_element["y"].(string); ok {
		y = strings.Replace(str, " ", "", -1)
	} else {
		y = ""
	}
	if y == "" {
		logic_error := &handshakejserrors.LogicError{"required", "y", "y cannot be blank"}
		return text_element, logic_error
	}

	content = text_element["content"].(string)
	if content == "" {
		logic_error := &handshakejserrors.LogicError{"required", "content", "content cannot be blank"}
		return text_element, logic_error
	}

	if str, ok := text_element["page_number"].(string); ok {
		page_number = strings.Replace(str, " ", "", -1)
	} else {
		page_number = ""
	}
	if page_number == "" {
		logic_error := &handshakejserrors.LogicError{"required", "page_number", "page_number cannot be blank"}
		return text_element, logic_error
	}

	if str, ok := text_element["signing_id"].(string); ok {
		signing_id = strings.Replace(str, " ", "", -1)
	} else {
		signing_id = ""
	}
	if signing_id == "" {
		logic_error := &handshakejserrors.LogicError{"required", "signing_id", "signing_id cannot be blank"}
		return text_element, logic_error
	}
	text_element["x"] = x
	text_element["y"] = y
	text_element["content"] = content
	text_element["page_number"] = page_number
	text_element["signing_id"] = signing_id
	key := uuid.New()
	text_element["id"] = key

	conn := Conn()
	_, err := conn.Put(TEXT_ELEMENTS, text_element["id"].(string), text_element)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return text_element, nil
}

func TextElementsShow(id string) (map[string]interface{}, *handshakejserrors.LogicError) {
	conn := Conn()
	result, err := conn.Get(TEXT_ELEMENTS, id)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	signature_element := make(map[string]interface{})
	err = result.Value(&signature_element)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return signature_element, nil
}

func TextElementsUpdate(new_text_element map[string]interface{}) (map[string]interface{}, *handshakejserrors.LogicError) {
	text_element, logic_error := TextElementsShow(new_text_element["id"].(string))
	if logic_error != nil {
		return nil, logic_error
	}

	var x string
	var y string
	if str, ok := new_text_element["x"].(string); ok {
		x = strings.Replace(str, " ", "", -1)
	} else {
		x = ""
	}
	if x == "" {
		logic_error := &handshakejserrors.LogicError{"required", "x", "x cannot be blank"}
		return new_text_element, logic_error
	}

	if str, ok := new_text_element["y"].(string); ok {
		y = strings.Replace(str, " ", "", -1)
	} else {
		y = ""
	}
	if y == "" {
		logic_error := &handshakejserrors.LogicError{"required", "y", "y cannot be blank"}
		return new_text_element, logic_error
	}

	text_element["x"] = x
	text_element["y"] = y

	conn := Conn()
	_, err := conn.Put(TEXT_ELEMENTS, text_element["id"].(string), text_element)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return text_element, nil
}
func Conn() *gorc.Client {
	client := gorc.NewClient(ORCHESTRATE_API_KEY)
	return client
}
