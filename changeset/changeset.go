package changeset

import (
	"reflect"
)

type Changeset struct {
	errors  []error
	changes map[string]interface{}
	data    map[string]interface{}
	types   map[string]reflect.Type
}

func (changeset *Changeset) Errors() []error {
	return changeset.errors
}

func (changeset *Changeset) Error() error {
	if changeset.errors != nil {
		return changeset.errors[0]
	}
	return nil
}

func (changeset *Changeset) Changes() map[string]interface{} {
	return changeset.changes
}

func (changeset *Changeset) Data() map[string]interface{} {
	return changeset.data
}
