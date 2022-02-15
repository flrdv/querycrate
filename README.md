# QueryCrate
### A simple library for loading & getting string queries from files.

## How to install
`go get github.com/fakefloordiv/querycrate/tree/master`

## Docs
#### Code is placed in qc folder. So import should look like this: `import "github.com/fakefloordiv/querycrate/qc"`.

### Library also has it's own queries locators - this is unix-like path to file, including file name, but excluding file extension. For example, `path/to/my/query.sql` may be gotten from QueryCrate by `path/to/my/query` locator. Also locator does not starts with a root path
---
### Filters
Filters are just structs that implement interface `querycrate.Filter` with a single method `IsAllowed(file querycrate.File) bool`. Library already has 2 simple filters - `AllowExtensions(extensions... string)` and `IgnoreExtensions(extensions... string)` (to be honest they're just returning an initialized struct). Filters may stacking, so you can add multiple filters


## Usage
```golang
import (
  "log"

  querycrate "github.com/fakefloordiv/querycrate/qc"
)

func main() {
  root := "somefolder/"
  qc, err := querycrate.NewQueryCrate(root) // by default only .sql files are allowed
  // or querycrate.NewQueryCrate(root, SomeFilter1, SomeFilter2, ...)
  
  if err != nil {
    log.Fatal("unexpected error:", err)
  }
  
  query, err := qc.Get("somerepo/myquery")
  
  if err != nil {
    if err = qc.AddQuery("somefolder/somerepo/myquery.sql"); err != nil {
      log.Fatal("unexpected error during adding query")
    }
    
    instead := "default query"
    query = qc.GetOr("somefolder/somerepo/myquery", instead).(string)
  }
  
  // For example
  sql.Exec(query)
}
```

## What problems this library solves?
Not all libs provide adequate way to keep queries not directly in the code, but in files with .sql extensions. This small library solves this problem by reading query files recursively and providing a friendly interface to access them. Flexible files filters are allowed, so you can keep any text files here (binary aren't supported as internally using string type for keeping query files content).
