package util

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/ccc469/go-mybatis-generator/config"
	"github.com/ccc469/go-mybatis-generator/db"
)

var (
	OutFileDir = "./out/"
	Author     string
)

// JavaModel Java model
type JavaModel struct {
	Package      string
	Imports      string
	TableName    string
	Name         string
	Fields       string
	Descriptions string
	Annotations  string
}

// Mapper Java mapper
type Mapper struct {
	Package      string
	Imports      string
	Descriptions string
	Annotations  string
	Name         string
	Model        string
}

type XmlModel struct {
	Mapper  string
	Model   string
	Results string
}

// GetModelTpl
func GeneratorModel(items []map[string]string, table map[string]string) {
	tpl, err := template.ParseFiles("./template/tk/model.tpl")
	if err != nil {
		panic(err)
	}
	javaName := ToJavaName(table["table_name"])

	var (
		_fields       strings.Builder
		_imports      strings.Builder
		_annotations  strings.Builder
		_descriptions strings.Builder
	)
	_imports.WriteString("\n")

	for _, it := range items {

		// 导包 (排除 java.lang.*)
		if !strings.Contains(_imports.String(), config.FieldTypes[it["data_type"]]) && !strings.Contains(config.FieldTypes[it["data_type"]], "java.lang.") {
			_imports.WriteString("import " + config.FieldTypes[it["data_type"]] + ";\n")
		}
		// 写入lombok 注解
		if !strings.Contains(_imports.String(), "lombok.Data") {
			_imports.WriteString("import lombok.Data;\n")
		}
		if !strings.Contains(_imports.String(), "lombok.Builder") {
			_imports.WriteString("import lombok.Builder;\n")
		}
		if !strings.Contains(_imports.String(), "lombok.AllArgsConstructor") {
			_imports.WriteString("import lombok.AllArgsConstructor;\n")
		}
		if !strings.Contains(_imports.String(), "lombok.NoArgsConstructor") {
			_imports.WriteString("import lombok.NoArgsConstructor;\n")
		}

		// 注释
		if !strings.Contains(_descriptions.String(), "@author") {
			_descriptions.WriteString(WriteDescriptions(table["table_name"], table["table_comment"]))
		}

		// 注解
		if !strings.Contains(_annotations.String(), "@Data") {
			_annotations.WriteString(WriteAnnotations())
		}

		// 字段及注释
		_fields.WriteString("\n")
		// 字段注释
		_fields.WriteString(ToJavaBeanFieldCommennt(it["column_comment"] + "\n"))

		// 判断是否是主键生成主键的注解
		if strings.Contains(it["column_key"], "PRI") && !strings.Contains(_imports.String(), "import javax.persistence.*;") {
			_fields.WriteString("	@Id\n")
			// 如果是主键则倒入@Id需要的包
			_imports.WriteString("import javax.persistence.*;\n")
		}
		// 判断是否是自增ID
		if strings.Contains(it["extra"], "auto_increment") {
			_fields.WriteString("	@GeneratedValue(strategy = GenerationType.IDENTITY)\n")
		}
		// 字段
		_fields.WriteString(ToJavaBeanField(it["column_name"], it["data_type"]))
		_fields.WriteString("\n")

	}

	javaModel := &JavaModel{
		Package:      config.ModelPackage,
		Name:         javaName,
		TableName:    table["table_name"],
		Imports:      _imports.String(),
		Fields:       _fields.String(),
		Descriptions: _descriptions.String(),
		Annotations:  _annotations.String(),
	}

	filePath := OutFileDir + strings.Replace(config.ModelPackage, ".", "/", -1)
	CheckPath(filePath)
	file, err := os.OpenFile(filePath+"/"+javaName+".java", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open failed err:", err)
		return
	}

	tpl.Execute(file, javaModel)
}

// GeneratorMapper 生成mapper
func GeneratorMapper(table map[string]string) {
	tpl, err := template.ParseFiles("./template/tk/mapper.tpl")
	if err != nil {
		panic(err)
	}

	var (
		_imports      strings.Builder
		_descriptions strings.Builder
		_annotations  strings.Builder
	)

	_imports.WriteString("\n")

	javaName := ToJavaName(table["table_name"])

	// 导入实体包
	if !strings.Contains(_imports.String(), javaName) {
		_imports.WriteString(fmt.Sprintf("import %s.%s;", config.ModelPackage, javaName))
	}

	// 类注释
	if !strings.Contains(_descriptions.String(), "@author") {
		_descriptions.WriteString(WriteDescriptions(table["table_name"], table["table_comment"]))
	}

	// // 注解
	// if !strings.Contains(_annotations.String(), "") {
	// 	_annotations.WriteString("")
	// }

	// Mapper Java mapper
	type Mapper struct {
		Package      string
		Imports      string
		Descriptions string
		Annotations  string
		Name         string
		Model        string
	}

	mapper := &Mapper{
		Package:      config.MapperPackage,
		Imports:      _imports.String(),
		Descriptions: _descriptions.String(),
		Annotations:  _annotations.String(),
		Name:         javaName + "Mapper",
		Model:        javaName,
	}

	filePath := OutFileDir + strings.Replace(config.MapperPackage, ".", "/", -1)
	CheckPath(filePath)
	file, err := os.OpenFile(filePath+"/"+javaName+"Mapper.java", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open failed err:", err)
		return
	}

	tpl.Execute(file, mapper)
}

// GeneratorXml 生成xml
func GeneratorXml(items []map[string]string, table map[string]string) {

	tpl, err := template.ParseFiles("./template/tk/xml.tpl")
	if err != nil {
		panic(err)
	}

	javaName := ToJavaName(table["table_name"])
	var _results strings.Builder

	for _, it := range items {

		if !strings.Contains(_results.String(), it["column_name"]) {
			if strings.Contains(it["column_key"], "PRI") {
				// 主键
				_results.WriteString("<id column=\"" + it["column_name"] + "\" jdbcType=\"" + config.JdbcTypes[it["data_type"]] + "\" property=\"" + ToHumpField(it["column_name"]) + "\" />\n")
			} else {
				// 普通字段
				_results.WriteString("		<result column=\"" + it["column_name"] + "\" jdbcType=\"" + config.JdbcTypes[it["data_type"]] + "\" property=\"" + ToHumpField(it["column_name"]) + "\" />\n")
			}
		}
	}

	xmlModel := &XmlModel{
		Mapper:  fmt.Sprintf("%s.%sMapper", config.MapperPackage, javaName),
		Model:   fmt.Sprintf("%s.%s", config.ModelPackage, javaName),
		Results: _results.String(),
	}

	filePath := OutFileDir + strings.Replace(config.XmlPackage, ".", "/", -1)
	CheckPath(filePath)
	file, err := os.OpenFile(filePath+"/"+javaName+"Mapper.xml", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open failed err:", err)
		return
	}

	tpl.Execute(file, xmlModel)

}

// ToJavaName 转换Java名称
func ToJavaName(s string) string {
	arr := strings.Split(s, "_")
	var result string = ""
	for _, str := range arr {
		slen := len(str)
		result = result + strings.ToUpper(str[0:1]) + str[1:slen]
	}
	return result
}

// ToJavaBeanField 转换属性
func ToJavaBeanField(field string, fieldType string) string {
	_fieldType := config.FieldTypes[fieldType]
	_fieldType = GetTypeName(_fieldType)
	_field := ToHumpField(field)
	return "	private " + _fieldType + " " + _field + ";"
}

// ToJavaBeanFieldCommennt 字段备注
func ToJavaBeanFieldCommennt(commennt string) string {
	return "	/**\n" +
		"	 * " + commennt +
		"	 */" + "\n"
}

// GetTypeName 获取类型
func GetTypeName(str string) string {
	arr := strings.Split(str, ".")
	lens := len(arr)
	result := arr[lens-1]
	return result
}

// ToHumpField 转驼峰
func ToHumpField(field string) string {
	arr := strings.Split(field, "_")
	var result string = ""
	for i, str := range arr {
		if i != 0 {
			slen := len(str)
			result = result + strings.ToUpper(str[0:1]) + str[1:slen]
		} else {
			result = result + str
		}
	}
	return result
}

// WriteDescriptions 写入注释
func WriteDescriptions(table string, tableComment string) string {
	return "/**\n" + " * @author " + Author + "\n" +
		" * @time " + time.Now().Format("2006-01-02 15:04:05") + "\n" +
		" * @description " + tableComment + "\n" + " */"
}

// WriteAnnotations 写入注解
func WriteAnnotations() string {
	return "@Data\n" + "@AllArgsConstructor\n" + "@NoArgsConstructor\n" + "@Builder"
}

// 执行生成方法
func Run() {
	var (
		tables  []map[string]string
		columns []map[string]string
	)
	// 查询所有表
	tables = db.GetTables()
	defer db.Close()
	for i := 0; i < len(tables); i++ {
		columns = db.GetTableColumns(tables[i]["table_name"])
		// 生成model
		GeneratorModel(columns, tables[i])
		// 生成mapper
		GeneratorMapper(tables[i])
		// 生成xml
		GeneratorXml(columns, tables[i])
	}

	log.Println("Finish...")
}
