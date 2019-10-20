/* viper: github.com/spf13/viper
  1. 支持Yaml、Json、TOML、HCL 等格式的配置
  2. 可以从文件、io、环境变量、command line中提取配置
  3. 支持自动转换的类型解析
  4. 可以远程从etcd中读取配置
  5. 监听配置变化，默认值

  每一种配置文件的解析都有包，viper.go里都有，可以精简
  yaml "gopkg.in/yaml.v2"，官网
  "golang.org/x/text/transform" 去 https://github.com/golang/text下载
*/

package mconfig

import (
	"fmt"

	"github.com/Unknwon/goconfig"

	"github.com/spf13/viper"
)

type cfgHelper struct {
	v *viper.Viper
}

var (
	mcfgHelper cfgHelper
)

type userInfo struct {
	Username string
	Hobbies  []string
}

func init() {
	//-------------------goconfig----------------------
	iniCfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	str, _ := iniCfg.GetValue(goconfig.DEFAULT_SECTION, "google")
	fmt.Println("str:", str)

	// str = iniCfg.MustValue("What's this?", "name", "maoyinan")
	// fmt.Println("str2:", str)

	// // [What's this?] no comments
	// comments := iniCfg.GetSectionComments("Demo")
	// fmt.Println("str3:", comments)

	// keycomments := iniCfg.GetKeyComments("What's this?", "name")
	// fmt.Println("key str:", keycomments)

	keyList := iniCfg.GetKeyList("Demo")
	fmt.Println(keyList)

	sectionMap, err := iniCfg.GetSection("Demo")
	if err == nil {
		fmt.Println(sectionMap)
	}

	b := iniCfg.DeleteKey(goconfig.DEFAULT_SECTION, "google")
	fmt.Println("b:", b)
	b = iniCfg.SetValue(goconfig.DEFAULT_SECTION, "google", "www.baidu.com")
	fmt.Println("b:", b)

	// 保存修改
	err = goconfig.SaveConfigFile(iniCfg, "config.ini")
	if err != nil {
		fmt.Println(err)
	}

	//-------------------viper----------------------
	// mcfgHelper.v = viper.New()
	// var cfg *viper.Viper = mcfgHelper.v
	// cfg.SetConfigName("config")
	// cfg.SetConfigName("simple")

	// cfg.AddConfigPath("%GOPATH/src/demo/mcfg/") // C:\\Users\\maoyi\\go\\%GOPATH\\src\\demo\\mcfg
	// cfg.AddConfigPath("../mcfg")       // C:\\Users\\maoyi\\go 或者 demo\\slice\（debug）
	// cfg.AddConfigPath("src/demo/mcfg") // ctrl + shift + B

	// cfg.SetDefault("Version", "1.1.0") //配置文件中有Version字段就不行

	//ReadConfig(io.Reader)
	// if err := cfg.ReadInConfig(); err != nil {
	// 	fmt.Println("can not find config.yaml!")
	// } else {
	// 	fmt.Println("read config.yaml.")
	// }

	// userInfo字段必须和yaml一样，首字母必须大写
	// var usrinfo userInfo
	// if err := cfg.Unmarshal(&usrinfo); err != nil {
	// 	fmt.Println("unmarshal error:", err)
	// }

	// ver := cfg.GetString("Version")
	// author := cfg.GetString("Author")
	// fmt.Println(ver, author)
	// cfg.Set("Version", "1.1.1") //没有Version字段，则不起作用
	// if err := cfg.WriteConfig(); err != nil {
	// 	fmt.Println("write config.xml error: ", err.Error())
	// }

	// WriteConfig | SafeWriteConfig  | WriteConfigAs  | SafeWriteConfigAs

	// cfg.WatchConfig()
	// cfg.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name) // 触发2遍？
	// })
}
