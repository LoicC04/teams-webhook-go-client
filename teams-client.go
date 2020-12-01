package teams_client

import (
	"bytes"
	"github.com/LoicC04/teams-webhook-go-client/model"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"text/template"
)

func Send(message model.MessageCard) {
	webhook := os.Getenv("TEAMS_WEBHOOK")
	if webhook == "" {
		log.Fatal("TEAMS_WEBHOOK environment variable is not set")
		os.Exit(2)
	}

	tmpl := template.Must(template.New("MessageCard").Funcs(template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
	}).Parse(MessageCardTemplate))
	var tpl bytes.Buffer
	tmpl.ExecuteTemplate(&tpl, "MessageCard", message)
	// Send request
	requestBody := tpl.String()
	res, err := http.Post(webhook, "application/json", strings.NewReader(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	// Read response
	res.Body.Close()

	// Convert response
	log.Println(res.Status)
}
