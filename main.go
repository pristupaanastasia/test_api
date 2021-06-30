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
    Value int64  
}

const Num1None = "Number #1 nonexist"
const Num2None = "Number #2 nonexist"
const NumNone = "Numbers nonexist"

func CheckNum(vals map[string][]string) (n1,n2 int,err string){
	var e error


	num1, ok := vals["a"]
	log.Println("num2",num1)
	if !ok{
		return 0, 0, Num1None
	} 
	num2, ok := vals["b"]
	log.Println("num2",num2)
	if !ok{
		return 0, 0, Num2None
	} 
	n1, e = strconv.Atoi(num1[0])
	log.Println(n1)
	if e != nil{
		return 0, 0, e.Error()
	}
	n2, e = strconv.Atoi(num2[0])
	log.Println(n2)
	if e != nil{
		return 0, 0, e.Error()
	}
	return n1, n2, ""
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var res Result
	vals := r.URL.Query()

	n1, n2, err:= CheckNum(vals)
	
	res.ErrCode = err
	if res.ErrCode == ""{
		res.Value = n1 + n2 
		res.Success = true
	}else{
		res.Value = 0
		res.Success = false
	}
	log.Println(res.ErrCode)
	json_data, e := json.Marshal(res)
	if e != nil {     
        http.Error(w, e.Error(), 404)
        return
    }
    w.Write(json_data)  

}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	var res Result
	vals := r.URL.Query()

	n1,n2, err:= CheckNum(vals)

	res.ErrCode = err
	if res.ErrCode == ""{
		res.Value = n1 - n2 
		res.Success = true
	}else{
		res.Value = 0
		res.Success = false
	}
	json_data, e := json.Marshal(res)
	if e != nil {     
        http.Error(w, e.Error(), 404)
        return
    }
    w.Write(json_data) 
}

func MulHandler(w http.ResponseWriter, r *http.Request) {
	var res Result
	vals := r.URL.Query()

	n1,n2, err:= CheckNum(vals)

	res.ErrCode = err
	if res.ErrCode == ""{
		res.Value = n1 * n2 
		res.Success = true
	}else{
		res.Value = 0
		res.Success = false
	}
	json_data, e := json.Marshal(res)
	if e != nil {     
        http.Error(w, e.Error(), 404)
        return
    }
    w.Write(json_data) 
}

func DivHandler(w http.ResponseWriter, r *http.Request) {
	var res Result
	vals := r.URL.Query()

	n1,n2, err:= CheckNum(vals)

	res.ErrCode = err
	
	if res.ErrCode == ""{
		if n2 != 0 {
			res.Value = n1 / n2 
			res.Success = true
		}else{
			res.ErrCode = "Num2 is zero"
			res.Value = 0
			res.Success = false
		}
	}else{
		res.Value = 0
		res.Success = false
	}
	json_data, e := json.Marshal(res)
	if e != nil {     
        http.Error(w, e.Error(), 404)
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