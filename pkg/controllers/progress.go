package controllers

import (
	"github.com/meritlabs/achievement-engine/pkg/db/stores"
	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/pkg/db/models/progress"
	"github.com/meritlabs/achievement-engine/pkg/dto"
	"net/http"
	"github.com/meritlabs/achievement-engine/pkg/db/models"
	_ "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func GetProgress(store *stores.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*models.User)

		p, err := getProgress(store, user.ID)

		if err != nil || &p == nil {
			c.AbortWithError(http.StatusBadRequest, dto.BadRequestError{Message: "No progress found"})
			return
		}

		c.JSON(http.StatusOK, *p)
	}
}

func UpdateTask(store *stores.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var taskProgress progress.TaskProgress
		user := c.MustGet("user").(*models.User)
		err := c.BindJSON(&taskProgress)

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, dto.BadRequestError{Message: "Invalid request object"})
			return
		}

		p, err := getProgress(store, user.ID)
		var found bool
		tasks := []progress.TaskProgress{}

		if err != nil || &p == nil {
			p = &progress.Progress{
				UserID: user.ID,
			}
		} else {
			for _, t := range p.Tasks {
				if t.Slug == taskProgress.Slug {
					found = true
					t.Status = taskProgress.Status
				}

				tasks = append(tasks, t)
			}
		}

		if !found {
			tasks = append(tasks, progress.TaskProgress{
				Slug: taskProgress.Slug,
				Status: taskProgress.Status,
			})
		}

		p.Tasks = tasks
		err = store.SetProgress(user.ID, *p)

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, dto.BadRequestError{Message: "Unable to save progress"})
			return
		}

		c.JSON(http.StatusOK, *p)
	}
}

func getProgress(store *stores.Store, userId bson.ObjectId) (*progress.Progress, error) {
	p, err := store.GetProgress(userId)
	return p, err
}