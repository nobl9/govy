package messagetemplates

import (
	"fmt"
	"sync"
	"text/template"
)

// Get returns a message template by its key.
// If the template is not found, it panics.
// The first time a template is requested, it is parsed and stored in the cache.
func Get(key templateKey) *template.Template {
	if tpl := messageTemplatesCache.Lookup(key); tpl != nil {
		return tpl
	}
	text, ok := rawMessageTemplates[key]
	if !ok {
		panic(fmt.Sprintf("message template %q was not found", key))
	}
	text += commonTemplateSuffix
	tpl := template.New("")
	tpl = AddFunctions(tpl)
	tpl = template.Must(tpl.Parse(text))
	messageTemplatesCache.Register(key, tpl)
	return tpl
}

// messageTemplatesCache is a cache for message templates
// which can be modified and accessed concurrently.
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
