package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
