package pkg

import (
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

// MkDir 创建目录
func Mkdir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm) //生成多级目录
}

type File struct {
	Name         string
	Data         []byte
	ForceReplace bool
}

func Writes(dst string, files []File) {
	for _, file := range files {
		if file.Name == "" {
			continue
		}

		file.Name = dst + "/" + file.Name
		if _, err := os.Stat(file.Name); err != nil || file.ForceReplace { // 目标文件不存在或者强制覆盖
			if err := Write(file); err != nil {
				glog.Error(err)
				continue
			}
		} else {
			current, err := os.ReadFile(file.Name)
			if err != nil {
				glog.Error(err)
				continue
			}
			edits := myers.ComputeEdits(span.URIFromPath(file.Name), string(current), string(file.Data))
			diffs := fmt.Sprint(gotextdiff.ToUnified(file.Name, file.Name, string(current), edits))
			if diffs != "" {
				fmt.Printf("\n\x1b[33m%s\x1b[0m\n%s\x1b[33m%s\x1b[0m\n", ">>>>>>>>>>", diffs, "<<<<<<<<<<")
			}
		}
	}
}

// Writer 写入文件
func Write(file File) error {
	f, err := os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(file.Data)
	return err
}
