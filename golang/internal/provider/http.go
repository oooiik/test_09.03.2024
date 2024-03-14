package provider

import "github.com/oooiik/test_09.03.2024/internal/app"

type http struct {
	app app.Interface
}

func (p *http) Bind(app app.Interface) {
	p.app = app
}
func (p *http) Boot() {
}
func (p *http) Register() {
}
