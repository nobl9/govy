package messagetemplates

import (
	"fmt"
	"sync"
	"text/template"
)

// Get returns a message template by its key.
// If the template is not found, it panics.
// The first time a template is requested, it is parsed and stored in the cache.
// Custom functions and dependencies are added to the template automatically.
func Get(key templateKey) *template.Template {
	if tpl := messageTemplatesCache.Lookup(key); tpl != nil {
		return tpl
	}

	text, ok := rawMessageTemplates[key]
	if !ok {
		panic(fmt.Sprintf("message template %q was not found", key))
	}
	text += commonTemplateSuffix

	tpl := newTemplate(key, text)

	messageTemplatesCache.Register(key, tpl)
	return tpl
}

// messageTemplatesCache is a cache for message templates
// which can be modified and accessed concurrently.
var messageTemplatesCache = newMessageTemplatesMap()

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

func newTemplate(key templateKey, text string) *template.Template {
	tpl := template.New(key.String())
	tpl = AddFunctions(tpl)
	tpl = template.Must(tpl.Parse(text))

	if deps, ok := templateDependencies[key]; ok {
		addTemplatesToTemplateTree(tpl, deps...)
	}
	return tpl
}

// addTemplatesToTemplateTree adds message templates to a provided message template parse tree.
// This way they are available for use in the template.
// Example:
//
//	{{ template "TemplateName" . }}
func addTemplatesToTemplateTree(tpl *template.Template, keys ...templateKey) {
	for _, key := range keys {
		text, ok := rawMessageTemplates[key]
		if !ok {
			panic(fmt.Sprintf("dependency template %q was not found", key))
		}
		depTpl := newTemplate(key, text)
		if depTpl.Name() == "" {
			panic(fmt.Sprintf("dependency template %q has no name", key))
		}
		if _, err := tpl.AddParseTree(depTpl.Name(), depTpl.Tree); err != nil {
			panic(fmt.Sprintf("failed to add message template %q as a dependency for %q: %v",
				depTpl.Name(), tpl.Name(), err))
		}
	}
}

func newMessageTemplatesMap() messageTemplatesMap {
	return messageTemplatesMap{
		tmpl: make(map[templateKey]*template.Template),
		mu:   sync.RWMutex{},
	}
}
