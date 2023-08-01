package celeritas

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Validation struct {
	Data   url.Values
	Errors map[string]string
}

func (c *Celeritas) Validator(data url.Values) *Validation {
	return &Validation{
		Data:   data,
		Errors: make(map[string]string),
	}
}

func (v *Validation) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validation) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validation) Has(field string, r *http.Request) bool {
	if x := r.Form.Get(field); strings.TrimSpace(x) != "" {
		return true
	}
	return false
}

func (v *Validation) Required(r *http.Request, fields ...string) {
	for _, field := range fields {
		if x := r.Form.Get(field); strings.TrimSpace(x) == "" {
			v.AddError(field, "This field cannot be blank")
		}
	}
}

func (v *Validation) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validation) IsEmail(field,value string) {
	if !govalidator.IsEmail(value) {
		v.AddError(field,"Invalid email address")
	}
}

func (v *Validation) IsInt(field,value string) {
	_, err := strconv.Atoi(value)
	if err != nil {
		v.AddError(field, "This field must be an integer")
	}
}


