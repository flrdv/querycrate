package querycrate

import (
	"errors"
	"fmt"
)

type QueryCrate interface {
}

type queryCrate struct {
	queries map[string]string
}

func NewQueryCrate(rootPath string, filters ...Filter) (QueryCrate, error) {
	filesTree, err := buildFilesTree(rootPath)

	if err != nil {
		return nil, err
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

	return queryCrate{
		queries: queries,
	}, nil
}

func (q queryCrate) GetOrEmpty(name string) string {
	if query, found := q.queries[name]; found {
		return query
	}

	return ""
}

func (q queryCrate) Get(name string) (string, error) {
	if query, found := q.queries[name]; found {
		return query, nil
	}

	return "", errors.New(fmt.Sprintf(`ErrQueryNotFound: query named "%s" not found`, name))
}
