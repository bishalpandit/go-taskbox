package cache

import (

	"github.com/bishalpandit/taskbox/constants"
)

var TasksCache = make(map[string][]constants.Task)