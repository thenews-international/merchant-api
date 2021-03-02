package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"merchant/util/validator"
)

var (
	valErrPasswordMatchFailure = "Passwords didn't match. Try again."
)

func (srv Server) handleValidationErrors(w http.ResponseWriter, form interface{}) bool {
	if err := srv.Validator.Struct(form); err != nil {
		//srv.Logger.Warn().Err(err).Msg("")

		resp := validator.ToErrResponse(err)
		if resp == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "%v"}`, "srvErrFormErrResponseFailure")
			return true
		}

		respBody, err := json.Marshal(resp)
		if err != nil {
			//srv.Logger.Warn().Err(err).Msg("")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "%v"}`, "srvErrJsonCreationFailure")
			return true
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write(respBody)
		return true
	}

	return false
}
