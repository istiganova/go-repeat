package file

import "os"

func WriteFile(path string, content string) error {
	data := []byte(content)

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
