package db

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

const port = 0600

var taskBucket = []byte("tasks")
var db *bolt.DB

// Task -
type Task struct {
	Key   int
	Value string
	Done  bool
	TS    int64
}

// CreateTask -
func CreateTask(taskStr string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		task := Task{Key: int(id64), Value: taskStr}
		enc, err := json.Marshal(task)
		if err != nil {
			return err
		}
		return b.Put(itob(task.Key), enc)
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// AllTasks -
func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		var t Task
		for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &t)
			tasks = append(tasks, t)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, err
}

// ToDoTasks -
func ToDoTasks() ([]Task, error) {
	var tasks []Task
	ts, err := AllTasks()
	if err != nil {
		return tasks, err
	}
	for _, t := range ts {
		if !t.Done {
			tasks = append(tasks, t)
		}
	}
	return tasks, nil
}

// DoneTasks -
func DoneTasks() ([]Task, error) {
	var tasks []Task
	ts, err := AllTasks()
	if err != nil {
		return tasks, err
	}
	for _, t := range ts {
		if t.Done && dateEqual(time.Unix(t.TS, 0), time.Now()) {
			tasks = append(tasks, t)
		}
	}
	return tasks, nil
}

func deleteTask(taskID int, b *bolt.Bucket) error {
	var t Task
	id := itob(taskID)
	if err := json.Unmarshal(b.Get(id), &t); err != nil {
		return err
	}
	if !t.Done {
		return b.Delete(id)
	}
	return nil
}

// DeleteTasks -
func DeleteTasks(idxs []int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		ts, _ := ToDoTasks()
		for _, idx := range idxs {
			taskID, err := getTaskID(idx, ts)
			if err != nil {
				return err
			}
			if err = deleteTask(taskID, b); err != nil {
				return err
			}
		}
		return nil
	})
}

func doTask(taskID int, b *bolt.Bucket, tx *bolt.Tx) error {
	var t Task
	id := itob(taskID)
	if err := json.Unmarshal(b.Get(id), &t); err != nil {
		return err
	}
	t.Done = true
	t.TS = time.Now().Unix()
	enc, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return tx.Bucket(taskBucket).Put(id, enc)
}

// CompleteTasks -
func CompleteTasks(idxs []int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		ts, _ := ToDoTasks()
		for _, idx := range idxs {
			taskID, err := getTaskID(idx, ts)
			if err != nil {
				return err
			}
			if err = doTask(taskID, b, tx); err != nil {
				return err
			}
		}
		return nil
	})
}

// Init -
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, port, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}
