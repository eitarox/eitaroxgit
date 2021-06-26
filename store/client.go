package store

import (
	"compress/zlib"
	"os"
	"path/filepath"

	"github.com/eitarox/eitaroxgit/object"
	"github.com/eitarox/eitaroxgit/sha"
	"github.com/eitarox/eitaroxgit/util"
)

type Client struct {
	objectDir string
}

func newClient(path string) (*Client, error) {
	rootDir, err := util.FindGitRoot(path)
	if err != nil {
		return nil, err
	}
	return &Client{
		objectDir: filepath.Join(rootDir, ".git", "object"),
	}, nil
}

func (c *Client) GetObject(hash sha.SHA1) (*object.Object, error) {
	hashString := hash.String()
	objectPath := filepath.Join(c.objectDir, hashString[:2], hashString[2:])

	objectFile, err := os.Open(objectPath)
	if err != nil {
		return nil, err
	}
	defer objectFile.Close()

	zr, err := zlib.NewReader(objectFile)
	if err != nil {
		return nil, err
	}

	obj, err := object.ReadObject(zr)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
