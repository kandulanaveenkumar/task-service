package repositories

import (
	"task-service/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) GetAll(status string, page, limit int) ([]models.Task, error) {
	var tasks []models.Task
	query := r.DB.Model(&models.Task{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetByID(id uint) (models.Task, error) {
	var task models.Task
	err := r.DB.First(&task, id).Error
	return task, err
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Task{}, id).Error
}
