package bb

import (
	"os"

	"golang.org/x/exp/slices"
)

const (
	Directory       = "directory"
	RegularFile     = "regular_file"
	SymbolicLink    = "symbolic_link"
	Socket          = "socket"
	HardLink        = "hard_link"
	NamedPipe       = "named_pipe"
	BlockDevice     = "block_device"
	CharacterDevice = "character_device"
	Hidden          = "hidden"
)

type FileTraits struct {
	IsDirectory       bool `json:"is_directory"`
	IsRegularFile     bool `json:"is_regular_file"`
	IsSymbolicLink    bool `json:"is_symbolic_link"`
	IsSocket          bool `json:"is_socket"`
	IsHardLink        bool `json:"is_hard_link"`
	IsNamedPipe       bool `json:"is_named_pipe"`
	IsBlockDevice     bool `json:"is_block_device"`
	IsCharacterDevice bool `json:"is_character_device"`
	IsHidden          bool `json:"is_hidden"`
}

func (traits FileTraits) ToList() []string {
	checks := map[string]bool{
		BlockDevice:     traits.IsBlockDevice,
		CharacterDevice: traits.IsCharacterDevice,
		Directory:       traits.IsDirectory,
		HardLink:        traits.IsHardLink,
		Hidden:          traits.IsHidden,
		NamedPipe:       traits.IsNamedPipe,
		RegularFile:     traits.IsRegularFile,
		Socket:          traits.IsSocket,
		SymbolicLink:    traits.IsSymbolicLink,
	}
	list := make([]string, 0)
	for trait, check := range checks {
		if check {
			list = append(list, trait)
		}
	}
	slices.Sort(list)
	return list
}

func GetFileTraits(path string) (*FileTraits, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	traits := &FileTraits{
		IsBlockDevice:     isBlockDevice(st),
		IsCharacterDevice: isCharacterDevice(st),
		IsDirectory:       isDirectory(st),
		IsHardLink:        isHardLink(st),
		IsNamedPipe:       isNamedPipe(st),
		IsRegularFile:     isRegularFile(st),
		IsSocket:          isSocket(st),
		IsSymbolicLink:    isSymbolicLink(st),
	}
	traits.IsHidden, _ = IsHidden(path)
	return traits, nil
}

func IsHidden(path string) (bool, error) {
	return isHidden(path)
}

func hasTrait(path string, f func(os.FileInfo) bool) (bool, error) {
	st, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return f(st), nil
}

func IsDirectory(path string) (bool, error) {
	return hasTrait(path, isDirectory)
}

func IsRegularFile(path string) (bool, error) {
	return hasTrait(path, isRegularFile)
}

func IsSymbolicLink(path string) (bool, error) {
	return hasTrait(path, isSymbolicLink)
}

func IsHardLink(path string) (bool, error) {
	return hasTrait(path, isHardLink)
}

func IsSocket(path string) (bool, error) {
	return hasTrait(path, isSocket)
}

func IsNamedPipe(path string) (bool, error) {
	return hasTrait(path, isNamedPipe)
}

func IsBlockDevice(path string) (bool, error) {
	return hasTrait(path, isBlockDevice)
}

func IsCharacterDevice(path string) (bool, error) {
	return hasTrait(path, isCharacterDevice)
}

func isDirectory(st os.FileInfo) bool {
	return st.Mode().IsDir()
}

func isRegularFile(st os.FileInfo) bool {
	return st.Mode().IsRegular()
}

func isSymbolicLink(st os.FileInfo) bool {
	return st.Mode()&os.ModeSymlink != 0
}

func isHardLink(st os.FileInfo) bool {
	return st.Mode()&os.ModeDevice != 0
}

func isSocket(st os.FileInfo) bool {
	return st.Mode()&os.ModeSocket != 0
}

func isNamedPipe(st os.FileInfo) bool {
	return st.Mode()&os.ModeNamedPipe != 0
}

func isBlockDevice(st os.FileInfo) bool {
	return st.Mode()&os.ModeDevice != 0
}

func isCharacterDevice(st os.FileInfo) bool {
	return st.Mode()&os.ModeCharDevice != 0
}
