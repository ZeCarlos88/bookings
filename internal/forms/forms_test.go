package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/qualquercoisa", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error(("got invalid when should be valid"))
	}

}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/qualquercoisa", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("from shows valid when required fields should be missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "qualquercoisa", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatwever", nil)
	form := New(r.PostForm)

	has := form.Has("whatwever")
	if has {
		t.Error("form shows has field when should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error(" shows form does not has field when should ")
	}

}

func TestForm_MinLenght(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLenght("x", 10)
	if form.Valid() {
		t.Error("form shows Minlenght for non existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("shouuld have an error but didnt have one")
	}

	postedData = url.Values{}
	postedData.Add("some_field", "some value")
	form = New(postedData)

	form.MinLenght("some_field", 100)
	if form.Valid() {
		t.Error("shows minlenght of 100 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abc123")

	form = New(postedData)

	form.MinLenght("another_field", 1)
	if !form.Valid() {
		t.Error("shows minlenght of 1 is not met when it is ")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("shouuld not have an error but have one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "me@here.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got invalid email when it shouldnt")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "x")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid email when it shouldnt be")
	}
}
