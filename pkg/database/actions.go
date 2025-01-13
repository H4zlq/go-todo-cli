package database

import (
	"github.com/h4zlq/cli/model"
)

func AddTask(task model.Task) {
	db := GetInstance()
	db.Create(&task)
}

func UpdateTask(id string, task model.Task) {
	db := GetInstance()
	db.Model(&task).Where("id = ?", id).Updates(&task)
}

func GetTask(id string) model.Task {
	db := GetInstance()
	var task model.Task
	db.Where("id = ?", id).First(&task)
	return task
}

func DeleteTask(id string, task model.Task) {
	db := GetInstance()
	db.Where("id = ?", id).Delete(&task)
}

func GetTasks() []model.Task {
	db := GetInstance()
	var tasks []model.Task
	db.Find(&tasks)
	return tasks
}

func GetTasksByStatus(status string) []model.Task {
	db := GetInstance()
	var tasks []model.Task
	db.Where("status = ?", status).Find(&tasks)
	return tasks
}
