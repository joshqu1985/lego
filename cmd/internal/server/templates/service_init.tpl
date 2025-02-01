package service

import (
  // "{{$.ServerName}}/internal/repo"
)

type Service struct {
  {{$.Name}} I{{$.Name}}
}

func New() *Service {
	s := &Service{
		{{$.Name}}:  New{{$.Name}}(),
  }
	return s
}
