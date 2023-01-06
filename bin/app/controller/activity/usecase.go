package activity

import (
	"fmt"
	"net/http"
	"time"
	res "todo/bin/packages/response"
)

func GetList(reply *res.Response) {
	var activities = make([]Activities, 0)
	db.Model(&Activities{}).Where("deleted_at IS NULL").Find(&activities)
	reply.Data = activities
}

func GetDetail(id int, reply *res.Response) {
	var activity = Activities{ID: id}
	if q := db.First(&activity); q.RowsAffected == 0 {
		res.ReplyError(http.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id), reply)
		return
	}
	reply.Data = activity
}

func Create(payload *Activities, reply *res.Response) {
	db.Create(payload)
	reply.Data = payload
}

func Update(id int, payload *Activities, reply *res.Response) {
	var activity = Activities{ID: id}
	payload.UpdatedAt = time.Now()
	if q := db.Model(&activity).Where("deleted_at IS NULL").Updates(payload); q.RowsAffected == 0 {
		res.ReplyError(http.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id), reply)
		return
	}

	db.First(&activity)
	reply.Data = activity
}

func Delete(id int, reply *res.Response) {
	var activity = Activities{ID: id}
	var now = time.Now()
	if q := db.Model(&activity).Where("deleted_at IS NULL").Updates(&Activities{DeletedAt: &now}); q.RowsAffected == 0 {
		res.ReplyError(http.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id), reply)
		return
	}
}