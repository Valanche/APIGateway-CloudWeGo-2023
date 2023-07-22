package dbdata

import (
	serverz "day3/kxS/kitex_gen/kitex/serverZ"
	"time"
)

type College struct {
	Name      string `gorm:"primarykey"`
	Address   string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func NewCollege(dto *serverz.Student) (res College) {
	res.Name = dto.College.Name
	res.Address = dto.College.Address
	return
}
