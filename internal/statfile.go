package wrabbit

import (
	"fmt"
	"os"
	"time"

	"github.com/zp4rker/wrabbit/internal/data"
)

type Statfile struct {
	id   string
	File *os.File
	Data data.StatfileData
}

func PrepareStatfile() (*Statfile, error) {
	sf := &Statfile{id: RandToken()}

	dir := fmt.Sprintf("%v/wrabbit/%v", os.TempDir(), sf.id)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(fmt.Sprintf("%v/statfile.json", dir))
	if err != nil {
		return nil, err
	} else {
		sf.File = file
		sf.Data = data.StatfileData{StartDate: time.Now()}
		return sf, nil
	}
}

func (sf *Statfile) Touch() error {
	now := time.Now()

	err := os.Chtimes(sf.File.Name(), now, now)
	if err != nil {
		return err
	} else {
		return nil
	}
}
