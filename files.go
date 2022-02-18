package helpers

import "os"

func StringToFile(data string, file string) error {
	data += "\n"
	fo, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
