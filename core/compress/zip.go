package compress

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Zip Compressed file.
// src is the abbreviation of source, which is a directory of files to be compressed.
// dst is the abbreviation of destination, which is the compressed file directory.
func Zip(src, dest string) (err error) {
	// Create a file to be written.
	fw, err := os.Create(dest)
	defer fw.Close()
	if err != nil {
		return err
	}

	// Create zip.Write through fw.
	zw := zip.NewWriter(fw)
	defer func() {
		// Check if it is successfully closed.
		if err := zw.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Write the file to zw, because there may be many directories and files, so recursive processing.
	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}

		// Create zip file information through file information.
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return
		}

		// Replace the file name in the file information
		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))

		// Judge it is not a directory.
		if fi.IsDir() {
			fh.Name += "/"
		}

		// Write file information and return a Write structure.
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return
		}

		// If it is not a standard file, only the header information is written, and the file data is not written to w.
		// Such as a directory, there is no data to write.
		if !fh.Mode().IsRegular() {
			return nil
		}

		// Open the file to be compressed.
		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			return
		}

		// Copy the opened file to w .
		n, err := io.Copy(w, fr)
		if err != nil {
			return
		}
		// Output compressed content.
		fmt.Printf("Zip success %s, A total of %d characters of data are written\n", path, n)

		return nil
	})
}

// UnZip decompressed file.
// src is the abbreviation of source, which is a directory of files to be decompressed.
// dst is the abbreviation of destination,  which is the decompressed file directory.
func UnZip(dst, src string) (err error) {
	// Open compressed file.
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		return
	}

	// If it is not placed in the current directory after decompression, create a directory according to the save directory.
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	// Traverse zr and write the file to disk
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		// If it is a directory, create a directory.
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			// Because it is a directory, skip the current loop, because the following is the processing of files.
			continue
		}

		// Get the Reader.
		fr, err := file.Open()
		if err != nil {
			//Also close when abnormal
			close(fr, nil)
			return err
		}

		// Create the Write corresponding to the file to be written.
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			//Also close when abnormal
			close(fr, fw)
			return err
		}

		n, err := io.Copy(fw, fr)
		if err != nil {
			//Also close when abnormal
			close(fr, fw)
			return err
		}

		// Output the decompressed result.
		fmt.Printf("UnZip success %s A total of %d characters have been written\n", path, n)

		//Finally remember to close
		close(fr, fw)
	}
	return nil
}

// close Close file stream.
func close(fr io.ReadCloser, fw *os.File) {
	if fr != nil {
		fr.Close()
	}
	if fw != nil {
		fw.Close()
	}

}
