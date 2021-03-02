package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"

	"merchant/model"
)

// ListMerchant godoc
// @Summary List merchant
// @Description get merchant list
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.MerchantDtos
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /merchants [get]
func (srv *Server) HandleListMerchant(w http.ResponseWriter, r *http.Request) {

	merchants, err := srv.DB.ListMerchants()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataAccessFailure)
		return
	}

	if len(merchants) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	dtos := merchants.ToDto()

	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrJsonCreationFailure)
		return
	}
}

// ReadMerchant godoc
// @Summary Read merchant
// @Description get a merchant
// @Produce  json
// @Param id path string true "Merchant ID"
// @Success 200 {object} model.Merchant
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /merchants/{id} [get]
func (srv *Server) HandleReadMerchant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	merchant, err := srv.DB.ReadMerchantById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataAccessFailure)
		return
	}

	dto := merchant.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrJsonCreationFailure)
		return
	}
}

// UpdateMerchant godoc
// @Summary Update merchant
// @Description update merchant
// @Param id path string true "Merchant ID"
// @Success 200 {string} string	"accepted"
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /merchants/{id} [put]
func (srv *Server) HandleUpdateMerchant(w http.ResponseWriter, r *http.Request) {
	form := &model.MerchantForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrFormDecodingFailure)
		return
	}
	if srv.handleValidationErrors(w, form) {
		return
	}

	id := chi.URLParam(r, "id")

	if _, err := srv.DB.ReadMerchantById(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataAccessFailure)
		return
	}

	if err := srv.DB.UpdateMerchantDescriptionById(id, form.Description); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataUpdateFailure)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// DeleteMerchant godoc
// @Summary Delete merchant
// @Description delete a merchant
// @Param id path string true "Merchant ID"
// @Success 200 {string} string	"ok"
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /merchants/{id} [delete]
func (srv *Server) HandleDeleteMerchant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if _, err := srv.DB.ReadMerchantById(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataAccessFailure)
		return
	}

	if err := srv.DB.DeleteMerchant(id); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataDeleteFailure)
		return
	}

}
