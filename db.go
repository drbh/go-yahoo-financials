package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/rand"
	// "os"
	// "path/filepath"
)

var client BoltClient

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var integers = []rune("123456789")
var database = []byte("values")
var ratioDatabase = []byte("ratios")
var targetDatabase = []byte("targets")

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb(path string) {
	var err error

	bc.boltDB, err = bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) test() {
	var result []string
	bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(database)
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			result = append(result, string(k))
		}
		return nil
	})
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Read(keyString string) string {
	var result string
	key := []byte(keyString)

	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(database)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", database)
		}
		val := bucket.Get(key)
		result = string(val)
		return nil
	})

	return result
}

func ReadRatios(keyString string) string {
	var result string
	key := []byte(keyString)

	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(ratioDatabase)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", ratioDatabase)
		}
		val := bucket.Get(key)
		result = string(val)
		return nil
	})
	return result
}
func ReadTargets(keyString string) string {
	var result string
	key := []byte(keyString)

	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(ratioDatabase)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", targetDatabase)
		}
		val := bucket.Get(key)
		result = string(val)
		return nil
	})
	return result
}

func Write(stringkey string, jsonData []byte) []byte {
	key := []byte(stringkey)
	value := jsonData
	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(database)
		if err != nil {
			return err
		}
		bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	return key
}

func WriteRatios(stringkey string, jsonData []byte) []byte {
	key := []byte(stringkey)
	value := jsonData
	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(ratioDatabase)
		if err != nil {
			return err
		}
		bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	return key
}

func WriteTargets(stringkey string, jsonData []byte) []byte {
	key := []byte(stringkey)
	value := jsonData
	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(targetDatabase)
		if err != nil {
			return err
		}
		bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	return key
}
