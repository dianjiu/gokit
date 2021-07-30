package file

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// SelfDir gets compiled executable file directory
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func FileIsDir(file string) bool {
	f, err := os.Stat(file)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// SearchFile Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullpath string, err error) {
	for _, path := range paths {
		if fullpath = filepath.Join(path, filename); FileExists(fullpath) {
			return
		}
	}
	err = errors.New(fullpath + " not found in paths")
	return
}

func ReadFile(filename string) (string, error) {
	body, err := ioutil.ReadFile(filename)
	return string(body), err
}

func IsRegularFile(fi os.FileInfo) bool {
	return fi.Mode()&(os.ModeDir|os.ModeSymlink|os.ModeNamedPipe|os.ModeSocket|os.ModeDevice) == 0
}

// like command grep -E
// for example: GrepFile(`^hello`, "hello.txt")
// \n is striped while read
func GrepFile(patten string, filename string) (lines []string, err error) {
	re, err := regexp.Compile(patten)
	if err != nil {
		return
	}

	fd, err := os.Open(filename)
	if err != nil {
		return
	}
	lines = make([]string, 0)
	reader := bufio.NewReader(fd)
	prefix := ""
	isLongLine := false
	for {
		byteLine, isPrefix, er := reader.ReadLine()
		if er != nil && er != io.EOF {
			return nil, er
		}
		if er == io.EOF {
			break
		}
		line := string(byteLine)
		if isPrefix {
			prefix += line
			continue
		} else {
			isLongLine = true
		}

		line = prefix + line
		if isLongLine {
			prefix = ""
		}
		if re.MatchString(line) {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

// 按分块数分块
func SplitFileByPartNum(fileSize int64, chunkNum int) ([]FileChunk, error) {
	if chunkNum <= 0 || chunkNum > 10000 {
		return nil, errors.New("chunkNum invalid")
	}

	if int64(chunkNum) > fileSize {
		return nil, errors.New("oss: chunkNum invalid")
	}

	var chunks []FileChunk
	var chunk = FileChunk{}
	var chunkN = (int64)(chunkNum)
	for i := int64(0); i < chunkN; i++ {
		chunk.Number = int(i + 1)
		chunk.Offset = i * (fileSize / chunkN)
		if i == chunkN-1 {
			chunk.Size = fileSize/chunkN + fileSize%chunkN
		} else {
			chunk.Size = fileSize / chunkN
		}
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

type FileChunk struct {
	Number int   // Chunk number
	Offset int64 // Chunk offset
	Size   int64 // Chunk size.
}

// 按大小分块， 参考阿里云
func SplitFileByPartSize(fileSize int64, chunkSize int64) ([]FileChunk, error) {
	if chunkSize <= 0 {
		return nil, errors.New("chunkSize invalid")
	}

	if fileSize <= 0 {
		return nil, errors.New("fileSize invalid")
	}

	if chunkSize > fileSize {
		return nil, errors.New("chunkSize must less then fileSize")
	}

	var chunkN = fileSize / chunkSize
	if chunkN >= 10000 {
		return nil, errors.New("Too many parts, please increase part size ")
	}

	var chunks []FileChunk
	var chunk = FileChunk{}
	for i := int64(0); i < chunkN; i++ {
		chunk.Number = int(i + 1)
		chunk.Offset = i * chunkSize
		chunk.Size = chunkSize
		chunks = append(chunks, chunk)
	}

	if fileSize%chunkSize > 0 {
		chunk.Number = len(chunks) + 1
		chunk.Offset = int64(len(chunks)) * chunkSize
		chunk.Size = fileSize % chunkSize
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}
