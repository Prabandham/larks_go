package interfaces

import (
	"github.com/gin-gonic/gin"

	"github.com/Prabandham/larks_go/managers"
)

type Interface struct{}

// Gets the list of all projects
// Accepted params
// page=1
// limit=10
// sort=
// sort_direction
func (int Interface) ListProjects(c *gin.Context) {
	c.JSON(200, managers.ListProjects())
}

func (i Interface) CreateProject(c *gin.Context) {
	c.String(200, "TODO")
}

func (i Interface) ShowProject(c *gin.Context) {
	id := c.Param("id")
	obj, err := managers.GetObjectByID(id)
	if err != nil {
		c.JSON(200, gin.H{
			"data": nil,
			"err":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": obj,
		"err":  nil,
	})
}

func (i Interface) UpdateProject(c *gin.Context) {
	c.String(200, "TODO")
}

func (i Interface) DeleteProject(c *gin.Context) {
	c.String(200, "TODO")
}
