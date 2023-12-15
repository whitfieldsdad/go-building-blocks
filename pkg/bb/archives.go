package bb

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"

	"filippo.io/age"
	"github.com/pkg/errors"
)

func NewEncryptedTarball(w io.Writer, files []string, password string) error {
	gw := gzip.NewWriter(w)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, file := range files {
		err := addFileToTarFile(file, tw)
		if err != nil {
			return errors.Wrapf(err, "failed to add file to archive: %s", file)
		}
	}

	identity, err := age.NewScryptRecipient(password)
	if err != nil {
		return errors.Wrap(err, "failed to create scrypt identity")
	}
	ew, err := age.Encrypt(tw, identity)
	if err != nil {
		return errors.Wrap(err, "failed to encrypt tarball")
	}
	defer ew.Close()

	return nil
}

func NewTarball(w io.Writer, files []string) error {
	gw := gzip.NewWriter(w)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, file := range files {
		err := addFileToTarFile(file, tw)
		if err != nil {
			return errors.Wrapf(err, "failed to add file to archive: %s", file)
		}
	}
	return nil
}

func addFileToTarFile(path string, w *tar.Writer) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}
	defer f.Close()

	s, err := f.Stat()
	if err != nil {
		return errors.Wrap(err, "failed to stat file")
	}
	h := &tar.Header{
		Name:    path,
		Size:    s.Size(),
		Mode:    int64(s.Mode()),
		ModTime: s.ModTime(),
	}
	err = w.WriteHeader(h)
	if err != nil {
		return errors.Wrap(err, "failed to write tar header")
	}
	_, err = io.Copy(w, f)
	if err != nil {
		return errors.Wrap(err, "failed to copy file to tar writer")
	}
	return nil
}

func ReadTarball(r io.Reader) (*tar.Reader, error) {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	tr := tar.NewReader(gzr)
	return tr, nil
}

func ReadEncryptedTarball(r io.Reader, password string) (*tar.Reader, error) {
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	id, err := age.NewScryptIdentity(password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create identity")
	}
	r, err = age.Decrypt(r, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decrypt tarball")
	}
	return ReadTarball(r)
}
