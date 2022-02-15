package qc

type Filter interface {
	IsAllowed(File) bool
}

type acceptExtensions struct {
	exts []string
}

func (a acceptExtensions) IsAllowed(file File) bool {
	return contains(a.exts, file.Extension)
}

/*
	Passes only case-sensitive extensions that are allowed
*/
func AllowExtensions(exts ...string) Filter {
	return acceptExtensions{
		exts: exts,
	}
}

type excludeExtensions struct {
	exts []string
}

func (e excludeExtensions) IsAllowed(file File) bool {
	return !contains(e.exts, file.Extension)
}

/*
	Passes only files with extensions that aren't listed
*/
func IgnoreExtensions(exts ...string) Filter {
	return excludeExtensions{
		exts: exts,
	}
}

func contains(arr []string, thing string) bool {
	for _, elem := range arr {
		if elem == thing {
			return true
		}
	}

	return false
}
