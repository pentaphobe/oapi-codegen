package codegen

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

var (
	_template_singleton_do_not_use_directly_ *template.Template
)

func GetTemplateSingleton() *template.Template {
	if _template_singleton_do_not_use_directly_ == nil {
		// SMELL: this is not how to do this properly
		//        Though arguably, singletons are an anti-pattern anyway.
		//        but beats updating every function to pass around a template instance.

		// Normally a singleton would be self-initialising
		// but since there's a lot of logic and setup involved, we're leaving it to
		// codegen.go to set this up.
		panic("Template singleton not initialised")
	}
	return _template_singleton_do_not_use_directly_
}

func SetTemplateSingleton(t *template.Template) {
	if _template_singleton_do_not_use_directly_ != nil {
		panic("BUG: Template singleton initialised more than once")
	}
	_template_singleton_do_not_use_directly_ = t
}

func ExecuteTemplate(w io.Writer, templateName string, data any) error {
	err := GetTemplateSingleton().ExecuteTemplate(w, templateName, data)
	if err != nil {
		return fmt.Errorf("error executing template %s: %w args:%#v", templateName, err, data)
	}
	return nil
}

func ExecuteTemplateToString(templateName string, data any) (string, error) {
	var buf strings.Builder
	err := ExecuteTemplate(&buf, templateName, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Must is a helper to wrap Go tuple style function returns such that they
// just return the value.  Panics on error.
//
// Example:
//
//	var s String = Must(ExecuteTemplateToString("mytemplate", mydata))
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
