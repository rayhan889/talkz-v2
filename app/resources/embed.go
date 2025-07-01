package resources

import (
	_ "embed"
)

var (
	//go:embed templates/welcome_template.html
	WelcomeEmailTemplate string
)
