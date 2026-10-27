//go:build !darwin

package spotlight

import (
	"errors"
	"time"
)

type SearchOptions struct {
	Query      string
	Scope      []string
	MaxResults int
}

type FileResult struct {
	Path string
	Name string
}

type FileInfo struct {
	Path     string
	Name     string
	Size     int64
	Modified time.Time
	IsDir    bool
}

func Search(_ SearchOptions) ([]FileResult, error) {
	return nil, errors.New("spotlight search is unsupported on this platform (requires macOS)")
}

func SearchWithMetadata(_ SearchOptions) ([]FileInfo, error) {
	return nil, errors.New("spotlight search is unsupported on this platform (requires macOS)")
}
