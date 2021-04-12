package compress

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// Tar Compressed file.
// src is the abbreviation of source, which is a directory of files to be compressed.
// dst is the abbreviation of destination, which is the compressed file directory.
func Tar(src, dest string) (err error) {
	// file write
	fw, err := os.Create(dest)
	defer fw.Close()
	if err != nil {
		return err
	}
	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// 打开文件夹
	dir, err := os.Open(src)
	defer dir.Close()
	if err != nil {
		return err
	}
	// 读取文件列表
	fis, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	// 遍历文件列表
	for _, fi := range fis {
		// 逃过文件夹, 我这里就不递归了
		if fi.IsDir() {
			continue
		}
		// 打印文件名称
		fmt.Println(fi.Name())
		// 打开文件
		fr, err := os.Open(dir.Name() + "/" + fi.Name())
		if err != nil {
			return err
		}
		defer fr.Close()
		// 信息头
		h := new(tar.Header)
		h.Name = fi.Name()
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()
		// 写信息头
		err = tw.WriteHeader(h)
		if err != nil {
			return err
		}
		// 写文件
		_, err = io.Copy(tw, fr)
		if err != nil {
			return err
		}
	}
	fmt.Println("tar.gz ok")
	return nil
}

// UnTar decompressed file.
// src is the abbreviation of source, which is a directory of files to be decompressed.
// dst is the abbreviation of destination,  which is the decompressed file directory.
func UnTar(src, dest string) (err error) {
	// file read
	fr, err := os.Open(src)
	defer fr.Close()
	if err != nil {
		panic(err)
	}
	// tar read
	tr := tar.NewReader(fr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if h.FileInfo().IsDir() {
			err = os.MkdirAll(dest+h.Name, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}
		// 打开文件
		fw, err := os.OpenFile(dest+h.Name, os.O_CREATE|os.O_WRONLY, 0666 /*os.FileMode(h.Mode)*/)
		if err != nil {
			return err
		}
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	fmt.Println("un tar.gz ok")
	return nil
}
