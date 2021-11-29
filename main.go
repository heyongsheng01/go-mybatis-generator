package main

import (
	"flag"
	"fmt"

	"github.com/ccc469/go-mybatis-generator/config"
	"github.com/ccc469/go-mybatis-generator/db"
	util "github.com/ccc469/go-mybatis-generator/utils"
)

var (
	ToolName = "go-mybatis-generator"
	Version  = "1.0"
	Usage    = "go-mybatis-generator [OPTIONS] ... COMMAND"
	Options  = `OPTIONS:
	--H value              Host (default: 127.0.0.1)
	--p value              Port (default: 3306)
	--u value              UserName (default: root)
	--P value              Password (default: 123456)
	--d value              Database (default: nil)
	--t value              单个表名 (default: nil)
	--all value            生成所有表 (default: false)
	--model value          Model包名 (default: com.xxx.entity)
	--mapper value         Mapper接口包名 (default: com.xxx.mapper)
	--xml value            Xml文件包名 (default: com.xxx.xml)
	--h value              显示帮助`
	PrintHelp = false
)

func main() {

	SetOptions()
	ShowInfo()

	// 初始化sql
	db.InitDB()

	// 执行生成方法
	util.Run()
}

// 显示内容
func ShowInfo() {
	if PrintHelp {
		fmt.Println(Usage + "\n")
		fmt.Println(Options)
	} else {
		fmt.Println("\nYour Setting: ------------------------------------")
		fmt.Printf("%30s %v\n", "               host:", config.Host)
		fmt.Printf("%30s %v\n", "               port:", config.Port)
		fmt.Printf("%30s %v\n", "           username:", config.UserName)
		fmt.Printf("%30s %v\n", "           password:", config.Password)
		fmt.Printf("%30s %v\n", "           database:", config.Database)
		fmt.Printf("%30s %v\n", "              table:", config.Table)
		fmt.Printf("%30s %v\n", "      model package:", config.ModelPackage)
		fmt.Printf("%30s %v\n", "     mapper package:", config.MapperPackage)
		fmt.Printf("%30s %v\n", "        xml package:", config.XmlPackage)
		fmt.Printf("%30s %v\n", "generator all table:", config.IsAllTables)
		fmt.Println()
	}

}

// 设置参数
func SetOptions() {
	flag.BoolVar(&PrintHelp, "h", false, "使用-h查看详细")

	if !PrintHelp {
		flag.StringVar(&config.Table, "t", "nil", "Mysql Table")
		flag.BoolVar(&config.IsAllTables, "all", false, "是否生成所有表")
		flag.StringVar(&config.Database, "d", "nil", "Mysql Database")
		flag.StringVar(&config.Host, "H", "127.0.0.1", "Mysql Host")
		flag.IntVar(&config.Port, "p", 3306, "Mysql Port")
		flag.StringVar(&config.UserName, "u", "root", "Mysql Username")
		flag.StringVar(&config.Password, "P", "123456", "Mysql Password")
		flag.StringVar(&config.ModelPackage, "model", "com.example.entity", "实体类包名")
		flag.StringVar(&config.MapperPackage, "mapper", "com.example.mapper", "Mapper接口包名")
		flag.StringVar(&config.XmlPackage, "xml", "com.example.xml", "xml包名")
	}

	flag.Parse()
}
