package provider

import "github.com/oooiik/test_09.03.2024/internal/app"

type Interface interface {
	Bind(app.Interface)
	Boot()     // for init methods
	Register() // register: interface bind types
}

var List = []Interface{
	&http{},
}
