package config

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var pidfile string

type File struct {
	*os.File
	path string
}

func SetPidFile(p string) {
	pidfile = p
}

func GetPidFile() string {
	return pidfile
}

func WritePidToFile() error {
	if pidfile == "" {
		return errors.New("pidfile not set")
	}

	// mkdir
	if err := os.MkdirAll(filepath.Dir(pidfile), os.FileMode(0755)); err != nil {
		return err
	}

	file, err := New(pidfile, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("error opening pidfile %s: %s", pidfile, err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%d", os.Getpid())
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func New(path string, mode os.FileMode) (*File, error) {
	f, err := ioutil.TempFile(filepath.Dir(path), filepath.Base(path))
	if err != nil {
		return nil, err
	}
	if err := os.Chmod(f.Name(), mode); err != nil {
		f.Close()
		os.Remove(f.Name())
		return nil, err
	}
	return &File{File: f, path: path}, nil
}

func Read() (int, error) {
	if pidfile == "" {
		return 0, errors.New("pidfile not set")
	}

	d, err := ioutil.ReadFile(pidfile)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(bytes.TrimSpace(d)))
	if err != nil {
		return 0, fmt.Errorf("error parsing pid from %s: %s", pidfile, err)
	}

	return pid, nil
}

func (f *File) Close() error {
	if err := f.File.Close(); err != nil {
		os.Remove(f.File.Name())
		return err
	}
	if err := os.Rename(f.Name(), f.path); err != nil {
		return err
	}
	return nil
}
