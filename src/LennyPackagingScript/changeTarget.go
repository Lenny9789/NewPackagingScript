package LennyPackagingScript

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DHowett/go-plist"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
)

var (
	input_Json_Config_Path = "/Users/lenny/Desktop/a_ios/FEC/targetConfig.json"
	input_Plist_Config_Path = "/Users/lenny/Desktop/a_ios/FEC/Info/"
	input_Target_Name = "XY000t"
	inside_CurrentPath = ""
)

type ApplicationBundleInfo struct {
	LennyBundleTargetName 	string `plist:"LennyBundleTargetName"`
	CFBundleDevelopmentRegion string `plist:"CFBundleDevelopmentRegion"`
	CFBundleDisplayName string `plist:"CFBundleDisplayName"`
	CFBundleExecutable  string `plist:"CFBundleExecutable"`
	CFBundleIdentifier  string    `plist:"CFBundleIdentifier"`
	CFBundleInfoDictionaryVersion   string `CFBundleInfoDictionaryVersion"`
	CFBundleName        string `plist:"CFBundleName"`
	CFBundlePackageType string `plist:"CFBundlePackageType"`
	CFBundleShortVersionString string `plist:"CFBundleShortVersionString"`
	CFBundleSignature 	string `plist:"CFBundleSignature"`
	CFBundleURLTypes	[]map[string]interface{} `plist:"CFBundleURLTypes"`
	CFBundleVersion		string `plist:"CFBundleVersion"`
	LSApplicationCategoryType bool `plist:"LSApplicationCategoryType"`
	LSApplicationQueriesSchemes []string `plist:"LSApplicationQueriesSchemes"`
	LSRequiresIPhoneOS  bool `plist:"LSRequiresIPhoneOS"`
	NSAppTransportSecurity map[string]interface{} `plist:"NSAppTransportSecurity"`
	NSAppleMusicUsageDescription string `plist:"NSAppleMusicUsageDescription"`
	NSBluetoothPeripheralUsageDescription string `plist:"NSBluetoothPeripheralUsageDescription"`
	NSCalendarsUsageDescription string `plist:"NSCalendarsUsageDescription"`
	NSCameraUsageDescription	string `plist:"NSCameraUsageDescription"`
	NSContactsUsageDescription	string `plist:"NSContactsUsageDescription"`
	NSHealthShareUsageDescription string `plist:"NSHealthShareUsageDescription"`
	NSHealthUpdateUsageDescription	string `plist:"NSHealthUpdateUsageDescription"`
	NSLocationAlwaysAndWhenInUseUsageDescription string `plist:"NSLocationAlwaysAndWhenInUseUsageDescription"`
	NSLocationAlwaysUsageDescription string `plist:"NSLocationAlwaysUsageDescription"`
	NSLocationUsageDescription		string `plist:"NSLocationUsageDescription"`
	NSLocationWhenInUseUsageDescription string `plist:"NSLocationWhenInUseUsageDescription"`
	NSMicrophoneUsageDescription	string `plist:"NSMicrophoneUsageDescription"`
	NSMotionUsageDescription 		string `plist:"NSMotionUsageDescription"`
	NSPhotoLibraryAddUsageDescription string `plist:"NSPhotoLibraryAddUsageDescription"`
	NSPhotoLibraryUsageDescription		string `plist:"NSPhotoLibraryUsageDescription"`
	NSRemindersUsageDescription			string `plist:"NSRemindersUsageDescription"`
	UIBackgroundModes				[]string	`plist:"UIBackgroundModes"`
	UIRequiredDeviceCapabilities	[]string	`plist:"UIRequiredDeviceCapabilities"`
	UIStatusBarHidden				bool `plist:"UIStatusBarHidden"`
	UIStatusBarStyle				string `plist:"UIStatusBarStyle"`
	UISupportedInterfaceOrientations	[]string `plist:"UISupportedInterfaceOrientations"`
	UIViewControllerBasedStatusBarAppearance bool `plist:"UIViewControllerBasedStatusBarAppearance"`
	Com_openinstall_APP_KEY		string `plist:"com.openinstall.APP_KEY"`
}
type ApplicationConfigInfo struct {
	Application_TargetName		string `json:"TargetName"`
	Application_DisplayName	    string	`json:"DisplayName"`
	Application_BundleIdentifier    string	`json:"BundleIdentifier"`
	Application_Version		string 	 `json:"Version"`
	Application_BuildVersion	string	`json:"Build"`
}
type ApplicationConfigs struct {
	Configs 	[]ApplicationConfigInfo `json:"Configs"`
}

type AppIconContentConfig struct {
	Size	string `json:"size"`
	Idiom	string `json:"idiom"`
	FileName 	string `json:"filename"`
	Scale 		string `json:"scale"`
}
type AppIconConfig struct {
	Images	[]AppIconContentConfig `json:"images"`
}

type LaunchImageContentConfig struct {
	Extent	string `json:"extent"`
	Idiom 	string `json:"idiom"`
	Subtype 	string `json:"subtype"`
	FileName	string `json:"filename"`
	MinimumSystemVersion	string `json:"minimum-system-version"`
	Orientation		string `json:"orientation"`
	Scale 		string `json:"scale"`
}
type LaunchImageConfig struct {
	Images []LaunchImageContentConfig `json:"images"`
}
func changeTarget_By_InsteadFile(target string) error {
	var index_Targets int
	for index, item := range Targets {
		if item == strings.ToUpper(target) {
			index_Targets = index
			break
		}
	}
	files, _ := ioutil.ReadDir(inside_CurrentPath + "/InfoFiles")
	var theFileExist  = false
	var index_Files int
	for index, file := range files {
		fmt.Println(file.Name(), index_Targets)
		if strings.Contains(file.Name(), Targets[index_Targets]) {
			theFileExist = true
			index_Files = index
			break
		}
	}
	if theFileExist {
		//先删掉原来的 plist, 再把新的拷贝过去
		if isExist(inside_CurrentPath + "/FEC/Info/HBVertical.plist") {
			_ = execCommand("rm",  "./FEC/Info/HBVertical.plist")
		}
		if execCommand("cp", "./InfoFiles/" + files[index_Files].Name(), "./FEC/Info/HBVertical.plist") == nil {
			//if execCommand("mv", inside_CurrentPath + "/FEC/Info/" + files[index_Files].Name(), inside_CurrentPath + "/FEC/Info/HBVertical.plist") == nil {
				fmt.Println("Plist 文件替换成功....!!!!")
			//}
		}
	}
	return nil
}

func changeTargetAppIcon(info BuildInfo) error {

	var index_Assets int
	files_Assets, _ := ioutil.ReadDir("./FEC/AssetSource/")
	for index, file := range files_Assets {
		fmt.Println(file.Name())
		if file.Name() == Targets[info.TargetIndex] + ".xcassets" {
			index_Assets = index
			break
		}
	}
	files_PNG, _ := ioutil.ReadDir("./FEC/AssetSource/" + files_Assets[index_Assets].Name())
	for _, file_imageType := range files_PNG {
		fmt.Println(file_imageType.Name())
		if strings.Contains(file_imageType.Name(), "Image.imageset") {
			//把 shareImage 复制过去
			fmt.Println("修改 shareImage")
			file_ShareImages, _ := ioutil.ReadDir("./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name())
			for _, file := range file_ShareImages {
				fmt.Println(file.Name())
				if strings.Contains(file.Name(), ".png") {
					if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + file.Name(), "./FEC/AssetSource/UniqueAppIcon.xcassets/shareImage.imageset/shareImage.png") == nil {
						fmt.Println("修改 shareImage 成功!!!")
					}
				}
			}
		}
		if strings.Contains(file_imageType.Name(), ".appiconset") {
			fmt.Println("修改 Appicon...")
			//file_AppIcon, _ := ioutil.ReadDir("./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name())
			if isExist("./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/Contents.json") {
				// 读取Json配置文件
				j, err := ioutil.ReadFile("./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/Contents.json")
				if err != nil {
					c := color.New(color.FgHiRed)
					_, _ = c.Println("读取 json 配置文件失败.  ", err)
					os.Exit(0)
				}
				fmt.Println(string(j))
				json_Data := &AppIconConfig{}
				err = json.Unmarshal([]byte(j), json_Data)
				if err != nil {
					c := color.New(color.FgHiRed)
					_, _ = c.Println("json 配置文件格式错误.  ", err)
					os.Exit(0)
				}
				for _, config := range json_Data.Images {
					switch config.Size {
					case "20x20":
						if config.Scale  == "2x" && len(config.FileName) > 0 {
							if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/AppIcon.appiconset/appicon_20pt@2x.png") == nil {
								fmt.Println("修改 Appicon-> 20pt -> 2x  成功!!!")
							}
						}
						if config.Scale  == "3x" && len(config.FileName) > 0 {
							if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/AppIcon.appiconset/appicon_20pt@3x.png") == nil {
								fmt.Println("修改 Appicon-> 20pt -> 3x  成功!!!")
							}
						}
					case "29x29":
						if config.Scale  == "2x" && len(config.FileName) > 0 {
							if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/AppIcon.appiconset/appicon_29pt@2x.png") == nil {
								fmt.Println("修改 Appicon-> 29pt -> 2x  成功!!!")
							}
						}
						if config.Scale  == "3x" && len(config.FileName) > 0 {
							if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/AppIcon.appiconset/appicon_29pt@3x.png") == nil {
								fmt.Println("修改 Appicon-> 29pt -> 3x  成功!!!")
							}
						}
					case "40x40":
						if config.Scale  == "2x" && len(config.FileName) > 0 {
							if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/AppIcon.appiconset/appicon_40pt@2x.png") == nil {
								fmt.Println("修改 Appicon-> 40pt -> 2x  成功!!!")
							}
						}
						if config.Scale  == "3x"  && len(config.FileName) > 0 {
							if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/AppIcon.appiconset/appicon_40pt@3x.png") == nil {
								fmt.Println("修改 Appicon-> 40pt -> 3x  成功!!!")
							}
						}
					case "60x60":
						if config.Scale  == "2x" && len(config.FileName) > 0 {
							if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/AppIcon.appiconset/appicon_60pt@2x.png") == nil {
								fmt.Println("修改 Appicon-> 60pt -> 2x  成功!!!")
							}
						}
						if config.Scale  == "3x" && len(config.FileName) > 0 {
							if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/AppIcon.appiconset/appicon_60pt@3x.png") == nil {
								fmt.Println("修改 Appicon-> 60pt -> 3x  成功!!!")
							}
						}
					case "1024x1024":
						if config.Scale  == "1x" && len(config.FileName) > 0 {
							if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/AppIcon.appiconset/appicon_1024*1024.png") == nil {
								fmt.Println("修改 Appicon-> 1024   成功!!!")
							}
						}
					default:
						break
					}
				}
			} else {
				c := color.New(color.FgHiRed)
				_, _ = c.Println("json 配置文件不存在 !!!.  ")
				os.Exit(0)
			}
		}
		if strings.Contains(file_imageType.Name(), ".launchimage") {
			if isExist("./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/Contents.json") {
				// 读取Json配置文件
				j, err := ioutil.ReadFile("./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/Contents.json")
				if err != nil {
					c := color.New(color.FgHiRed)
					_, _ = c.Println("读取 json 配置文件失败.  ", err)
					os.Exit(0)
				}
				fmt.Println(string(j))
				json_Data := &LaunchImageConfig{}
				err = json.Unmarshal([]byte(j), json_Data)
				if err != nil {
					c := color.New(color.FgHiRed)
					_, _ = c.Println("json 配置文件格式错误.  ", err)
					os.Exit(0)
				}
				for _, config := range json_Data.Images {
					if config.Subtype == "2436h" && len(config.FileName) > 0 {
						if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/LaunchImage.launchimage/X.png") == nil {
							fmt.Println("修改 LaunchImage -> X   成功!!!")
						}
					}
					if config.Subtype == "736h" && len(config.FileName) > 0 {
						if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/LaunchImage.launchimage/launchimage_1242*2208.png") == nil {
							fmt.Println("修改 LaunchImage -> launchimage_1242*2208   成功!!!")
						}
					}
					if config.Subtype == "667h" && len(config.FileName) > 0 {
						if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/LaunchImage.launchimage/launchimage_750*1334.png") == nil {
							fmt.Println("修改 LaunchImage -> launchimage_750*1334  成功!!!")
						}
					}
					if config.Subtype == "retina4" && len(config.FileName) > 0 {
						if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/LaunchImage.launchimage/launchimage_640*1136.png") == nil {
							fmt.Println("修改 LaunchImage -> launchimage_640*1136   成功!!!")
						}
					}
					if config.Subtype == "" && len(config.FileName) > 0 {
						if execCommand("cp", "./FEC/AssetSource/" + files_Assets[index_Assets].Name() + "/" + file_imageType.Name() + "/" + config.FileName, "./FEC/AssetSource/UniqueAppIcon.xcassets/LaunchImage.launchimage/launchimage_640*960.png") == nil {
							fmt.Println("修改 LaunchImage -> launchimage_640*960   成功!!!")
						}
					}
				}
			}else {
				c := color.New(color.FgHiRed)
				_, _ = c.Println("json 配置文件不存在 !!!.  ")
				os.Exit(0)
			}
		}
	}
	return nil
}

func loadApplicationConfigInfo(info BuildInfo) error {

	j, err := ioutil.ReadFile(inside_CurrentPath + "/FEC/targetConfig.json")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(j))
	json_Data := &ApplicationConfigs{}
	err = json.Unmarshal([]byte(j), json_Data)
	if err != nil {
		fmt.Println("------------> Json 文件格式错误, 干你木.")
		panic(err)
	}
	var index_Config int
	for index, item := range json_Data.Configs {
		if item.Application_TargetName == strings.ToUpper(info.TargetName) {
			index_Config = index
		}
	}
	if strings.ToUpper(info.TargetName) != json_Data.Configs[index_Config].Application_TargetName {
		fmt.Println("Json 配置文件中没有这个盘口. 干!!!")
		os.Exit(0)
	}
	applicationConfigInfo = json_Data.Configs[index_Config]
	return nil
}
func changeTarget_By_ModifiFieldContent(info BuildInfo) error {

	//读取项目中的 Plist 文件  , 待修改
	f, err := ioutil.ReadFile(inside_CurrentPath + "/FEC/Info/HBVertical.plist")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(f))
	buf := bytes.NewReader([]byte(f))
	var data ApplicationBundleInfo
	//把读取的 Plist 格式化为结构体
	decoder := plist.NewDecoder(buf)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.CFBundleDisplayName + data.CFBundleVersion + data.CFBundleShortVersionString)
	// 读取Json配置文件, 并将配置文件中的信息写到 PList 结构体中.
	if loadApplicationConfigInfo(info) != nil {
		//
	}
	data.LennyBundleTargetName = applicationConfigInfo.Application_TargetName
	data.CFBundleDisplayName = applicationConfigInfo.Application_DisplayName
	data.CFBundleIdentifier = applicationConfigInfo.Application_BundleIdentifier
	data.CFBundleShortVersionString = applicationConfigInfo.Application_Version
	data.CFBundleVersion = applicationConfigInfo.Application_BuildVersion
	// 将已经修改完的 Plist 结构体对象转为 PList 文件并写入到工程文件中.
	bf := &bytes.Buffer{}
	encoder := plist.NewEncoder(bf)
	err = encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(bf.String())
	err = ioutil.WriteFile(inside_CurrentPath + "/FEC/Info/HBVertical.plist", []byte(bf.String()), 0777)
	if err != nil {
		fmt.Println("-------------> 写入文件失败.")
		panic(err)
	}
	fmt.Println("--------------> Plist 文件修改成功!!!")
	return nil
}

func isExist(path string) bool  {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) { return  true}
		return false
	}
	return true
}


type PbxprojConfig struct {
	Objects map[string]dict		`plist:"objects""`
	RootObject string	`plist:"rootObject"`
}
type dict map[string]interface{}

func changeXcodeProj_pbxproj(info BuildInfo) error {
	var tempMap map[string]interface{}
	var targets []interface{}
	var buildConfigurations []interface{}
	var buildSettingKey_Releasae string
	//path := "/Users/lenny/XcodeProjects/a_ios/Convert.xcodeproj/project.pbxproj"
	path := info.Path + "/Convert.xcodeproj/project.pbxproj"
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f))
	buf := bytes.NewReader([]byte(f))
	var data PbxprojConfig
	//把读取的 Plist 格式化为结构体
	decoder := plist.NewDecoder(buf)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	tempMap = data.Objects[data.RootObject]
	targets = tempMap["targets"].([]interface{})
	targetsBuildSetting := data.Objects[targets[0].(string)]
	fmt.Println("targetsBuildSetting-->: ", targetsBuildSetting)
	buildConfigurationList := targetsBuildSetting["buildConfigurationList"].(string)
	fmt.Println("buildConfigurationList-->: ", buildConfigurationList)
	buildConfigurations = data.Objects[buildConfigurationList]["buildConfigurations"].([]interface{})
	for _, item := range buildConfigurations {
		itemString := item.(string)
		buildSettingKey_Releasae = itemString
		if buildSettingKey_Releasae == "" {
			fmt.Println("没有找到 BuildSettings!!!")
			os.Exit(0)
		}
		key_Development_Team := "objects:" + buildSettingKey_Releasae + ":buildSettings:DEVELOPMENT_TEAM"
		if execCommand("/usr/libexec/PlistBuddy", "-c", "Set :" + key_Development_Team  + " " + codeSignTeamIdentifier(info) , path) == nil {
			fmt.Println("pbxproj文件修改成功--> DEVELOPMENT TEAM.!!!")
		}
		key_BundleID := "objects:" + buildSettingKey_Releasae + ":buildSettings:PRODUCT_BUNDLE_IDENTIFIER"
		_ = loadApplicationConfigInfo(Info_Build)
		if execCommand("/usr/libexec/PlistBuddy", "-c", "Set :" + key_BundleID  + " "  + applicationConfigInfo.Application_BundleIdentifier , path) == nil {
			fmt.Println("pbxproj文件修改成功--> Bundle ID.!!!")
		}
	}
	return nil
}

func changeTargetExportOptionsFile(info BuildInfo) error {
	return execCommand("cp", inside_CurrentPath + "/ExportOptionsConfig/" + codeSignTeamIdentifier(info) + ".plist", inside_CurrentPath + "/ExportOptions.plist")
}