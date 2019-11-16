package LennyPackagingScript

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	shareDisk_UserName = "dc_app"
	shareDisk_Password = "baidu.app"
	shareDisk_Domain = "192.168.56.56"
	shareDisk_Name = "apps"
)

func RemoteTransfer(info BuildInfo)  {
	err := execCommand("cd", "/Volumes/" + shareDisk_Name)
	if err != nil {
		err = execCommand("mkdir", "/Volumes/" + shareDisk_Name)
		if err != nil {
			panic(err)
		}
		fmt.Println("创建目录成功....-> /Volumes/shareDisk")
	}
	// 目录存在, 开始挂载远程共享盘
	err = execCommand("mount", "-t", "smbfs",
		"//" + shareDisk_UserName + ":" + shareDisk_Password + "@" + shareDisk_Domain + "/apps",
		"/Volumes/" + shareDisk_Name)
	fmt.Println("共享盘挂载成功")
	//切换目录
	err = os.Chdir("/Volumes/" + shareDisk_Name)
	if err != nil {
		panic(err)
	}
	fmt.Println("切换目录.../")
	fmt.Println(os.Getwd())
	//execCommand("ls", "-a")
	//遍历文件目录
	files, _ := ioutil.ReadDir("./")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func execCommand(name string, arg ...string) error {
	var args string
	for _, item := range arg {
		if len(arg) > 1 { args = args + item + " "
		}else { args = args + item }
	}
	fmt.Println(name, args)
	cmd := exec.Command(name, arg...)
	//cmd := &exec.Cmd{
	//	Path:name,
	//	Args:append([]string{name}, arg...),
	//}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		c := color.New(color.FgHiRed)
		c.Println(name, " failed with %s ", err)
		os.Exit(0)
		//log.Fatal(name, " failed with %s ", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Println(outStr, errStr)
	return err
}