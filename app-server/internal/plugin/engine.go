package engine

import (
	"errors"
	"fmt"

	"rogchap.com/v8go"
)

type Engine struct {
	context     *v8go.Context
	initialized bool
}

func NewEngine() *Engine {
	context := v8go.NewContext()

	return &Engine{context: context}
}

func (rs Engine) Reset() {
	rs.context = v8go.NewContext()
	rs.initialized = false
}

func (rs Engine) Initialize(initialState string, methods map[string]string) error {
	if rs.initialized {
		return errors.New("Already initialized")
	}

	rs.context.RunScript(initialState, "state_init.js")

	for name, source := range methods {
		rs.context.RunScript(source, name)
	}

	rs.initialized = true

	return nil
}

func (rs Engine) GetState() (*v8go.Object, error) {
	if !rs.initialized {
		return &v8go.Object{}, errors.New("Not initialized")
	}

	val, err := rs.context.RunScript("store", "get_store.js")

	if err != nil {
		return &v8go.Object{}, errors.New("JavaScript error: " + err.Error())
	}

	return val.Object(), nil
}

func (rs Engine) Run(method string, context string) error {
	if !rs.initialized {
		errors.New("Not initialized")
	}

	script := fmt.Sprintf("%s('%s')", method, context)
	rs.context.RunScript(script, method+".js")

	return nil
}
