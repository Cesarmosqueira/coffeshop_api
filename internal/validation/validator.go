package validation

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Cesarmosqueira/coffeshop_api/internal/util"
)

type Validator struct {
	ReflectValue reflect.Value
	Fields       []string
	Errors       map[string][]string
}

type ValidationError struct {
	Field  string   `json:"field"`
	Errors []string `json:"errors"`
}

func (v *Validator) Summary() []ValidationError {
	errors := make([]ValidationError, 0)
	for field, arr := range v.Errors {
		err := ValidationError{
			Field:  util.ToSnakeCase(field),
			Errors: arr,
		}

		errors = append(errors, err)
	}

	return errors
}

func (v *Validator) CheckField(fieldName string) error {
	found := false
	for _, field := range v.Fields {
		if field == fieldName {
			found = true
			break
		}
	}

	if !found {
		message := fmt.Sprintf("Field %s not found in model.", fieldName)
		log.Println(message)
		return errors.New(message)
	}

	return nil
}

func (v *Validator) NotBlank(fieldName string) error {
	if err := v.CheckField(fieldName); err != nil {
		return err
	}

	fieldValue := v.ReflectValue.FieldByName(fieldName).String()
	if fieldValue == "" {
		message := "Must not be blank."
		v.Errors[fieldName] = append(v.Errors[fieldName], message)
	}

	return nil
}

func (v *Validator) Email(fieldName string) error {
	if err := v.CheckField(fieldName); err != nil {
		return err
	}

	fieldValue := v.ReflectValue.FieldByName(fieldName).String()
	if !util.ValidEmail(fieldValue) {
		message := "Must be a valid email."
		v.Errors[fieldName] = append(v.Errors[fieldName], message)
	}

	return nil
}

func (v *Validator) MinLength(fieldName string, length int) error {
	if err := v.CheckField(fieldName); err != nil {
		return err
	}

	fieldValue := v.ReflectValue.FieldByName(fieldName).String()
	if len(fieldValue) < length {
		message := fmt.Sprintf("Must be at least %v characters long", length)
		v.Errors[fieldName] = append(v.Errors[fieldName], message)
	}

	return nil
}

func NewValidator(model interface{}) Validator {
	v := reflect.ValueOf(model)
	f := make([]string, 0)
	for i := 0; i < v.NumField(); i++ {
		f = append(f, v.Type().Field(i).Name)
	}

	return Validator{
		ReflectValue: v,
		Fields:       f,
		Errors:       make(map[string][]string),
	}
}

