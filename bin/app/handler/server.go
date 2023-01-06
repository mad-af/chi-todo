package handler

import (
	"net/http"
	activity "todo/bin/app/controller/activity"
	todo "todo/bin/app/controller/todo"
	rep "todo/bin/repositories"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Server struct {
	Chi *chi.Mux
}

type Service struct {
	Router chi.Router
	DB     *gorm.DB
}

func (s *Server) Routes() {
	var r = s.Chi
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello."))
	})

	r.Route("/activity-groups", func(r chi.Router) {
		var service = Service{Router: r, DB: rep.Init(rep.Mysql())}
		activity.Init((*activity.Service)(&service))
	})

	r.Route("/todo-items", func(r chi.Router) {
		var service = Service{Router: r, DB: rep.Init(rep.Mysql())}
		todo.Init((*todo.Service)(&service))
	})

}
