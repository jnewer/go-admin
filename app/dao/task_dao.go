package dao

import (
	"github.com/cilidm/toolbox/gconv"
	"github.com/jinzhu/gorm"
	"pear-admin-go/app/global"
	"pear-admin-go/app/model"
	"strings"
)

type TaskDao interface {
	Insert(model.Backup) error
	FindOne(int) (*model.Backup, error)
	Update(model.Backup, map[string]interface{}) error
	Delete(model.Backup) error
	Findtask(k, v string) (*model.Backup, error)
	Findtasks(k, v string) ([]model.Backup, error)
	FindByPage(pageNum, limit int, filters ...interface{}) ([]model.Backup, int, error)
}

func NewTaskDaoImpl() TaskDao {
	s := new(TaskDaoImpl)
	return s
}

type TaskDaoImpl struct {
}

func (t TaskDaoImpl) Insert(task model.Backup) error {
	err := global.DBConn.Create(&task).Error
	return err
}

func (t TaskDaoImpl) FindOne(id int) (*model.Backup, error) {
	var task model.Backup
	err := global.DBConn.Model(model.Backup{}).First(&task, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &task, err
}

func (t TaskDaoImpl) Update(task model.Backup, m map[string]interface{}) error {
	err := global.DBConn.Model(&task).Omit("id").Updates(m).Error
	return err
}

func (t TaskDaoImpl) Delete(task model.Backup) error {
	err := global.DBConn.Delete(&task).Error
	return err
}

func (t TaskDaoImpl) Findtask(k, v string) (*model.Backup, error) {
	var task model.Backup
	err := global.DBConn.Model(model.TaskServer{}).Where(k, v).First(&task).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (t TaskDaoImpl) Findtasks(k, v string) ([]model.Backup, error) {
	var tasks []model.Backup
	err := global.DBConn.Model(model.TaskServer{}).Where(k, v).Find(&tasks).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return tasks, nil
}

func (t TaskDaoImpl) FindByPage(pageNum, limit int, filters ...interface{}) ([]model.Backup, int, error) {
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
	var (
		tasks []model.Backup
		count int
	)
	query := global.DBConn.Model(model.TaskServer{})
	query.Where(strings.Join(queryArr, " AND "), values...).Count(&count)
	err := query.Where(strings.Join(queryArr, " AND "), values...).Order("id desc").Limit(limit).Offset(offset).Find(&tasks).Error
	return tasks, count, err
}
