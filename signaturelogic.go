package signaturelogic

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/handshakejs/handshakejserrors"
	"github.com/orchestrate-io/gorc"
	"strings"
)

const (
	DOCUMENTS  = "documents"
	PROCESSING = "processing"
)

var (
	ORCHESTRATE_API_KEY string
	client              *gorc.Client
)

func Setup(orchestrate_api_key string) {
	ORCHESTRATE_API_KEY = orchestrate_api_key
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

	conn := Conn()

	key := uuid.New()
	_, err := conn.Put(DOCUMENTS, key, document)
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return nil, logic_error
	}

	return document, nil
}

func Conn() *gorc.Client {
	client := gorc.NewClient(ORCHESTRATE_API_KEY)
	return client
}
