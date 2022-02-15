# QueryCrate
### A simple library for loading & getting string queries from files.

What does this thing solves? Not all libs provide adequate way to keep queries not directly in the code, but in files with .sql extensions. This small library solves this problem by reading query files recursively and providing a friendly interface to access them. Flexible files filters are allowed, so you can keep any text files here (binary aren't supported as internally using string type for keeping query files content).

## Docs
#### Code is placed in qc folder. So import should look like this: `import "github.com/fakefloordiv/querycrate/qc"`.

### Library also has it's own queries locators - this is unix-like path to file, including file name, but excluding file extension. For example, `path/to/my/query.sql` may be gotten from QueryCrate by `path/to/my/query` locator. Also locator does not starts with a root path
---
### Interfaces:
  - `QueryCrate` - query crate object. Implements such a methods:
    - `Get(path string) (query string, err error)` - get query by it's query locator
    - `GetOr(path string, instead interface{}) interface{}` - get query by it's locator, and return `instead` object if query locator does not presents any added queries
    - `AddQuery(path string) error` - add a query to the QueryCrate. Root path won't be excluded from the beginning as query is being added as an external

### Functions:
  - `NewQueryCrate(rootPath string, filters... querycrate.Filter) querycrate.QueryCrate` - returns an instance of `QueryCrate` with already added queries
