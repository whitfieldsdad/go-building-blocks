package bb

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"filippo.io/age"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/afero/tarfs"
)

type FileOptions struct {
	IncludeAll        bool `json:"include_all,omitempty"`
	IncludeFileHashes bool `json:"include_file_hashes,omitempty"`
}

func NewFileOptions() *FileOptions {
	return &FileOptions{
		IncludeAll: true,
	}
}

type File struct {
	Path      string  `json:"path"`
	Directory string  `json:"directory,omitempty"`
	Filename  string  `json:"filename"`
	Size      *int    `json:"size,omitempty"`
	Hashes    *Hashes `json:"hashes,omitempty"`
}

func NewFile(path string) *File {
	return &File{
		Path:      path,
		Filename:  filepath.Base(path),
		Directory: filepath.Dir(path),
	}
}

func GetFile(path string, opts *FileOptions) (*File, error) {
	file := NewFile(path)
	info, err := os.Stat(path)
	if err != nil {
		return file, err
	}
	if opts == nil {
		opts = NewFileOptions()
	}
	sz := int(info.Size())
	file.Size = &sz

	if opts.IncludeAll || opts.IncludeFileHashes {
		file.Hashes, _ = GetFileHashes(path)
	}
	return file, nil
}

type FileFilter struct {
	FilenamePatterns []string `json:"filenames"`
	PathPatterns     []string `json:"paths"`
}

func (f FileFilter) Matches(path string, info os.FileInfo) (bool, error) {
	if len(f.FilenamePatterns) > 0 {
		matches, err := StringMatchesAnyPattern(filepath.Base(path), f.FilenamePatterns)
		if !matches || err != nil {
			return false, err
		}
	}
	if len(f.PathPatterns) > 0 {
		matches, err := StringMatchesAnyPattern(path, f.PathPatterns)
		if !matches || err != nil {
			return false, err
		}
	}
	return true, nil
}

type FileStat struct {
	Path string
	Info os.FileInfo
}

func GetFS() afero.Afero {
	fs := afero.NewOsFs()
	return afero.Afero{Fs: fs}
}

func GetTarFS(path, password string) (*afero.Afero, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer file.Close()

	var reader io.Reader
	reader = file

	// If a password was provided, decrypt the tarball.
	if password != "" {
		id, err := age.NewScryptIdentity(password)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create identity")
		}
		reader, err = age.Decrypt(file, id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decrypt tarball")
		}
	}

	// If the tarball is GZIP compressed, decompress it.
	if strings.HasSuffix(path, ".tar.gz") || strings.HasSuffix(path, ".tar.gz.age") {
		reader, err = gzip.NewReader(reader)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create gzip reader")
		}
	}

	// Create a tar filesystem from the tarball.
	fs := tarfs.New(tar.NewReader(reader))
	afs := &afero.Afero{Fs: fs}
	return afs, nil
}

func FindFiles(afs afero.Afero, filter *FileFilter) ([]File, error) {
	var results []File
	afs.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filter != nil {
			matches, err := filter.Matches(path, info)
			if !matches || err != nil {
				return err
			}
		}
		result := NewFile(path)
		results = append(results, *result)
		return nil
	})
	return results, nil
}

func FindFilesInTarball(path, password string, filter *FileFilter) ([]File, error) {
	afs, err := GetTarFS(path, password)
	if err != nil {
		return nil, err
	}
	return FindFiles(*afs, filter)
}

func CopyFile(src, dst string) error {
	dir, err := IsDirectory(src)
	if err != nil {
		return err
	}
	if dir {
		return CopyDirectory(src, dst)
	} else {
		return copyFile(src, dst)
	}
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, "failed to open source file")
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return errors.Wrap(err, "failed to open destination file")
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return errors.Wrap(err, "failed to copy file")
	}
	return nil
}

func CopyDirectory(src, dst string) error {
	srcFiles, err := os.ReadDir(src)
	if err != nil {
		return errors.Wrap(err, "failed to read source directory")
	}
	err = os.MkdirAll(dst, 0755)
	if err != nil {
		return errors.Wrap(err, "failed to create destination directory")
	}
	for _, file := range srcFiles {
		srcPath := src + "/" + file.Name()
		dstPath := dst + "/" + file.Name()
		if file.IsDir() {
			err = CopyDirectory(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func DeleteFile(path string) error {
	dir, err := IsDirectory(path)
	if err != nil {
		return errors.Wrap(err, "failed to check if path is a directory")
	}
	if dir {
		return os.RemoveAll(path)
	}
	return os.Remove(path)
}

func MoveFile(src, dst string) error {
	return os.Rename(src, dst)
}
