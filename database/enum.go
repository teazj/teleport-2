package database

import (
	"github.com/pagarme/teleport/action"
	"github.com/pagarme/teleport/batcher/ddldiff"
)

type Enum struct {
	Oid  string `json:"oid"`
	Name string `json:"name"`
	Type *Type
}

func (post *Enum) Diff(other ddldiff.Diffable, context ddldiff.Context) []action.Action {
	actions := make([]action.Action, 0)

	if other == nil {
		actions = append(actions, &action.CreateEnum{
			context.Schema,
			post.Type.Name,
			post.Name,
		})
	}

	return actions
}

func (e *Enum) Children() []ddldiff.Diffable {
	return []ddldiff.Diffable{}
}

func (e *Enum) Drop(context ddldiff.Context) []action.Action {
	return []action.Action{}
}

func (e *Enum) IsEqual(other ddldiff.Diffable) bool {
	if other == nil {
		return false
	}

	if otherEnum, ok := other.(*Enum); ok {
		return (e.Oid == otherEnum.Oid)
	}

	return false
}
