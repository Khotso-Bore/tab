package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{})  {
	if body,err := io.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body),x); err != nil{
			return
		}
	}
}

func Response(w *http.ResponseWriter, statusCode int, response interface{}){
	
	res,_ := json.Marshal(response)
	(*w).Header().Set("Content-Type","application/json")
	(*w).WriteHeader(statusCode)
	(*w).Write(res)
}
