package singleton

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once init()  -- thread safety
// laziness
var once sync.Once
var instance *singletonDatabase

// readData reads a file from the given path, parses it, and returns a map of string keys to integer values
func readData(path string) (map[string]int, error) {
	// Get the executable path
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exPath := filepath.Dir(ex)

	// Open the file
	file, err := os.Open(filepath.Join(exPath, path)) // Use filepath.Join for proper path concatenation
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Map to store the parsed data
	data := make(map[string]int)

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split each line by space or tab to get key and value
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		// Parse the value as an integer
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid integer value: %s", parts[1])
		}

		// Store the key-value pair in the map
		data[parts[0]] = value
	}

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData("filename")
		db := singletonDatabase{caps}
		if e != nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}
