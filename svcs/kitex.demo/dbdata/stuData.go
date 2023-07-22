package dbdata

import (
	serverz "day3/kxS/kitex_gen/kitex/serverZ"
	"strings"
	"time"
)

type Student struct {
	Id          int32 `gorm:"primarykey"`
	Name        string
	CollegeName string
	Emails      string
	Sex         string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func NewStudent(dto *serverz.Student) (res Student) {
	res.Id = dto.Id
	res.Name = dto.Name
	res.CollegeName = dto.College.Name
	res.Emails = strings.Join(dto.Email, ",")
	res.Sex = dto.Sex
	return
}
