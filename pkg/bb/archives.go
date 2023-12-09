package bb

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"

	"filippo.io/age"
	"github.com/pkg/errors"
)

func NewEncryptedTarball(writer io.Writer, filePaths []string, password string) (io.Writer, error) {
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	tarWriter, err := NewTarball(writer, filePaths)
	if err != nil {
		return nil, err
	}
	defer tarWriter.Close()

	identity, err := age.NewScryptRecipient(password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create scrypt identity")
	}
	w, err := age.Encrypt(tarWriter, identity)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encrypt tarball")
	}
	return w, nil
}

func NewTarball(w io.Writer, filePaths []string) (*tar.Writer, error) {
	gzw := gzip.NewWriter(w)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	for _, f := range filePaths {
		err := tarFile(f, tw)
		if err != nil {
			errors.Wrapf(err, "failed to add file %s to tar writer", f)
		}
	}
	return tw, nil
}

func tarFile(path string, w *tar.Writer) error {
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
