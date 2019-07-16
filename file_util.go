package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

//复制文件
func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return -1, err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return -1, err
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)

}

//文件是否存在
func IsExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//判断目录是否存在
func IsDirExists(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

//检查目录是否存在，如果不存在则创建目录
func CheckAndCreateDir(dirname string, mode os.FileMode) error {
	ok := IsDirExists(dirname)
	if !ok {
		err := os.MkdirAll(dirname, mode)
		if err != nil {
			return fmt.Errorf("%s:make dir err!", dirname)
		}
	}
	return nil
}

//检查文件是否存在，如果不存在则创建
func CheckAndCreateFile(filename string, mode os.FileMode) error {
	if IsExists(filename) {
		return nil
	}
	pathFile := path.Dir(filename)
	if err := CheckAndCreateDir(pathFile, mode); err != nil {
		return err
	}

	fi, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fi.Close()
	return nil
}

func GetFileString(filename string) (content string, error error) {
	fi, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return "", err
	}
	return string(fd), nil
}

func GetBasePath() string {
	var dirAbsPath string
	ex, err := os.Executable()
	if err == nil {
		dirAbsPath = filepath.Dir(ex)
	} else {
		exReal, err := filepath.EvalSymlinks(ex)
		if err != nil {
			panic(err)
		}
		dirAbsPath = filepath.Dir(exReal)
	}
	dirAbsPath = filepath.Join(dirAbsPath, "..")
	return filepath.Clean(dirAbsPath)
}
