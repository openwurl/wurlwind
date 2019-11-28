package debug

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const (
	defaultFilename        = "debug-%s-%s"
	defaultFilePermissions = 0644
	defaultTimestampFormat = "3_04_PM"
)

// WriteDebugFile writes a payload to a debug file
func WriteDebugFile(name string, data interface{}) error {
	timestamp := time.Now().Format(defaultTimestampFormat)
	name = strings.ReplaceAll(name, " ", "_")
	filename := fmt.Sprintf(defaultFilename, name, timestamp)

	file, err := json.MarshalIndent(data, "", " ")
	var writeErr error
	if err != nil {
		writeErr = ioutil.WriteFile(filename, data.([]byte), 0644)
	} else {
		writeErr = ioutil.WriteFile(filename, file, 0644)
	}

	if writeErr != nil {
		Log("ERROR", "Error writing debug output to %s", filename)
	}

	Log("Debug", "Wrote debug output to %s", filename)

	return nil
}

// Log writes a debug log message to console when TF_LOG=debug
func Log(title string, message string, opts ...interface{}) {
	msg := fmt.Sprintf("===== [%s] %s", title, message)

	log.Printf(msg, opts...)
	log.Println("============================================")
}
