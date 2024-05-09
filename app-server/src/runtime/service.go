package runtime

import (
	"errors"
	"fmt"

	"rogchap.com/v8go"
)

type RuntimeService struct {
	context *v8go.Context
}

func NewRuntimeService() *RuntimeService {
	context := v8go.NewContext()

	context.RunScript("const store = {text: '', number: 12}", "game_init.js")

	context.RunScript("function addToText(context) {const {params} = JSON.parse(context); store.text += params.text}", "def_add_to_text.js")
	context.RunScript("function clearText() {store.text = ''}", "def_clear_text.js")
	context.RunScript("function incrementNumber() {store.number += 1}", "def_increment_number.js")
	context.RunScript("function decrementNumber() {store.number -= 1}", "def_decrement_number.js")

	return &RuntimeService{context: context}
}

func (rs RuntimeService) GetGameContext() (*v8go.Object, error) {
	val, err := rs.context.RunScript("store", "get_store.js")

	if err != nil {
		return &v8go.Object{}, errors.New("JavaScript error: " + err.Error())
	}

	return val.Object(), nil
}

func (rs RuntimeService) InvokeMethod(method string, context string) {
	script := fmt.Sprintf("%s('%s')", method, context)
	rs.context.RunScript(script, method+".js")
}
