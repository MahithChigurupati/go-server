// helper.go
package utils

import (
	"github.com/MahithChigurupati/go-server/models"
)

func IsCourseEmpty(course models.Course) bool {
	return course.CourseName == ""
}
