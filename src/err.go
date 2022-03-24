package grpt

import (
	"errors"
	"fmt"
	"net/http"
)

// Executing error cases to html
func ErrPrint(w http.ResponseWriter, errMsg *Error) {
	w.WriteHeader(errMsg.Code)
	if err := RenderTemplate(w, "err", errMsg); err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func MethodChecker(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}
	return nil
}
