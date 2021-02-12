package tencentcloud

import (
	"github.com/satori/go.uuid"
	"reflect"
)

const (
	fieldClientToken = "ClientToken"
)

func injectClientToken(obj interface{}) {
	// obj Must be struct ptr
	getType := reflect.TypeOf(obj)
	if getType.Kind() != reflect.Ptr || getType.Elem().Kind() != reflect.Struct {
		return
	}

	// obj Must exist named field
	_, ok := getType.Elem().FieldByName(fieldClientToken)
	if !ok {
		return
	}

	field := reflect.ValueOf(obj).Elem().FieldByName(fieldClientToken)

	// field Must be string ptr
	if field.Kind() != reflect.Ptr {
		return
	}

	// Set if nil or empty
	if field.IsNil() || (field.Kind() == reflect.String && field.Elem().Len() == 0) {
		uuidVal := uuid.NewV4().String()
		field.Set(reflect.ValueOf(&uuidVal))
	}
}
