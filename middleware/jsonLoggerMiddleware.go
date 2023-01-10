package middleware

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
	"userservice/database"
	"userservice/schema"
	"userservice/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ActivityLog schema.ActivityLog

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := util.GetDurationInMillseconds(start)
		duration_s := fmt.Sprintf("%f", duration)
		entry := log.WithFields(log.Fields{
			"client_ip":  util.GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_id":    util.GetUserID(c),
			"referrer":   c.Request.Referer(),
			"request_id": c.Param("ID"),
			"body":       c.Request.Body,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}

		var input ActivityLog

		input.Client = util.GetClientIP(c)
		input.Duration = duration_s
		input.Method = c.Request.Method
		input.Status = strconv.Itoa(c.Writer.Status())
		input.Path = c.Request.RequestURI
		input.PathOp = ""
		input.UserId = util.GetUserID(c)
		input.Source = ""
		input.Application = c.Param("APPS")
		input.RequestId = c.Param("ID")
		input.Referrer = c.Request.Referer()
		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)
		input.ReqBody = string(x)
		savedLog, _ := input.save()

		log.Print(savedLog)
	}
}

func (log ActivityLog) save() (ActivityLog, error) {
	err := database.Database.Create(&log).Error
	return log, err
}
