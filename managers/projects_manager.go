package managers

import (
	"github.com/Prabandham/larks_go/db"
	. "github.com/Prabandham/larks_go/objects"
)

func ListProjects() *[]Project {
	db := db.GetConnection()
	projects := []Project{}
	db.Connection.Find(&projects)
	return &projects
}
