package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeResponse(w http.ResponseWriter, resp interface{}, httpStatusCode int, err error) {
	if err != nil {
		log.Printf("error:%s\n", err.Error())
		resp = map[string]string{"message": apiErrorMessage(err)}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("json encode error: %s", err.Error())
	}
}

// Don't say everything about error to the api user, filter and/or convert to
// some meaningful messages when needed.
func apiErrorMessage(err error) string {
	switch err {
	case nil:
		return "Status ok"
	case someError:
		return err.Error()
	default:
		return "Internal error"
	}
}
