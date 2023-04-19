package test_function

import (
	"encoding/json"
	"github.com/OpenFunction/functions-framework-go/functions"
	"net/http"
)

func init() {
	functions.HTTP("TestFunction", TestFunction,
		functions.WithFunctionMethods("POST"))
}

func TestFunction(w http.ResponseWriter, r *http.Request) {
	responseBytes, _ := json.Marshal(map[string]string{
		"test": "test",
	})
	w.Header().Set("Content-type", "application/json")
	w.Write(responseBytes)
}
