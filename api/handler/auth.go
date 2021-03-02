package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"merchant/model"
)

var (
	jwtSecret   = "pacenow-token"
	jwtLifetime = "24h"
)

// RegisterMerchant godoc
// @Summary Register new merchant
// @Description Register new merchant
// @tags auth
//
// @Router /auth/register [POST]
// @Accept json
// @Param body body model.RegistrationForm true "Register merchant"
//
// @Success 201 {string} string
// @Failure 422 {object} validator.ErrResponse
// @Failure 409 {object} model.SrvError
// @Failure 500 {object} model.SrvError
func (srv *Server) HandleRegister(w http.ResponseWriter, r *http.Request) {
	form := &model.RegistrationForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrFormDecodingFailure)
		return
	}
	if srv.handleValidationErrors(w, form) {
		return
	}

	if form.Password != form.ConfirmPassword {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, valErrPasswordMatchFailure)
		return
	}

	merchant, err := srv.DB.ReadMerchantByEmail(form.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrHashGenerationFailure)
		return
	}

	if merchant != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataDuplicateInsertion)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrHashGenerationFailure)
		return
	}

	merchant = form.ToMerchantModel()
	merchant.Password = string(pass)

	if err = srv.DB.CreateMerchant(merchant); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataCreationFailure)
		return
	}

	srv.Logger.Info(fmt.Sprintf("New User created: %s", merchant.Email))
	w.WriteHeader(http.StatusCreated)
}

// LoginMerchant godoc
// @Summary Login
// @Description Login
// @tags auth
//
// @Router /auth/login [POST]
// @Produce  json
// @Param body body model.LoginForm true "Login merchant"
//
// @Success 200 {string} string
// @Failure 500 {object} model.SrvError
func (srv *Server) HandleLogin(w http.ResponseWriter, r *http.Request) {
	form := &model.LoginForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrFormDecodingFailure)
		return
	}
	if srv.handleValidationErrors(w, form) {
		return
	}

	user, err := srv.DB.ReadMerchantByEmail(form.Email)
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

	duration, err := time.ParseDuration(jwtLifetime)
	if err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, err.Error())
		return
	}

	expiresAt := time.Now().Add(duration).Unix()

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrAuthenticationFailure)
		return
	}

	tk := &model.Token{
		UserId: user.ID,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrJsonCreationFailure)
		return
	}

	var resp = map[string]interface{}{
		"token": tokenString,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrJsonCreationFailure)
		return
	}
}
