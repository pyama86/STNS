package stns

import (
	"reflect"
	"strconv"
)

type Attribute struct {
	Id int `toml:"id" json:"id"`
	*User
	*Group
}

type Attributes map[string]*Attribute

type Linker interface {
	LinkTargetValue() []string
	LinkValue() []string
	SetLinkValue([]string)
}

func (u Attributes) GetByName(name string) Attributes {
	attr := u[name]
	if attr == nil || reflect.ValueOf(attr).IsNil() {
		return nil
	}
	return Attributes{
		name: attr,
	}
}

func (u Attributes) GetById(_id string) Attributes {
	id, _ := strconv.Atoi(_id)
	for k, u := range u {
		if u.Id == id {
			return Attributes{
				k: u,
			}
		}
	}
	return nil
}
