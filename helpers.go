package celeritas

import "os"

func (c *Celeritas) CreateDirIfNotExist(path string) error {
	// create folder if it doesn't exist
	const mode = 0755
	if _,err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Celeritas) CreateFileIfNotExist(path string) error {
	// create folder if it doesn't exist
	const mode = 0644
	if _,err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}
