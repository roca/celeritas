package celeritas

import "os"

func (c *Celeritas) CreateDirIfNotExist(path string) error {
	// create folder if it doesn't exist
	const mode = 0755
	if _,err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}
