package dbApi

import (
	"github.com/boltdb/bolt"
)

const (
	DBFile     = "source/my.db"
	BucketName = "DB"
	DBPort     = 0600
)

func connect() (*bolt.DB, error) {
	dbs, err := bolt.Open(DBFile, DBPort, nil)
	if err != nil {
		return nil, err
	}
	return dbs, nil
}

func InitDB() error {
	dbs, err := connect()
	if err != nil {
		return err
	}

	dbs.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(BucketName))
		if err != nil {
			return err
		}
		b.Put([]byte("/urlshort"), []byte("https://github.com/gophercises/urlshort"))
		b.Put([]byte("/abba"), []byte("https://google.com/search?q=abba"))
		b.Put([]byte("/aha"), []byte("https://google.com/search?q=oci"))
		b.Put([]byte("/hack"), []byte("https://hackerrank.com"))
		return nil
	})
	defer dbs.Close()
	return nil
}

func Get(key string) (string, error) {
	var value string
	dbs, err := connect()
	if err != nil {
		return value, err
	}
	err = dbs.View(func(tx *bolt.Tx) error {
		value = string(tx.Bucket([]byte(BucketName)).Get([]byte(key)))
		return nil
	})
	defer dbs.Close()
	return value, err
}
