package rules

import (
	"bytes"
	"errors"
	"log/slog"
	"sync"
	"text/template"
)

type templateVariables[T any] struct {
	PropertyValue T
	MinLength     int
	MaxLength     int
}

func returnTemplatedError[T any](tpl *template.Template, getVars func() templateVariables[T]) error {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, getVars()); err != nil {
		slog.Error("failed to execute message template",
			slog.String("template", tpl.Name()),
			slog.String("error", err.Error()))
	}
	return errors.New(buf.String())
}

var messageTemplates = messageTemplatesMap{
	tmpl: make(map[string]*template.Template),
	mu:   sync.RWMutex{},
}

type messageTemplatesMap struct {
	tmpl map[string]*template.Template
	mu   sync.RWMutex
}

func (p *messageTemplatesMap) Lookup(name string) *template.Template {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.tmpl[name]
}

func (p *messageTemplatesMap) Register(name string, tpl *template.Template) {
	p.mu.Lock()
	p.tmpl[name] = tpl
	p.mu.Unlock()
}

func getMessageTemplate(name, msg string) *template.Template {
	if tpl := messageTemplates.Lookup(name); tpl != nil {
		return tpl
	}
	tpl := template.Must(template.New(name).Parse(msg))
	messageTemplates.Register(name, tpl)
	return tpl
}
