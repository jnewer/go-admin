package dao

import (
	"github.com/cilidm/toolbox/gconv"
	"github.com/jinzhu/gorm"
	"pear-admin-go/app/global"
	"pear-admin-go/app/model"
	"strings"
)

type TaskDao interface {
	Insert(model.Task) error
	FindOne(int) (*model.Task, error)
	Update(model.Task, map[string]interface{}) error
	Delete(model.Task) error
	Findtask(k, v string) (*model.Task, error)
	Findtasks(k, v string) ([]model.Task, error)
	FindByPage(pageNum, limit int, filters ...interface{}) ([]model.Task, int, error)
}

func NewTaskDaoImpl() TaskDao {
	s := new(TaskDaoImpl)
	return s
}

type TaskDaoImpl struct {
}

func (t TaskDaoImpl) Insert(task model.Task) error {
	err := global.DBConn.Create(&task).Error
	return err
}

func (t TaskDaoImpl) FindOne(id int) (*model.Task, error) {
	var task model.Task
	err := global.DBConn.Model(model.Task{}).First(&task, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &task, err
}

func (t TaskDaoImpl) Update(task model.Task, m map[string]interface{}) error {
	err := global.DBConn.Model(&task).Omit("id").Updates(m).Error
	return err
}

func (t TaskDaoImpl) Delete(task model.Task) error {
	err := global.DBConn.Delete(&task).Error
	return err
}

func (t TaskDaoImpl) Findtask(k, v string) (*model.Task, error) {
	var task model.Task
	err := global.DBConn.Model(model.TaskServer{}).Where(k, v).First(&task).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (t TaskDaoImpl) Findtasks(k, v string) ([]model.Task, error) {
	var tasks []model.Task
	err := global.DBConn.Model(model.TaskServer{}).Where(k, v).Find(&tasks).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return tasks, nil
}

func (t TaskDaoImpl) FindByPage(pageNum, limit int, filters ...interface{}) ([]model.Task, int, error) {
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
		tasks []model.Task
		count int
	)
	query := global.DBConn.Model(model.TaskServer{})
	query.Where(strings.Join(queryArr, " AND "), values...).Count(&count)
	err := query.Where(strings.Join(queryArr, " AND "), values...).Order("id desc").Limit(limit).Offset(offset).Find(&tasks).Error
	return tasks, count, err
}
