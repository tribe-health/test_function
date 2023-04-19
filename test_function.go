package test_function

import (
	"fmt"
	ofctx "github.com/OpenFunction/functions-framework-go/context"
	"github.com/OpenFunction/functions-framework-go/functions"
	"net/http"
)

func init() {
	functions.HTTP("TestFunction", TestFunction,
		functions.WithFunctionPath("/{echo}"))
}

func TestFunction(w http.ResponseWriter, r *http.Request) {
	vars := ofctx.VarsFromCtx(r.Context())
	_, err := fmt.Fprintf(w, "Echo: %s\n", vars["echo"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
