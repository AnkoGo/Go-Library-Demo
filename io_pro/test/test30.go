package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const (
	LOGDIR  = "logs"
	LOGFILE = "oradev.log"
)

//别人的代码，不用理会这个文件，也不要执行
type LogStruct struct {
	LogDirFilePath string
	file           *os.File
}
type LogInterface interface {
	InitLog() error
	CloseLog()
	Write()
}

//初始化日志系统接口
func (l *LogStruct) InitLog() error {
	fmt.Printf("Start adlog system initialization...\n\n")

	// 检查日志文件是否存在,
	res, err := FileCreate(LOGDIR, LOGFILE)
	if err != nil {
		return err
	}
	if res != nil {
		l.LogDirFilePath = *res
		logFile, err := os.OpenFile(l.LogDirFilePath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Log system initialization failed: %s\n", err)
		}
		l.file = logFile
		fmt.Printf("Log system initialization complete.\n\n")
	}
	return nil
}
func (l *LogStruct) Write(tp string, info interface{}) {
	fmt.Println("lfile: ", l.file)
	fmt.Println("path: ", l.LogDirFilePath)
	adlog := log.New(l.file, "", log.Ldate|log.Ltime)
	if tp == "INFO: " {
		adlog.Println(tp, info)
	}
	if tp == "WARNING: " {
		adlog.Println(tp, info)
	}
	if tp == "ERROR: " {
		adlog.Println(tp, info)
	}
	if tp == "FATAL: " {
		adlog.Println(tp, info)
		os.Exit(1)
	}
}
func INFO(info interface{}) error {
	l := &LogStruct{}
	l.Write("INFO: ", info)
	return fmt.Errorf("++++++++++++++++++++++++++++++++++++++++++++++++++")
}
func WARNING(info interface{}) {
	l := LogStruct{}
	l.Write("WARNING: ", info)
}
func ERROR(info interface{}) {
	l := LogStruct{}
	l.Write("ERROR: ", info)
}
func FATAL(info interface{}) {
	l := LogStruct{}
	l.Write("FATAL: ", info)
}

//关闭日志文件
func (l *LogStruct) CloseLog() {
	if l.file != nil {
		l.file.Close()
	} else {
		fmt.Println("close log: ", l.file)
	}
}

//关闭日志文件
func CLog() {
	l := LogStruct{}
	l.CloseLog()
}

// directory and file create
func FileCreate(dir string, filename string) (*string, error) {
	var pfres string

	//homedir := GetHomeDirectory()
	homedir := `C:\Users\Administrator\Desktop\go_pro\src\io_pro\test`
	if homedir != "" {
		exist, err := PathExists(path.Join(homedir, dir))
		if err != nil {
			return nil, fmt.Errorf("Get directory property information occurred Error: %v\n", err)
		}
		if exist {
			fmt.Printf("Check if directory exists - exists: %s\n\n", path.Join(homedir, dir))
		} else {
			fmt.Printf("Check if directory exists - Non-exists: %s\n", path.Join(homedir, dir))
			// 创建文件夹
			err := os.Mkdir(path.Join(homedir, dir), os.ModePerm)
			if err != nil {
				return nil, fmt.Errorf("directory created failed: %s\n\n", err)
			} else {
				fmt.Printf("directory created successfully: %s\n\n", path.Join(homedir, dir))
			}
		}
	} else {
		return nil, fmt.Errorf("Program home directory acquisition failed\n\n")
	}
	if filename != "" {
		exist, err := PathExists(filepath.Join(path.Join(homedir, dir), filename))
		if err != nil {
			return nil, fmt.Errorf("Get file property information occurred Error: %v\n", err)
		}
		if exist {
			fmt.Printf("Check if file exists - exists: %s\n", filepath.Join(path.Join(homedir, dir), filename))
			pfres = filepath.Join(path.Join(homedir, dir), filename)
			return &pfres, nil
		} else {
			fmt.Printf("Check if file exists - Non-exists: %s\n\n", filepath.Join(path.Join(homedir, dir), filename))
			_, err := os.OpenFile(filepath.Join(path.Join(homedir, dir), filename), os.O_EXCL|os.O_CREATE, os.ModePerm)
			if err != nil {
				return nil, fmt.Errorf("file creation failed: %s\n\n", err)
			} else {
				fmt.Printf("file created successfully: %s\n\n", filepath.Join(path.Join(homedir, dir), filename))
				pfres = filepath.Join(path.Join(homedir, dir), filename)
				return &pfres, nil
			}
		}
	} else {
		return nil, fmt.Errorf("File name cannot be empty")
	}
	return nil, nil
}

//Get program home directory
func GetHomeDirectory() (homedir string) {
	file, _ := exec.LookPath(os.Args[0])
	ExecFilePath, _ := filepath.Abs(file)
	execfileslice := strings.Split(ExecFilePath, "/")
	HomeDirectory := execfileslice[:len(execfileslice)-2]
	for _, v := range HomeDirectory {
		if v != "" {
			homedir += `/` + v
		}
	}
	if homedir != "" {
		fmt.Printf("Program home directory initialized successfully: %s\n", homedir)
		return homedir
	} else {
		fmt.Printf("Program home directory initialized failed\n")
	}
	return ""
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main8938373() {
	l := LogStruct{}
	err := (&l).InitLog()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := INFO(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("r:", r)

}
