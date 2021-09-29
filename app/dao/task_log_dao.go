package dao

import (
	"github.com/cilidm/toolbox/gconv"
	"pear-admin-go/app/core/db"
	"pear-admin-go/app/model"
	"strings"
)

type TaskLogDao interface {
	Insert(model.TaskLog) error
	Update(model.TaskLog, map[string]interface{}) error
	Delete(int) error
	FindByPage(pageNum, limit int, filters ...interface{}) ([]model.TaskLog, int, error)
}

func NewTaskLogDaoImpl() TaskLogDao {
	return &TaskLogDaoImpl{}
}

type TaskLogDaoImpl struct {
}

func (this *TaskLogDaoImpl) Insert(data model.TaskLog) error {
	err := db.Instance().Create(&data).Error
	return err
}

func (this *TaskLogDaoImpl) Delete(id int) error {
	err := db.Instance().Where("id = ?", id).Delete(model.TaskLog{}).Error
	return err
}

func (this *TaskLogDaoImpl) Update(info model.TaskLog, updateAttrMap map[string]interface{}) error {
	err := db.Instance().Model(&info).Where("id = ?", info.Id).Updates(updateAttrMap).Error
	return err
}

func (this *TaskLogDaoImpl) FindByPage(pageNum, limit int, filters ...interface{}) (infos []model.TaskLog, count int, err error) {
	client := db.Instance()
	offset := (pageNum - 1) * limit

	var queryArr []string
	var values []interface{}
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			queryArr = append(queryArr, gconv.String(filters[k]))
			values = append(values, filters[k+1])
		}
	}

	query := client.Model(model.TaskLog{})
	query.Where(strings.Join(queryArr, " AND "), values...).Count(&count)
	err = query.Where(strings.Join(queryArr, " AND "), values...).Order("id desc").Limit(limit).Offset(offset).Find(&infos).Error
	return
}
