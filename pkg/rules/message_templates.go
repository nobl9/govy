package rules

import (
	"bytes"
	"fmt"
	"log/slog"
	"sync"
	"text/template"
)

type templateKey int

const (
	stringLengthTemplateKey templateKey = iota + 1
)

var rawMessageTemplates = map[templateKey]string{
	stringLengthTemplateKey: "length must be between {{ .MinLength }} and {{ .MaxLength }}",
}

// ruleErrorTemplateVars lists all the possible variables that can be used builtin rules' message templates.
type ruleErrorTemplateVars[T any] struct {
	PropertyValue T
	Error         string
	MinLength     int
	MaxLength     int
}

func mustExecuteTemplate[T any](tpl *template.Template, vars ruleErrorTemplateVars[T]) string {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, vars); err != nil {
		slog.Error("failed to execute message template",
			slog.String("template", tpl.Name()),
			slog.String("error", err.Error()))
	}
	return buf.String()
}

var messageTemplatesCache = messageTemplatesMap{
	tmpl: make(map[templateKey]*template.Template),
	mu:   sync.RWMutex{},
}

type messageTemplatesMap struct {
	tmpl map[templateKey]*template.Template
	mu   sync.RWMutex
}

func (p *messageTemplatesMap) Lookup(key templateKey) *template.Template {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.tmpl[key]
}

func (p *messageTemplatesMap) Register(key templateKey, tpl *template.Template) {
	p.mu.Lock()
	p.tmpl[key] = tpl
	p.mu.Unlock()
}

func getMessageTemplate(key templateKey) *template.Template {
	if tpl := messageTemplatesCache.Lookup(key); tpl != nil {
		return tpl
	}
	text, ok := rawMessageTemplates[key]
	if !ok {
		panic(fmt.Sprintf("message template %q was not found", key))
	}
	tpl := template.Must(template.New("").Parse(text))
	messageTemplatesCache.Register(key, tpl)
	return tpl
}
