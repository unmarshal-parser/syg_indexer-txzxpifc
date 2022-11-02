// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"os"
	"sync"
)

var indexerConfig IndexerConfig

func init() {
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if len(configFilePath) == 0 {
		panic("missing CONFIG_FILE_PATH value")
	}
	err := LoadConfig("run", configFilePath, &indexerConfig)
	if err != nil {
		panic(err)
	}
	indexerConfig.AssignDefaults()
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		RunIndexer(indexerConfig)
		wg.Done()
	}()
	wg.Wait()
}
