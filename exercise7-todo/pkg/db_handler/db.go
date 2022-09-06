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

		err := b.Put(util.Itob(int(id)), []byte(task))
		util.IsErrorOccured(err)
		return nil
	})
}

func DeleteTask(db *bolt.DB, task string) {
	to_del, err := strconv.Atoi(task)
	util.IsErrorOccured(err)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		err := b.Delete(util.Itob(to_del))
		util.IsErrorOccured(err)
		return nil
	})
}

func PrintTasks(db *bolt.DB) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("Task %d: %s\n", util.Btoi(k), v)
		}
		return nil
	})
}
