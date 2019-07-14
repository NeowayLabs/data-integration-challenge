package services

import (
	"github.com/mholt/archiver"
)

// ExtractFiles funcao para descompactar arquivos do tipo RAR
func ExtractFiles(filepath string, outpath string) error {

	err := archiver.Unarchive(filepath, outpath)
	if err != nil {
		return err
	}

	return nil
}
