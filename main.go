package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	"log"
)

type Result struct{
    Success bool
    ErrCode string 
    Value int  
}

const Num1None = "Number #1 nonexist"
const Num2None = "Number #2 nonexist"
const NumZero = "Number #2 is zero"
const FloatRes = "result is float"

func Calculate(nOne, nTwo int, oper string) (res Result){
	switch oper{
	case "+":
		res.Value = nOne + nTwo
	case "-":
		res.Value = nOne - nTwo
	case "*":
		res.Value = nOne * nTwo
	case "/":
		res.Value = nOne / nTwo
	default:
		res.Value = 0
		res.ErrCode = "operations doesn't exist"
		res.Success = false
		return
	}
	res.ErrCode = ""
	res.Success = true
	return res
}

func CheckNum(vals map[string][]string) (nOne,nTwo int,err string){
	var er error


	numOne, ok := vals["a"]
	if !ok{
	   return 0, 0, Num1None
	} 
	numTwo, ok := vals["b"]
	if !ok{
	   return 0, 0, Num2None
	} 
	nOne, er = strconv.Atoi(numOne[0])
	if er != nil{
	   return 0, 0, er.Error()
	}
	nTwo, er = strconv.Atoi(numTwo[0])
	if er != nil{
	   return 0, 0, er.Error()
	}
	return nOne, nTwo, ""
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var res Result
	vals := r.URL.Query()

	nOne, nTwo, err := CheckNum(vals)
	
	res.ErrCode = err
	if res.ErrCode == ""{
	   res = Calculate(nOne,nTwo,"+")
	}else{
	   res.Success = false
	}

	json_data, er := json.Marshal(res)
	if er != nil {     
           http.Error(w, er.Error(), 404)
           return
        }
        w.Write(json_data)  

}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	var res Result
	vals := r.URL.Query()

	nOne, nTwo, err:= CheckNum(vals)

	res.ErrCode = err
	if res.ErrCode == ""{
	   res = Calculate(nOne,nTwo,"-")
	}else{
	   res.Success = false
	}
	json_data, er := json.Marshal(res)
	if er != nil {     
           http.Error(w, er.Error(), 404)
           return
        }
        w.Write(json_data) 
}

func MulHandler(w http.ResponseWriter, r *http.Request) {
	var res Result
	vals := r.URL.Query()

	nOne, nTwo, err:= CheckNum(vals)

	res.ErrCode = err
	if res.ErrCode == ""{
	    res = Calculate(nOne,nTwo,"*")
	}else{
	    res.Success = false
	}

	json_data, er := json.Marshal(res)
	if er != nil {     
           http.Error(w, er.Error(), 404)
           return
        }
        w.Write(json_data) 
}

func DivHandler(w http.ResponseWriter, r *http.Request) {
	var res Result
	vals := r.URL.Query()

	nOne, nTwo, err:= CheckNum(vals)

	res.ErrCode = err

	switch{
	case nTwo == 0:
	   res.ErrCode = NumZero
	case nOne * nOne < nTwo * nTwo:
	   res.ErrCode = FloatRes
	}

	if res.ErrCode == "" {
 	   res = Calculate(nOne,nTwo,"/")
	}else{
	   res.Success = false
	}
	json_data, er := json.Marshal(res)
	if er != nil {     
           http.Error(w, er.Error(), 404)
           return
        }
    w.Write(json_data) 
}

func main(){
	fmt.Println("")
	r := mux.NewRouter()
	sub := r.PathPrefix("/api/").Subrouter()
        sub.HandleFunc("/add", AddHandler).Methods("GET")
	sub.HandleFunc("/sub", SubHandler).Methods("GET")
        sub.HandleFunc("/mul", MulHandler).Methods("GET")
	sub.HandleFunc("/div", DivHandler).Methods("GET")
        http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
