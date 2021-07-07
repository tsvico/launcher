//go:generate goversioninfo
package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

// 运行时隐藏自己的cmd
// go build -ldflags -H=windowsgui /  go build  -ldflags="-H windowsgui"

//定义配置文件解析后的结构 //首字母需要大写
type LaunchConfig struct {
	Target  string   `json:"target"`
	WorkDir string   `json:"workDir"`
	Params  []string `json:"params"`
}

var (
	file *os.File
)

//go:embed launcher.json
var launchstr string

func main() {
	JsonParse := NewJsonConfig()
	v := LaunchConfig{}
	JsonParse.LoadConfig("launcher.json", &v)
	// 获取去除第一个的配置参数
	args := os.Args[1:]
	// fmt.Println(args)
	_, err := os.Stat(".lock.loop")
	if err == nil || os.IsExist(err) {
		file, _ = os.OpenFile("README.Use.txt", os.O_CREATE|os.O_WRONLY, 0666)
		file.WriteAt([]byte("u must shutdown the exe and then remove .lock.loop file"), 0)
		file.Close()
		fmt.Println("u must shutdown the exe and then remove .lock.loop file")
		return
	}
	// 加锁
	file, _ = os.Create(".lock.loop")
	// 将命令行参数 合并到配置文件的参数后面
	cmd := exec.Command(v.Target, append(v.Params, args...)...)
	cmd.Dir = v.WorkDir
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	data, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	cmd.Start()
	cmd.Wait()

	file.Close()
	os.Remove(".lock.loop")
}

/**
 * JSON 对象
 */
type JsonStruct struct {
}

func NewJsonConfig() *JsonStruct {
	return &JsonStruct{}
}

/**
* 加载配置文件
 */
func (jst *JsonStruct) LoadConfig(filename string, v interface{}) {
	_, err := os.Stat(filename)
	var data []byte
	if err == nil {
		// ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
		data, err = ioutil.ReadFile(filename)
		if err != nil {
			return
		}
	}

	// 外部配置文件不存在
	if os.IsNotExist(err) {
		data = []byte(launchstr)
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
