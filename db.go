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
var technicalDatabase = []byte("technical")
var ohlcDatabase = []byte("ohlc")

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

func ReadOHLC(keyString string) string {
	var result string
	key := []byte(keyString)

	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(ohlcDatabase)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", ohlcDatabase)
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

func ReadTechnicals(keyString string) string {
	var result string
	key := []byte(keyString)

	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(technicalDatabase)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", technicalDatabase)
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

func WriteOHLC(stringkey string, jsonData []byte) []byte {
	key := []byte(stringkey)
	value := jsonData
	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(ohlcDatabase)
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

func WriteTechnicals(stringkey string, jsonData []byte) []byte {
	key := []byte(stringkey)
	value := jsonData
	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(technicalDatabase)
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

func GetKeys(dbKey string) int {
	keyCount := 0
	keyList := []string{}

	client.boltDB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(dbKey))
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", []byte(dbKey))
		}
		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			// fmt.Printf("key=%s, value=%s\n", k, v)
			// fmt.Printf("key=%s\n", k)
			keyList = append(keyList, string(k))
			keyCount++
		}

		return nil
	})
	fmt.Println(dbKey, keyCount)
	return keyCount
}
