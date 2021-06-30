package main

import (
	"testing"
	"net/http/httptest"
	"encoding/json"
)

func TestResultAdd(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/add?a=1&b=3", nil)
	w := httptest.NewRecorder()
	AddHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 4 {
		t.Errorf("Invalid value  %d expected %d", res.Value, 4)
	 }
}
func TestResultAddTwo(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/add?a=-100&b=3", nil)
	w := httptest.NewRecorder()
	AddHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != -97 {
		t.Errorf("Invalid value %d expected %d", res.Value, -97)
	 }
}
func TestResultAddTree(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/add?a=f00&b=3", nil)
	w := httptest.NewRecorder()
	AddHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.ErrCode == "" || res.Success == true {
		t.Errorf("Invalid ErrCode %+v or Success %+v", res.ErrCode, res.Success)
	 }
}
func TestResultSub(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/sub?a=6&b=3", nil)
	w := httptest.NewRecorder()
	SubHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 3 {
		t.Errorf("Invalid value  %d expected %d", res.Value, 3)
	 }
}
func TestResultSubTwo(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/sub?a=100&b=-6", nil)
	w := httptest.NewRecorder()
	SubHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 106 {
		t.Errorf("Invalid value %d expected %d", res.Value, 106)
	 }
}
func TestResultSubTree(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/sub?a=100", nil)
	w := httptest.NewRecorder()
	SubHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.ErrCode == "" || res.Success == true {
		t.Errorf("Invalid ErrCode %+v or Success %+v", res.ErrCode, res.Success)
	 }
}
func TestResultMul(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/mul?a=2&b=3", nil)
	w := httptest.NewRecorder()
	MulHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 6 {
		t.Errorf("Invalid value  %d expected %d", res.Value, 6)
	 }
}
func TestResultMulTwo(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/mul?a=100&b=-6", nil)
	w := httptest.NewRecorder()
	MulHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != -600 {
		t.Errorf("Invalid value %d expected %d", res.Value, -600)
	 }
}
func TestResultMulTree(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/mul?a=100&a=30&b=50", nil)
	w := httptest.NewRecorder()
	MulHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 5000 {
		t.Errorf("Invalid value %d expected %d", res.Value, 5000)
	 }
}
func TestResultDiv(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/div?a=6&b=3", nil)
	w := httptest.NewRecorder()
	DivHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 2 {
		t.Errorf("Invalid value  %d expected %d", res.Value, 2)
	 }
}
func TestResultDivTwo(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/div?a=100&b=-50", nil)
	w := httptest.NewRecorder()
	DivHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != -2 {
		t.Errorf("Invalid value %d expected %d", res.Value, -2)
	 }
}
func TestResultDivTree(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/div?a=100&b=0", nil)
	w := httptest.NewRecorder()
	DivHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.ErrCode == "" || res.Success == true {
		t.Errorf("Invalid ErrCode %+v or Success %+v", res.ErrCode, res.Success)
	 }
}
func TestResultDivFour(t *testing.T) {
	res := Result{}
	r := httptest.NewRequest("GET", "http://localhost:8080/api/div?a=1&b=3", nil)
	w := httptest.NewRecorder()
	DivHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.ErrCode == "" || res.Success == true {
		t.Errorf("Invalid ErrCode %+v or Success %+v", res.ErrCode, res.Success)
	 }
}