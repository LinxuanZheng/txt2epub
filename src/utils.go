package src

import (
	"fmt"
	"github.com/LinxuanZheng/customZip"
	"io"
	"io/ioutil"
	"os"
	"path"
)

const TXTSuffix = ".txt"

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func ZipCompress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		if path.Base(file.Name()) == "mimetype" {
			err := compress(file, "", w, zip.Store)
			if err != nil {
				return err
			}
		} else {
			err := compress(file, "", w, zip.Deflate)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer, method uint16) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		if prefix != "" {
			prefix = prefix + "/" + info.Name()
		} else {
			prefix = info.Name()
		}
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw, method)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		if prefix != "" {
			header.Name = prefix + "/" + header.Name
		}
		if err != nil {
			return err
		}
		header.Method = method

		if header.Name == "mimetype" {
			writer, err := zw.CreateHeaderIgnoringExtra(header)
			if err != nil {
				return err
			}
			_, err = io.Copy(writer, file)
			file.Close()
			if err != nil {
				return err
			}
		} else {
			writer, err := zw.CreateHeader(header)
			if err != nil {
				return err
			}
			_, err = io.Copy(writer, file)
			file.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			//f, d, err := GetFilesAndDirs(dirPth + PthSep + fi.Name())
			//if err != nil {
			//	return nil, nil, err
			//}
			//files = append(files, f...)
			//dirs = append(dirs, d...)
		} else {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, dirs, nil
}

func GetTxtFile() (file string, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir, err := ioutil.ReadDir(wd)
	if err != nil {
		return "", err
	}

	for _, fi := range dir {
		if !fi.IsDir() { // 目录, 递归遍历
			if path.Ext(fi.Name()) == TXTSuffix {
				return fi.Name(), nil
			}
		}
	}

	return "", nil
}

func ErrorReport(a ...interface{}) {
	fmt.Println(a)
	os.Exit(0)
}
