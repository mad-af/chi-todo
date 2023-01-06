package todo

import (
	"fmt"
	"net/http"
	"time"
	res "todo/bin/packages/response"
)

func GetList(activityGroupID *int, reply *res.Response) {
	var todos = make([]Todos, 0)

	var query = db.Model(&Todos{}).Where("deleted_at IS NULL")
	if activityGroupID != nil {
		query.Where(&Todos{ActivityGroupID:*activityGroupID}, "activity_group_id")
	}
	query.Find(&todos)
	reply.Data = todos
}

func GetDetail(id int, reply *res.Response) {
	var activity = Todos{ID: id}
	if q := db.First(&activity); q.RowsAffected == 0 {
		res.ReplyError(http.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id), reply)
		return
	}
	reply.Data = activity
}

func Create(payload *Todos, reply *res.Response) {
	if payload.Priority == "" {
		payload.Priority = "very-high"
	}
	db.Create(payload)
	reply.Data = payload
}

func Update(id int, payload *Todos, reply *res.Response) {
	var activity = Todos{ID: id}
	payload.UpdatedAt = time.Now()
	if q := db.Model(&activity).Where("deleted_at IS NULL").Updates(payload); q.RowsAffected == 0 {
		res.ReplyError(http.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id), reply)
		return
	}

	db.First(&activity)
	reply.Data = activity
}

func Delete(id int, reply *res.Response) {
	var activity = Todos{ID: id}
	var now = time.Now()
	if q := db.Model(&activity).Where("deleted_at IS NULL").Updates(&Todos{DeletedAt: &now}); q.RowsAffected == 0 {
		res.ReplyError(http.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id), reply)
		return
	}
}