package db_handler

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	util "github.com/cpprian/cpprian-gophercises/exercise7-todo/pkg/utility"
)

func InitDb() (*bolt.DB, func()) {
	db, err := bolt.Open("todo.db", 0600, nil)
	util.IsErrorOccured(err)
	closeDB := func() {
		defer db.Close()
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todo"))
		util.IsErrorOccured(err)
		return nil
	})

	return db, closeDB
}

func AddTask(db *bolt.DB, tasks []string) {
	task := util.WrapTask(tasks)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		id, _ := b.NextSequence()

		newTask, err := GenerateTask(int(id), task)
		util.IsErrorOccured(err)

		err = b.Put(util.Itob(int(id)), newTask)
		util.IsErrorOccured(err)
		return nil
	})
}

func MarkTaskAsCompleted(db *bolt.DB, task string) {
	to_mark, err := strconv.Atoi(task)
	util.IsErrorOccured(err)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		data := b.Get(util.Itob(to_mark))

		task := ReadTask(data)
		task.MarkAsCompleted()

		byteTask, err := util.MarshalJSON(task)
		util.IsErrorOccured(err)

		err = b.Put(util.Itob(task.GetIndex()), byteTask)
		util.IsErrorOccured(err)
		return nil
	})
}

func RemoveTask(db *bolt.DB, task string) {
	to_del, err := strconv.Atoi(task)
	util.IsErrorOccured(err)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		err := b.Delete(util.Itob(to_del))
		util.IsErrorOccured(err)
		return nil
	})
}

func PrintTasks(db *bolt.DB, check func(Task) bool) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		c := b.Cursor()

		i := 1
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if data := ReadTask(v); check(data) {
				fmt.Printf("Task %d: %s\n", i, data.GetTaskName())
				i++
			}
		}
		return nil
	})
}

// delete out of date tasks which is completed
// order all tasks
func UpdateDb(db *bolt.DB) {
	DeleteOutdatedTasks(db)

	taskList := InitStruct()
	taskList.FillTaskList(db)
	taskList.OrderList()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))

		DeleteAllTasks(b)
		for i, task := range *taskList {
			byteTask, err := util.MarshalJSON(task)
			util.IsErrorOccured(err)

			err = b.Put(util.Itob(i), byteTask)
			util.IsErrorOccured(err)
		}
		return nil
	})
}

func DeleteOutdatedTasks(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if data := ReadTask(v); data.IsCompleted && !util.IsToday(data.Datatime) {
				b.Delete(k)
			}
		}
		return nil
	})
}

func DeleteAllTasks(b *bolt.Bucket) {
	c := b.Cursor()

	for k, _ := c.First(); k != nil; k, _ = c.Next() {
		b.Delete(k)
	}
}
