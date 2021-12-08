package util

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	u "github.com/Truth1984/awadau-go"
)

func FileExists(path ...string) bool {
	fullpath := PathJoin(path...)
	_, err := os.Stat(fullpath)
	if err != nil {
		LTP("FileExists - Unable to get file info", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	} else {
		LTP("FileExists - File exists", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	}
	return err == nil
}

func FileIsDir(path ...string) bool {
	fullpath := PathJoin(path...)
	info, err := os.Stat(fullpath)
	if err != nil {
		LTP("FileIsDir - Unable to get file info", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
		return false
	}
	if info.IsDir() {
		LTP("FileIsDir - File is directory", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	} else {
		LTP("FileIsDir - File is not directory", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	}
	return info.IsDir()
}

func FileIsFile(path ...string) bool {
	fullpath := PathJoin(path...)
	info, err := os.Stat(fullpath)
	if err != nil {
		LTP("FileIsFile - Unable to get file info", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
		return false
	}
	if info.Mode().IsRegular() {
		LTP("FileIsFile - File is regular file", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	} else {
		LTP("FileIsFile - File is not a regular file", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	}
	return info.Mode().IsRegular()
}

func FileStat(path ...string) os.FileInfo {
	fullpath := PathJoin(path...)
	info, err := os.Stat(fullpath)
	if err != nil {
		EHWarn(err, "FileStat - Unable to get file info", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
		return nil
	}
	LTP("FileStat - File info", LogMap(u.Map("path", path), u.Map("fullpath", fullpath, "info", info)))
	return info
}

func FileSize(path ...string) int64 {
	fullpath := PathJoin(path...)
	info, err := os.Stat(fullpath)
	if err != nil {
		EHWarn(err, "FileSize - Unable to get file size", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
		return 0
	}
	size := info.Size()
	LTP("FileSize - File size", LogMap(u.Map("path", path), u.Map("fullpath", fullpath, "size", size)))
	return size
}

func FileLs(path ...string) []string {
	fullpath := PathJoin(path...)
	files, err := filepath.Glob(fullpath)
	if err != nil {
		EHWarn(err, "FileLs - Unable to list files", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
		return nil
	}
	LTP("FileLs - Listing files", LogMap(u.Map("path", path), u.Map("fullpath", fullpath, "files", files)))
	return files
}

func FileMkdir(path ...string) {
	fullpath := PathJoin(path...)
	LTP("FileMkdir - Creating directory", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	EHWarn(os.MkdirAll(fullpath, 0755), "FileMkdir - Unable to create directory", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
}

func FileMove(src string, dst string) {
	LTP("FileMove - Moving file", LogMap(u.Map("src", src), u.Map("dst", dst)))
	EHWarn(os.Rename(src, dst), "FileMove - Unable to move file", LogMap(u.Map("src", src, "dst", dst), nil))
}

func FileExt(path ...string) string {
	fullpath := PathJoin(path...)
	ext := filepath.Ext(fullpath)
	LTP("FileExt - Getting file extension", LogMap(u.Map("path", path), u.Map("fullpath", fullpath, "ext", ext)))
	return ext
}

func PathAbsolute(path ...string) string {
	fullpath := PathJoin(path...)
	fullpathAbs, err := filepath.Abs(fullpath)

	if err != nil {
		EHWarn(err, "PathAbsolute - Unable to get absolute path", LogMap(u.Map("path", path), u.Map("fullpath", fullpath, "fullpathAbs", fullpathAbs)))
		return ""
	}

	LTP("PathAbsolute - Getting absolute path", LogMap(u.Map("path", path), u.Map("fullpath", fullpath, "fullpathAbs", fullpathAbs)))
	return fullpathAbs
}

func PathJoin(path ...string) string {
	fullpath := filepath.Join(path...)
	LTP("PathJoin - Joining path", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	return fullpath
}

func PathDirname(path ...string) string {
	fullpathabs := PathAbsolute(path...)
	dirname := filepath.Dir(fullpathabs)
	LTP("PathDirname - Getting directory name", LogMap(u.Map("path", path), u.Map("fullpath", fullpathabs, "dirname", dirname)))
	return dirname
}

func PathBasename(path ...string) string {
	fullpathabs := PathAbsolute(path...)
	basename := filepath.Base(fullpathabs)
	LTP("PathBasename - Getting basename", LogMap(u.Map("path", path), u.Map("fullpath", fullpathabs, "basename", basename)))
	return basename
}

func PathExt(path ...string) string {
	fullpath := PathJoin(path...)
	ext := filepath.Ext(fullpath)
	LTP("PathExt - Getting extension", LogMap(u.Map("path", path), u.Map("fullpath", fullpath, "ext", ext)))
	return ext
}

func FileRead(path ...string) []byte {
	fullpath := PathJoin(path...)
	data, err := ioutil.ReadFile(fullpath)
	if err != nil {
		EHWarn(err, "FileRead - Unable to read file", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
		return nil
	}
	LTP("FileRead - Reading file", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	return data
}

func FileReadStr(path ...string) string {
	content := string(FileRead(path...))
	LTP("FileReadStr - Reading file", LogMap(u.Map("path", path), u.Map("content", content)))
	return content
}

func FileRemove(path ...string) {
	fullpath := PathJoin(path...)
	LTP("FileRemove - Removing file", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
	EHWarn(os.Remove(fullpath), "FileRemove - Unable to remove file", LogMap(u.Map("path", path), u.Map("fullpath", fullpath)))
}

func FileWrite(path string, body []byte) {
	err := ioutil.WriteFile(path, body, 0644)
	if err != nil {
		EHWarn(err, "FileWrite - Unable to write file", LogMap(u.Map("path", path), nil))
	}
	LTP("FileWrite - Writing file", LogMap(u.Map("path", path), nil))
}

func FileWriteStr(path string, body string) {
	FileWrite(path, []byte(body))
	LTP("FileWriteStr - Writing file", LogMap(u.Map("path", path), nil))
}

func FileZip(path []string, dst string) {
	archive, err := os.Create(dst)
	if err != nil {
		EHWarn(err, "FileZip - Unable to create archive", LogMap(u.Map("path", path, "dst", dst), nil))
		return
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	for _, file := range path {
		zipFile, err := os.Open(file)
		if err != nil {
			EHWarn(err, "FileZip - Unable to open file", LogMap(u.Map("path", path, "dst", dst), nil))
			return
		}
		defer zipFile.Close()

		info, err := zipFile.Stat()
		if err != nil {
			EHWarn(err, "FileZip - Unable to get file info", LogMap(u.Map("path", path, "dst", dst), nil))
			return
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			EHWarn(err, "FileZip - Unable to get file header", LogMap(u.Map("path", path, "dst", dst), nil))
			return
		}

		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			EHWarn(err, "FileZip - Unable to create file", LogMap(u.Map("path", path, "dst", dst), nil))
			return
		}

		_, err = io.Copy(writer, zipFile)
		if err != nil {
			EHWarn(err, "FileZip - Unable to copy file", LogMap(u.Map("path", path, "dst", dst), nil))
			return
		}

		LTP("FileZip - Zipping file", LogMap(u.Map("path", path, "dst", dst), nil))
	}
}

func FileUnzip(path string, dst string) {
	if !FileIsDir(dst) {
		FileMkdir(dst)
	}
	zipReader, err := zip.OpenReader(path)
	if err != nil {
		EHWarn(err, "FileUnzip - Unable to open archive", LogMap(u.Map("path", path, "dst", dst), nil))
		return
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		fileReader, err := file.Open()
		if err != nil {
			EHWarn(err, "FileUnzip - Unable to open file", LogMap(u.Map("path", path, "dst", dst), nil))
			return
		}
		defer fileReader.Close()

		targetFile := PathJoin(dst, file.Name)
		targetFileDir := PathDirname(targetFile)
		if !FileIsDir(targetFileDir) {
			FileMkdir(targetFileDir)
		}

		fileWriter, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			EHWarn(err, "FileUnzip - Unable to create file", LogMap(u.Map("path", path, "dst", dst), nil))
			return
		}
		defer fileWriter.Close()

		_, err = io.Copy(fileWriter, fileReader)
		if err != nil {
			EHWarn(err, "FileUnzip - Unable to copy file", LogMap(u.Map("path", path, "dst", dst), nil))
			return
		}
		LTP("FileUnzip - Unzipping file", LogMap(u.Map("path", path, "dst", dst), nil))
	}
}
