package db

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"io"
	"main/app/db/build"
	"os"
)

const (
	dbPath = "./app/db/data/ip2region.xdb"
)

var searcher *xdb.Searcher

func InitSearcher() error {
	file, err := os.Open(dbPath)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(b)

	searcher, err = xdb.NewWithBuffer(buffer.Bytes())
	return nil
}

func Search(ip string) (string, error) {
	if searcher == nil {
		return "", errors.New("empty searcher")
	}
	loc, err := searcher.SearchByStr(ip)
	if err != nil {
		return "", err
	}
	return loc, nil
}

func GenerateDB(srcFile, dstFile string) error {
	var err error
	var indexPolicy = build.VectorIndexPolicy

	if srcFile == "" || dstFile == "" {
		return fmt.Errorf("param cannot be null")
	}

	// make the binary file
	maker, err := build.NewMaker(indexPolicy, srcFile, dstFile)
	if err != nil {
		return fmt.Errorf("failed to create %s\n", err)
	}

	err = maker.Init()
	if err != nil {
		return fmt.Errorf("failed Init: %s\n", err)
	}

	err = maker.Start()
	if err != nil {
		return fmt.Errorf("failed Start: %s\n", err)
	}

	err = maker.End()
	if err != nil {
		return fmt.Errorf("failed End: %s\n", err)
	}

	return nil
}
