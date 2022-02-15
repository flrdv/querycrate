package qc

import (
	"errors"
	"fmt"
)

const maxRecursionDepth = 15

type QueryCrate interface {
	GetOr(name string, otherwise interface{}) interface{}
	Get(name string) (string, error)
	AddQuery(path string) error
}

type queryCrate struct {
	queries map[string]string
}

func NewQueryCrate(rootPath string, filters ...Filter) (QueryCrate, error) {
	filesTree, err := buildFilesTree(rootPath, maxRecursionDepth)

	if err != nil {
		return nil, err
	}

	if len(filters) == 0 {
		// all files with .sql extensions are allowed by default
		filters = append(filters, AllowExtensions(".sql"))
	}

	queryFiles := getFilteredFiles(filesTree, filters...)
	queries := make(map[string]string)

	for _, file := range queryFiles {
		fileContent, err := file.Read()

		if err != nil {
			return nil, errors.New(fmt.Sprintf(`ErrInvalidQueryFile: failed to read query file: "%s"`, file.Name))
		}

		queries[getQueryPath(file)] = string(fileContent)
	}

	return &queryCrate{
		queries: queries,
	}, nil
}

func (q queryCrate) GetOr(name string, instead interface{}) interface{} {
	if query, found := q.queries[name]; found {
		return query
	}

	return instead
}

func (q queryCrate) Get(name string) (string, error) {
	if query, found := q.queries[name]; found {
		return query, nil
	}

	return "", errors.New(fmt.Sprintf(`ErrQueryNotFound: query named "%s" not found`, name))
}

func (q *queryCrate) AddQuery(path string) error {
	queryName, queryValue, err := getFile(path)

	if err != nil {
		return err
	}

	q.queries[queryName] = queryValue

	return nil
}
