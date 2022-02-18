package helpers

import "os"

func WriteToFile(file string, data string) error {
	fo, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	_, err = fo.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
