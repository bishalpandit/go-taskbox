package cache

import (
	"errors"

	"github.com/bishalpandit/taskbox/constants"
)



/*
Values of key: "all", "sort-time-asc", "sort-time-desc"
*/

func SetTasks(key string, tasks []constants.Task) {
	_, ok := TasksCache[key]
	if !ok {
		TasksCache[key] = tasks
	}
}

func GetTasks(key string) ([]constants.Task, error) {
	tasks, ok := TasksCache[key]
	if ok {
		return tasks, nil
	} else {
		return make([]constants.Task, 0), errors.New("no cache")
	}
}