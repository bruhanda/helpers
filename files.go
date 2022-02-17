package helpers

import "os"

func WriteToFile(file string, data []byte) error {
	fo, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	_, err = fo.Write(data)
	if err != nil {
		return err
	}
	return nil
}
