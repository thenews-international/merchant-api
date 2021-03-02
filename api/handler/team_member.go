package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"

	"merchant/model"
)

// ListTeamMember godoc
// @Summary List team members
// @Description get team members list
// @Produce  json
// @Success 200 {array} model.TeamMemberDtos
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /team-members [get]
func (srv *Server) HandleListTeamMember(w http.ResponseWriter, r *http.Request) {
	userDetails, _ := r.Context().Value(model.CtxKeyXUser).(model.CtxUser)
	merchantId := userDetails.UserId.String()

	members, err := srv.DB.ListTeamMembersByMerchantId(merchantId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataAccessFailure)
		return
	}

	if len(members) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	dtos := members.ToDto()

	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrJsonCreationFailure)
		return
	}
}

// CreateTeamMember godoc
// @Summary Create Team Member
// @Description add a member to merchant team
// @Accept  json
// @Param body body model.TeamMemberCreateForm true "Create a team member"
// @Success 200 {string} string "ok"
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /team-members [post]
func (srv *Server) HandleCreateTeamMember(w http.ResponseWriter, r *http.Request) {
	form := &model.TeamMemberCreateForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrFormDecodingFailure)
		return
	}
	if srv.handleValidationErrors(w, form) {
		return
	}

	userDetails, _ := r.Context().Value(model.CtxKeyXUser).(model.CtxUser)
	merchantId := userDetails.UserId.String()

	if _, err := srv.DB.ReadMerchantById(merchantId); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataAccessFailure)
		return
	}

	member := form.ToModel(merchantId)

	if err := srv.DB.CreateTeamMember(member); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataCreationFailure)
		return
	}

	srv.Logger.Info(fmt.Sprintf("New Team Member created: %s", member.Email))
	w.WriteHeader(http.StatusCreated)
}

// ReadTeamMember godoc
// @Summary Read team member
// @Description get a member
// @Produce  json
// @Param id path string true "Team Member ID"
// @Success 200 {object} model.TeamMemberDto
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /team-members/{id} [get]
func (srv *Server) HandleReadTeamMember(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	member, err := srv.DB.ReadTeamMemberById(id)
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

	dto := member.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrJsonCreationFailure)
		return
	}
}

// UpdateTeamMember godoc
// @Summary Update team member
// @Description update member
// @Produce  json
// @Param body body model.TeamMemberUpdateForm true "Update a team member"
// @Param id path string true "Team Member ID"
// @Success 200 {string} string	"accepted"
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /team-members/{id} [put]
func (srv *Server) HandleUpdateTeamMember(w http.ResponseWriter, r *http.Request) {
	form := &model.TeamMemberUpdateForm{}
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

	if _, err := srv.DB.ReadTeamMemberById(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataAccessFailure)
		return
	}

	member := form.ToModel(id)

	if err := srv.DB.UpdateTeamMemberById(id, member); err != nil {
		srv.Logger.Warn(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, srvErrDataUpdateFailure)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// DeleteTeamMember godoc
// @Summary Remove a team member from merchant
// @Description remove a member
// @Param id path string true "Team Member ID"
// @Success 200 {string} string	"ok"
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /team-members/{id} [delete]
func (srv *Server) HandleDeleteTeamMember(w http.ResponseWriter, r *http.Request) {
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
