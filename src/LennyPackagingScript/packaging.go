package LennyPackagingScript

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"os/user"
	"reflect"
	"strings"
)

var (
	Info_Build = BuildInfo{}
	Targets = []string{
		"XY000T",
		"XY001", "XY002", "XY003", "XY005", "XY006", "XY007", "XY008", "XY009", "XY010", "XY011", "XY012", "XY013",
		"XY015", "XY016", "XY017", "XY018", "XY019", "XY020", "XY021", "XY022", "XY023", "XY025", "XY026", "XY027", "XY028",
		"XY029", "XY030", "XY031", "XY032", "XY033", "XY035", "XY036", "XY037", "XY039", "XY050", "XY051", "XY052",
		"XY053", "XY055", "XY056", "XY057", "XY058", "XY059", "XY060", "XY061", "XY062", "XY063", "XY065", "XY066", "XY067",
		"XY068", "XY069", "XY070", "XY071", "XY072", "XY073", "XY075", "XY076", "XY077", "XY078", "XY079", "XY080", "XY081",
		"XY082", "XY083", "XY085", "XY086", "XY087", "XY088", "XY089", "XY090", "XY091", "XY092", "XY093", "XY095", "XY096",
		"XY097", "XY098", "XY099", "XY100", "XY101", "XY102", "XY103", "XY105", "XY106", "XY107", "XY108", "XY109", "XY110",
		"XY111", "XY112", "XY113", "XY115", "XY116", "XY117", "XY118", "XY119", "XY120", "XY121", "XY122", "XY123", "XY125",
		"XY126", "XY127", "XY128", "XY129", "XY130", "XY131", "XY132", "XY133", "XY135", "XY136", "XY137", "XY138", "XY139",
		"XY150", "XY151", "XY152", "XY153", "XY155", "XY156", "XY157", "XY158", "XY159", "XY160", "XY161", "XY162", "XY163",
		"XY165", "XY166", "XY167", "XY168", "XY169", "XY170", "XY171", "XY172", "XY173", "XY175", "XY176", "XY177", "XY178",
		"XY179", "XY180", "XY181", "XY182", "XY183", "XY185", "XY186", "XY187",
	}
	userName = ""
	applicationConfigInfo = ApplicationConfigInfo{}
)
type BuildInfo struct {
	Scheme		string `json:scheme`
	Path 		string `json:"path"`
	WorkSpace	string `json:workspace`
	ArchPath	string `json:"archPath"`
	ExportPath	string `json:"exportPath"`
	Config 		string `json:"config"`
	ApiToken 	string `json:"apiToken"`
	TargetName	string `json:"target_name"`
	TargetIndex	int		`json:"target_index"`
}

func init() {
	//工程的 Target
	Info_Build.Scheme = "HBVertical"   //默认
	//工程文件的路径  --> *.xcodeproj
	Info_Build.Path  = "/Users/lenny/XcodeProjects/a_ios"
	//工程文件的名字  --> *.xcodeproj, 或 *.xcworkspace
	Info_Build.WorkSpace = "Convert"
	//保存 archive 文件的路径
	Info_Build.ArchPath = "/Users//Desktop/archiveDirectory"
	Info_Build.ExportPath = "/Users/lenny/Desktop/exportDirectory"
	path, _ := os.Getwd()
	inside_CurrentPath = path
	u, _ := user.Current()
	userName = u.Username
	Info_Build.Path = path
	Info_Build.ArchPath = "/Users/" + userName + "/Desktop/archiveDirectory"
	Info_Build.ExportPath = "/Users/" + userName + "/Desktop/exportDirectory"
	Info_Build.TargetName = "XY000T"
	applicationConfigInfo.Application_TargetName = ""
	applicationConfigInfo.Application_DisplayName = ""
	applicationConfigInfo.Application_BuildVersion = ""
	applicationConfigInfo.Application_Version = ""
	applicationConfigInfo.Application_BundleIdentifier = ""
}


func archive(info BuildInfo) error {
	fmt.Println("Target: -->", info.Scheme)
	fmt.Println("Path: -->", info.Path)
	fmt.Println("WorkSpace: -->", info.WorkSpace)
	fmt.Println("ArchPath: -->", info.ArchPath)
	fmt.Println("ExportPath: -->", info.ExportPath)

	err := os.Chdir(info.Path)
	if err != nil {
		panic(err.Error())
	}
	_, err = os.Stat(info.ArchPath)
	if err != nil && os.IsNotExist(err) {
		//文件夹不存在
		err = os.Mkdir(info.ArchPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	_, err = os.Stat(info.ExportPath)
	if err != nil && os.IsNotExist(err) {
		//文件夹不存在
		err = os.Mkdir(info.ExportPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	workspace := info.WorkSpace + ".xcworkspace"
	archCommand := exec.Command("xcodebuild", "archive",
		"-workspace", workspace,
		"-scheme", info.Scheme,
		"-archivePath", info.ArchPath + "/" + info.TargetName,
		"-configuration", "Release")
		//"ONLY_ACTIVE_ARCH=NO",
		//"CODE_SIGN_IDENTITY=" + "iPhone Developer",
		//"PROVISIONING_PROFILE_SPECIFIER=" + "Automatic",
		//"PROVISIONING_PROFILE=" + "Automatic",
		//"CODE_SIGN_STYLE=" + "Automatic",
		//"PROVISIONING_STYLE=" + "Automatic",
		//"DEVELOPMENT_TEAM=" + codeSignTeamIdentifier())
	archCommand.Stdout = os.Stdout
	fmt.Println("Method: --> archive Command -->\n", archCommand)
	err  = archCommand.Run()
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func export(info BuildInfo) error {
	fmt.Println("Method: -->", "export")
	arch := info.ArchPath + "/" + info.TargetName + ".xcarchive"
	fmt.Println("Method: --> arch -->", arch)
	export := exec.Command("xcodebuild", "-exportArchive",
		"-archivePath", arch,
		"-configuration", "Release",
		"-exportPath", info.ExportPath + "/" + info.TargetName,
		"-exportOptionsPlist", info.Path + "/" + "ExportOptions.plist",
		"-allowProvisioningUpdates", "-quiet")
	export.Stdout = os.Stdout
	fmt.Println("Method: --> export Command -->\n", export)
	err := export.Run()
	if err != nil {
		panic(err.Error())
	}else {
		c := color.New(color.FgHiRed)
		_, _ = c.Println("Export ", info.Scheme, " Success!!!")
	}
	return nil
}

func PackagingFrom(target string) {
	for index, item := range Targets {
		if item == strings.ToUpper(target) {
			Info_Build.TargetIndex = index
			break
		}
	}
	for ; Info_Build.TargetIndex <= len(Targets); Info_Build.TargetIndex++ {
		Info_Build.TargetName = Targets[Info_Build.TargetIndex]
		if changeTarget_By_ModifiFieldContent(Info_Build) != nil {
			fmt.Println("修改 Plist 文件出错...!!!")
			os.Exit(0)
		}
		if changeTargetAppIcon(Info_Build) != nil {
			fmt.Println("修改 AppIcon 出错...!!!")
			os.Exit(0)
		}
		if changeXcodeProj_pbxproj(Info_Build) != nil {
			fmt.Println("切换账号出错-->..!!!")
			os.Exit(0)
		}
		if archive(Info_Build) != nil {
			fmt.Println("Archive 出错-->.!!!")
			os.Exit(0)
		}
		if export(Info_Build) != nil {
			fmt.Println("导出 IPA 出错-->. !!!")
		}
		if RemoteTransfer(Info_Build) != nil {
			fmt.Println("上传到共享盘失败...!!!")
			os.Exit(0)
		}
	}
}
func Packaging(from string, to string) {
	for index, item := range Targets {
		if item == strings.ToUpper(from) {
			Info_Build.TargetIndex = index
			break
		}
	}
	var index_To int
	for index, item := range Targets {
		if item == strings.ToUpper(to) {
			index_To = index
			break
		}
	}
	for ; Info_Build.TargetIndex <= index_To; Info_Build.TargetIndex++ {
		Info_Build.TargetName = Targets[Info_Build.TargetIndex]
		if changeTarget_By_ModifiFieldContent(Info_Build) != nil {
			fmt.Println("修改 Plist 文件出错...!!!")
			os.Exit(0)
		}
		if changeTargetAppIcon(Info_Build) != nil {
			fmt.Println("修改 AppIcon 出错...!!!")
			os.Exit(0)
		}
		if changeXcodeProj_pbxproj(Info_Build) != nil {
			fmt.Println("切换账号出错-->..!!!")
			os.Exit(0)
		}
		if archive(Info_Build) != nil {
			fmt.Println("Archive 出错-->.!!!")
			os.Exit(0)
		}
		if export(Info_Build) != nil {
			fmt.Println("导出 IPA 出错-->. !!!")
		}
		if RemoteTransfer(Info_Build) != nil {
			fmt.Println("上传到共享盘失败...!!!")
			os.Exit(0)
		}
	}
}

func PackagingTarget(target string) {
	for index, item := range Targets {
		if item == strings.ToUpper(target) {
			Info_Build.TargetIndex = index
			break
		}
	}
	Info_Build.TargetName = Targets[Info_Build.TargetIndex]
	if changeTarget_By_ModifiFieldContent(Info_Build) != nil {
		fmt.Println("修改 Plist 文件出错...!!!")
		os.Exit(0)
	}
	if changeTargetAppIcon(Info_Build) != nil {
		fmt.Println("修改 AppIcon 出错...!!!")
		os.Exit(0)
	}
	if changeXcodeProj_pbxproj(Info_Build) != nil {
		fmt.Println("切换账号出错-->..!!!")
		os.Exit(0)
	}
	if changeTargetExportOptionsFile(Info_Build) != nil {
		//
	}
	if archive(Info_Build) != nil {
		fmt.Println("Archive 出错-->.!!!")
		os.Exit(0)
	}
	if export(Info_Build) != nil {
		fmt.Println("导出 IPA 出错-->. !!!")
	}
	if RemoteTransfer(Info_Build) != nil {
		fmt.Println("上传到共享盘失败...!!!")
		os.Exit(0)
	}
}

func codeSignTeamIdentifier(info BuildInfo) string {

	var container1 = []string{"XY000T", "XY001", "XY002", "XY003", "XY005", "XY006", "XY007", "XY008", "XY009", "XY010"}
	var container2 = []string{"XY011", "XY012", "XY013", "XY015", "XY016", "XY017", "XY018", "XY019", "XY020", "XY021", "XY022", "XY023", "XY025", "XY026", "XY027", "XY028", "XY029", "XY030"}
	var container3 = []string{"XY031", "XY032", "XY033", "XY035", "XY036", "XY037", "XY039", "XY050", "XY051", "XY052", "XY053", "XY055", "XY056", "XY057", "XY058", "XY059", "XY060"}
	var container4 = []string{"XY061", "XY062", "XY063", "XY065", "XY066", "XY067", "XY068", "XY069", "XY070", "XY071", "XY072", "XY073", "XY075", "XY076", "XY077", "XY078", "XY079", "XY080"}
	var container5 = []string{"XY081", "XY082", "XY083", "XY085", "XY086", "XY087", "XY088", "XY089", "XY090", "XY091", "XY092", "XY093", "XY095", "XY096", "XY097", "XY098", "XY099", "XY100"}
	var container6 = []string{"XY101", "XY102", "XY103", "XY105", "XY106", "XY107", "XY108", "XY109", "XY110", "XY111", "XY112", "XY113", "XY115", "XY116", "XY117", "XY118", "XY119", "XY120"}
	var container7 = []string{"XY121", "XY122", "XY123", "XY125", "XY126", "XY127", "XY128", "XY129", "XY130", "XY131", "XY132", "XY133", "XY135", "XY136", "XY137", "XY138", "XY139"}
	var container8 = []string{"XY150", "XY151", "XY152", "XY153", "XY155", "XY156", "XY157", "XY158", "XY159", "XY160", "XY161", "XY162", "XY163", "XY165", "XY166", "XY167", "XY168", "XY169", "XY170"}
	var container9 = []string{"XY171", "XY172", "XY173", "XY175", "XY176", "XY177", "XY178", "XY179", "XY180", "XY181", "XY182", "XY183", "XY185", "XY186", "XY187", "XY188", "XY189", "XY190"}
	if contains(container1, info.TargetName) >= 0 {
		return "2U94TLD4SJ"
	}else if contains(container2, info.TargetName) >= 0 {
		return "93Y4U3V625"
	}else if contains(container3, info.TargetName) >= 0 {
		return "CV4VR2BK32"
	}else if contains(container4, info.TargetName) >= 0 {
		return "8X36854WB2"
	}else if contains(container5, info.TargetName) >= 0 {
		return "3D37BN9GYF"
	}else if contains(container6, info.TargetName) >= 0 {
		return "P7M4387Y6U"
	}else if contains(container7, info.TargetName) >= 0 {
		return "K435P4QZ9J"
	}else if contains(container8, info.TargetName) >= 0 {
		return "PHS6ACGVHM"
	}else if contains(container9, info.TargetName) >= 0 {
		//return ""
		fmt.Println("盘口对应的开发者账号没有添加....!!!")
		os.Exit(0)
	}
	return ""
}

func contains(array interface{}, args interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(array)
		for i := 0; i < a.Len(); i++ {
			if reflect.DeepEqual(args, a.Index(i).Interface()) {
				index = i
				return
			}
		}
	}
	return
}
