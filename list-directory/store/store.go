package store

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type TextFileStore struct {
	coins map[string]string
	files []*file
	mu    sync.RWMutex
	dir   string
}

type file struct {
	filename     string
	lastModified time.Time
}

func NewTextFileStore(p string) (*TextFileStore, error) {
	tfs := &TextFileStore{
		dir: p,
	}
	// f, err := os.Open(tfs.path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// files, err := f.Readdir(-1)
	// f.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	files, err := filepath.Glob(filepath.Join(tfs.dir, "bnbusdt.json"))
	if err != nil {
		return nil, err
	}

	for _, e := range files {
		fi, _ := os.Stat(e)
		f := &file{
			filename:     e,
			lastModified: fi.ModTime(),
		}

		tfs.files = append(tfs.files, f)
	}
	go tfs.watchFileChanges()
	return tfs, nil
}

func (tfs *TextFileStore) watchFileChanges() {
	for {
		tfs.mu.Lock()
		for _, f := range tfs.files {
			// fileStats, err := os.Stat(filepath.Join(tfs.dir, f.filename))
			fileStats, err := os.Stat(f.filename)
			if err != nil {
				log.Println(err)
			}

			modTime := fileStats.ModTime()
			// log.Println("Mod Time", modTime)
			// log.Println("Last Mod Time", f.lastModified)
			if f.lastModified.Before(modTime) {
				f.lastModified = modTime
				log.Printf("%q has been modified at %s.", f.filename, modTime.String())
				log.Printf("%q -- %s.", f.filename, f.lastModified.String())
			}
		}
		tfs.mu.Unlock()
		time.Sleep(5 * time.Second)
	}
}

func (tfs *TextFileStore) ListFiles() {
	for _, e := range tfs.files {
		fmt.Printf("%s\t\t\t%s\n", e.filename, e.lastModified.String())
	}
}
