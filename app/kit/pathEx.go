package kit

import "os"

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) bool {
	//os.Stat获取文件信息
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}
	return true
}

// 判断所给路径是否为文件夹
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func PathIsFile(path string) bool {
	return !PathIsDir(path)
}
