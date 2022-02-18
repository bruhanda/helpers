package helpers

import "os"

func StringToFile(data string, file string) error {
	data += "\n"
	var fo *os.File
	var _, err = os.Stat(file)
	// create file if not exists
	if os.IsNotExist(err) {
		fo, err = os.Create(file)
		if err != nil {
			return err
		}
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()
	} else {
		fo, err = os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()
	}

	_, err = fo.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
