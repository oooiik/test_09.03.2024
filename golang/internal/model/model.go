package model

type Interface interface {
}

type Scan interface {
	Scan(dest ...any) error
}
