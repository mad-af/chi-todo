package todo

import (
	"fmt"
	"net/http"
	"strconv"
	res "todo/bin/packages/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

var db *gorm.DB

type Service struct {
	Router chi.Router
	DB     *gorm.DB
}

func Init(s *Service) {
	db = s.DB
	s.Router.Get("/", s.GetList)
	s.Router.Get("/{id}", s.GetDetail)
	s.Router.Post("/", s.Create)
	s.Router.Patch("/{id}", s.Update)
	s.Router.Delete("/{id}", s.Delete)
}

func (s *Service) GetList(w http.ResponseWriter, r *http.Request) {
	var reply = &res.Response{Code: http.StatusOK, Status: "Success", Message: "Success", Data: make([]Todos, 0)}

	var query = r.URL.Query().Get("activity_group_id") 
	var activityGroupID, err = strconv.Atoi(query)

	if query == "" {
		GetList(nil, reply)
	} else if err == nil {
		GetList(&activityGroupID, reply)
	}
	
	render.Status(r, reply.Code)
	render.Respond(w, r, reply)
}

func (s *Service) GetDetail(w http.ResponseWriter, r *http.Request) {
	var reply = &res.Response{Code: http.StatusOK, Status: "Success", Message: "Success"}

	var param = chi.URLParam(r, "id")
	if id, err := strconv.Atoi(param); err != nil {
		res.ReplyError(http.StatusNotFound, "Not Found", fmt.Sprintf("Todo with ID %s Not Found", param), reply)
	} else {
		GetDetail(id, reply)
	}

	render.Status(r, reply.Code)
	render.Respond(w, r, reply)
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	var reply = &res.Response{Code: http.StatusCreated, Status: "Success", Message: "Success"}

	var data = &Todos{}
	if err := render.Bind(r, data); err != nil {
		res.ReplyError(http.StatusBadRequest, "Bad Request", err.Error(), reply)
	} else {
		Create(data, reply)
	}

	render.Status(r, reply.Code)
	render.Respond(w, r, reply)
}

func (s *Service) Update(w http.ResponseWriter, r *http.Request) {
	var reply = &res.Response{Code: http.StatusOK, Status: "Success", Message: "Success"}

	var param = chi.URLParam(r, "id")
	var data = &Todos{}
	if id, err := strconv.Atoi(param); err != nil {
		res.ReplyError(http.StatusNotFound, "Not Found", fmt.Sprintf("Todo with ID %s Not Found", param), reply)
	} else if err := render.Bind(r, data); err != nil {
		res.ReplyError(http.StatusBadRequest, "Bad Request", err.Error(), reply)
	} else {
		Update(id, data, reply)
	}

	render.Status(r, reply.Code)
	render.Respond(w, r, reply)
}

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	var reply = &res.Response{Code: http.StatusOK, Status: "Success", Message: "Success", Data: make(map[string]interface{})}

	var param = chi.URLParam(r, "id")
	if id, err := strconv.Atoi(param); err != nil {
		res.ReplyError(http.StatusNotFound, "Not Found", fmt.Sprintf("Todo with ID %s Not Found", param), reply)
	} else {
		Delete(id, reply)
	}

	render.Status(r, reply.Code)
	render.Respond(w, r, reply)
}
