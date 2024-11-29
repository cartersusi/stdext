package filepath

import "os"

// ListDir returns a list of files in a directory.
// If recurse is true, it will return all files in all subdirectories.
// If recurse is false, it will return only the files in the directory.
//
// Parameters:
//   - fpath: the path to the directory
//   - recurse: whether to recurse into subdirectories
//
// Returns:
//   - a list of files in the directory
//   - an error if one occurred
func ListDir(fpath string, recurse ...bool) ([]string, error) {
	if len(recurse) > 0 && recurse[0] {
		return lsr(fpath)
	}
	var files []string
	dir, err := os.ReadDir(fpath)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		files = append(files, file.Name())
	}

	return files, nil
}

func lsr(fpath string) ([]string, error) {
	var files []string
	dir, err := os.ReadDir(fpath)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		if file.IsDir() {
			subdir, err := lsr(fpath + string(os.PathSeparator) + file.Name())
			if err != nil {
				return nil, err
			}
			files = append(files, subdir...)
		} else {
			files = append(files, file.Name())
		}

	}

	return files, nil
}
