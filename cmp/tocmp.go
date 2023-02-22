package cmp

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func CreateArchive(files []string, buf io.Writer) error {
	gw := gzip.NewWriter(buf)
	defer gw.Close()
	tarWriter := tar.NewWriter(gw)
	defer tarWriter.Close()

	// Iterate over files and add them to the tar archive
	for _, file := range files {
		err := addToArchive(tarWriter, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func addToArchive(tarWriter *tar.Writer, fileName string) error {
	// Open the file which will be written into the archive
	file, fileError := os.Open(fileName)
	if fileError != nil {
		log.Println(fileError.Error())
		return fileError
	}
	defer file.Close()

	// Get FileInfo about our file providing file size, mode, etc.
	fileInfo, fileInfoError := file.Stat()
	if fileInfoError != nil {
		log.Println(fileInfoError.Error())
		return fileInfoError
	}

	//Create tar header with the info of file
	tarHeaderInfo, tarHeaderInfoError := tar.FileInfoHeader(fileInfo, fileInfo.Name())
	if tarHeaderInfoError != nil {
		log.Println(tarHeaderInfoError.Error())
		return tarHeaderInfoError
	}

	tarHeaderInfo.Name = fileName

	// Write file header to the tar archive
	writeHeaderError := tarWriter.WriteHeader(tarHeaderInfo)
	if writeHeaderError != nil {
		return writeHeaderError
	}

	// Copy file content to tar archive
	_, copyError := io.Copy(tarWriter, file)
	if copyError != nil {
		return copyError
	}

	return nil
}
