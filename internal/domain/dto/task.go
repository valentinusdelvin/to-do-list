package dto

type CreateTask struct {
	TaskId   string `json:"task_id"`
	Title    string `json:"title" validate:"required"`
	Details  string `json:"details" validate:"required"`
	Deadline string `json:"deadline" validate:"required,datetime=2006-01-02"`
	UserId   string `json:"user_id" validate:"required"`
	Status   string `json:"status" validate:"required,oneof=pending in-progress completed"`
}

type GetAllTasks struct {
	TaskId   string `json:"task_id"`
	Title    string `json:"title"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
}

type GetTaskById struct {
	TaskId   string `json:"task_id"`
	Title    string `json:"title"`
	Details  string `json:"details"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
	UserId   string `json:"user_id"`
}

type UpdateProgress struct {
	TaskId string `json:"task_id"`
	Status string `json:"status" validate:"required,oneof=pending in-progress completed"`
}

type UpdateTask struct {
	TaskId      string `json:"task_id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"details" validate:"required"`
	Deadline    string `json:"deadline" validate:"required,datetime=2006-01-02"`
	UserId      string `json:"user_id" validate:"required"`
}
