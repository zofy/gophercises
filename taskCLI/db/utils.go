package db

import (
	"encoding/binary"
	"errors"
	"time"
)

func itob(n int) []byte {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func getTaskID(idx int, tasks []Task) (int, error) {
	for i, t := range tasks {
		if i+1 == idx {
			return t.Key, nil
		}
	}
	return -1, errors.New("Index out of range")
}

func dateEqual(a, b time.Time) bool {
	ya, ma, da := a.Date()
	yb, mb, db := b.Date()
	return ya == yb && ma == mb && da == db
}
