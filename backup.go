package main

import (
	tar2 "archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Settings struct {
	BackupDir []string `json:"backup_dir"`
	StoreDir  string   `json:"store_dir"`
}

func NewSettings() Settings {
	s := Settings{}
	if len(os.Args) == 1 {
		log.Fatalln("you should set path of settings.json")
		log.Fatalln("example: bamp ./settings.json")
		os.Exit(1)
	}
	settingsPath := os.Args[1]

	data, err := ioutil.ReadFile(settingsPath)
	if err != nil {
		log.Fatalln("can not read settings.json")
		os.Exit(1)
		return Settings{}
	}

	err = json.Unmarshal(data, &s)
	if err != nil {
		log.Fatalln("can not marshal")
		os.Exit(1)
	}
	return s
}

func main() {

	s := NewSettings()
	log.SetOutput(os.Stdout)

	now := time.Now()
	store := strconv.Itoa(now.Year()) + "-" + now.Month().String() + "-" + strconv.Itoa(now.Day()) + ".tar.gz"
	storePath := filepath.Join(s.StoreDir, store)
	file, err := os.Create(storePath)
	if err != nil {
		log.Fatalln(err)
	}

	gz := gzip.NewWriter(file)
	defer gz.Close()

	tw := tar2.NewWriter(gz)
	defer tw.Close()

	for _, root := range s.BackupDir {
		/*
			isIgnore := false

			for _, list := range BLACK_LIST {
				if root == list {
					isIgnore = true
				}
			}

			if isIgnore {
				continue
			}
		*/
		filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			log.Println("file name: " + path)
			err = walk(root, path, tw, info, err)
			if err != nil {
				return err
			}
			return nil
		})
	}

	tw.Flush()
	err = gz.Flush()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("success")
}
func walk(root, filePath string, tw *tar2.Writer, f fs.FileInfo, err error) error {
	if !f.Mode().IsRegular() {
		return nil
	}

	header, err := tar2.FileInfoHeader(f, f.Name())
	if err != nil {
		log.Fatalln(err)
		return err
	}
	header.Name = strings.TrimPrefix(strings.Replace(filePath, root, "", -1), string(filepath.Separator))

	err = tw.WriteHeader(header)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if !f.IsDir() {
		tmpFile, err := os.Open(filePath)

		if err != nil {
			log.Fatalln(err)
			return err
		}

		if _, err = io.Copy(tw, tmpFile); err != nil {
			log.Fatalln(err)
			return err
		}
		tmpFile.Close()
	}
	return nil
}
