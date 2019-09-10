package fileio

import "io/ioutil"

// FileToString reads the file into a string and returns it or an error
func FileToString(filepath string) (string, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
