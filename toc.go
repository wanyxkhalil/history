package main

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
	"net/url"
)

type I18nIDs struct {
	values map[string]bool
}

func newIDs() *I18nIDs {
	return &I18nIDs{
		values: map[string]bool{},
	}
}

// Generate generates a new element id.
func (s *I18nIDs) Generate(value []byte, kind ast.NodeKind) []byte {
	value = util.TrimLeftSpace(value)
	value = util.TrimRightSpace(value)
	result := []byte(url.QueryEscape(string(value)))
	if _, ok := s.values[util.BytesToReadOnlyString(result)]; !ok {
		s.values[util.BytesToReadOnlyString(result)] = true
		return result
	}
	for i := 1; ; i++ {
		newResult := fmt.Sprintf("%s-%d", result, i)
		if _, ok := s.values[newResult]; !ok {
			s.values[newResult] = true
			return []byte(newResult)
		}
	}
}

// Put puts a given element id to the used ids table.
func (s *I18nIDs) Put(value []byte) {
	s.values[util.BytesToReadOnlyString(value)] = true
}
