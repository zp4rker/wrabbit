package wrabbit

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/zp4rker/wrabbit/pkg/data"
)

type Statfile struct {
	id   string
	File *os.File
	Data data.StatfileData
}

// create statfile, generate id and return pointer to statfile
func PrepareStatfile() (*Statfile, error) {
	sf := &Statfile{id: RandToken()}

	// create necessary dirs
	dir := fmt.Sprintf("%v/wrabbit/%v", os.TempDir(), sf.id)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	// create statfile
	file, err := os.Create(fmt.Sprintf("%v/statfile.json", dir))
	if err != nil {
		return nil, err
	} else {
		// initialise statfile
		sf.File = file
		sf.Data = data.StatfileData{StartDate: time.Now()}
		return sf, nil
	}
}

// update statfile with current data
func (sf *Statfile) Update() error {
	// marshal data
	data, err := json.Marshal(sf.Data)
	if err != nil {
		return err
	}

	// write data
	_, err = sf.File.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// update modtime of statfile
func (sf *Statfile) Touch() error {
	now := time.Now()

	err := os.Chtimes(sf.File.Name(), now, now)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// poll (touch) statfile every 5 seconds
func (sf *Statfile) StartPoll(pollStop *chan bool) {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-*pollStop:
			return
		case <-ticker.C:
			err := sf.Touch()
			if err != nil {
				fmt.Println("Failed to touch statfile!")
			}
		}
	}
}
