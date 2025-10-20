package model

import (
	"ktfs/config"

	"time"

	softDel "gorm.io/plugin/soft_delete"
	"github.com/google/uuid"
)

type Student struct {
	ID uuid.UUID `gorm:"primaryKey; column:id; type:uuid; default:gen_random_uuid();" json:"id"`
	StudentID string `gorm:"unique; column:student_id; size:255; uniqueIndex:udx_student_id_user;" json:"student_id"`
	Name string `gorm:"column:name; default:null; size:255;" json:"name"`
	Gender string `gorm:"column:gender; default:null; size:255;" json:"gender"`
	Address string `gorm:"column:address; default:null; size:1000;" json:"address"`
	EntryYear uint `gorm:"column:entry_year; default:null; size:255;" json:"entry_year"`

	CreatedAt time.Time `gorm:"autoCreateTime; column:created_at;" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime; column:updated_at;" json:"-"`
	DeletedAt	time.Time `gorm:"column:deleted_at;" json:"-"`
	IsDeleted softDel.DeletedAt `gorm:"column:is_deleted; softDelete:flag; DeletedAtField:DeletedAt; uniqueIndex:udx_email_user;" json:"-"`
}

func IsStudentExistsByID(id string) (bool, error) {
	var studentFound int64
	err := config.DB.Model(Student{}).Where("id  = ?", id).Count(&studentFound).Error
	if err != nil {
		return false, err
	}
	if studentFound == 0 {
		return false, nil
	} else {
		return true, nil
	}
}