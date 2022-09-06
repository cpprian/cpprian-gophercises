package db_handler

import (
	"time"

	"github.com/boltdb/bolt"
	util "github.com/cpprian/cpprian-gophercises/exercise7-todo/pkg/utility"
)

type Task struct {
	Id          int    `json:"id"`
	Task        string `json:"task"`
	Datatime    string `json:"datatime"`
	IsCompleted bool   `json:"is_completed"`
}

type TaskList []Task

func InitStruct() *TaskList {
	return &TaskList{}
}

func GenerateTask(id int, task string) ([]byte, error) {
	today := time.Now()

	newTask := Task{
		Id:          int(id),
		Task:        task,
		Datatime:    today.Format("2006-01-02"), // YYYY-MM-DD
		IsCompleted: false,
	}

	return util.MarshalJSON(newTask)
}

func ReadTask(task []byte) Task {
	var data Task
	util.UnmarshalJSON(task, &data)
	return data
}

func (t *Task) GetTaskName() string {
	return t.Task
}

func (t *Task) GetIndex() int {
	return t.Id
}

func (t *Task) MarkAsCompleted() {
	t.IsCompleted = true
}

func IsUncompletedTask(t Task) bool {
	return !t.IsCompleted
}

func IsCompletedTaskToday(t Task) bool {
	return t.IsCompleted && util.IsToday(t.Datatime)
}

func (tl *TaskList) FillTaskList(db *bolt.DB) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := ReadTask(v)
			*tl = append(*tl, task)
		}
		return nil
	})
}

func (tl *TaskList) OrderList() {
	for i := 0; i < len(*tl); i++ {
		(*tl)[i].Id = i + 1
	}
}
