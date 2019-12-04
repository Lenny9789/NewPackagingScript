package LennyPackagingScript

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	shareDisk_UserName = "dc_app"
	shareDisk_Password = "baidu.app"
	shareDisk_Domain = "192.168.56.56"
	shareDisk_Name = "apps"
)

func RemoteTransfer(info BuildInfo) error {
	if isExist("/Users/" + userName + "/Desktop/" + shareDisk_Name) == false {
		err := execCommand("mkdir", "/Users/" + userName + "/Desktop/" + shareDisk_Name)
		if err != nil {
			panic(err)
		}
		fmt.Println("创建目录成功....-> ")
	}
	// 目录存在, 开始挂载远程共享盘
	//先 删除, 再重新加载
	files, _ := ioutil.ReadDir("/Users/" + userName + "/Desktop/" + shareDisk_Name)
	if len(files) > 0 {
		if execCommand("umount", "/Users/" + userName + "/Desktop/" + shareDisk_Name) == nil {
			fmt.Println("清除本地挂载成功...!")
		}
	}

	if execCommand("mount", "-t", "smbfs",
		"//" + shareDisk_UserName + ":" + shareDisk_Password + "@" + shareDisk_Domain + "/apps",
		"/Users/" + userName + "/Desktop/" + shareDisk_Name) == nil {
		fmt.Println("重新挂载成功....!")
	}

	//切换目录
	err := os.Chdir("/Users/" + userName + "/Desktop/" + shareDisk_Name)
	if err != nil {
		panic(err)
	}
	fmt.Println("切换目录.../")
	fmt.Println(os.Getwd())
	//遍历文件目录
	files, _ = ioutil.ReadDir("./")
	var file_Index int
	for index, file := range files {
		fmt.Println(file.Name())
		if strings.HasSuffix(file.Name(), info.TargetName[2:]) {
			file_Index = index
			break
		}
	}
	if file_Index == 0 {
		fmt.Println("未找到匹配的文件夹路径...!!!")
		if info.TargetName != "XY000T" {
			os.Exit(0)
		}
		fmt.Println("默认 X0Test ....!")
	}
	path_From := "/Users/" + userName + "/Desktop/exportDirectory/" + info.TargetName + "/HBVertical.ipa"
	path_To := "/Users/" + userName + "/Desktop/" + shareDisk_Name + "/" + files[file_Index].Name() + "/ios_ipa/" + info.TargetName + ".ipa"
	if info.TargetName == "XY000T" {
		path_To = "/Users/" + userName + "/Desktop/" + shareDisk_Name + "/xy00000" + "/ios_ipa/" + info.TargetName + ".ipa"
	}
	if execCommand("cp", path_From, path_To) == nil {
		fmt.Println("上传到共享盘成功...!!!")
	}
	return nil
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

func notifi(subTitle string)  {
	//osascript -e 'display notification "Hello world!" with title "Hi!"'
	execCommand("osascript", "-e", )
}