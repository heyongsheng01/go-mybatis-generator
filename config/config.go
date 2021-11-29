package config

var (
	Table         = ""
	IsAllTables   = false
	Database      = ""
	Host          = "127.0.0.1"
	Port          = 3306
	UserName      = "root"
	Password      = "123456"
	ModelPackage  = ""
	MapperPackage = ""
	XmlPackage    = ""
)

var (
	FieldTypes = make(map[string]string)
	JdbcTypes  = make(map[string]string)
)

// InitFieldTypes 初始化数据库字段类型对照表
func InitFieldTypes() {
	FieldTypes["bigint"] = "java.lang.Long"
	FieldTypes["bit"] = "java.lang.Boolean"
	FieldTypes["blob"] = "java.lang.byte[]"
	FieldTypes["char"] = "java.lang.String"
	FieldTypes["date"] = "java.util.Date"
	FieldTypes["datetime"] = "java.util.Date"
	FieldTypes["decimal"] = "java.math.BigDecimal"
	FieldTypes["double"] = "java.lang.Double"
	FieldTypes["double"] = "java.lang.Double"
	FieldTypes["enum"] = "java.lang.String"
	FieldTypes["float"] = "java.lang.Float"
	FieldTypes["int"] = "java.lang.Integer"
	FieldTypes["integer"] = "java.lang.Long"
	FieldTypes["longblob"] = "java.lang.byte[]"
	FieldTypes["longtext"] = "java.lang.String"
	FieldTypes["mediumblob"] = "java.lang.byte[]"
	FieldTypes["mediumint"] = "java.lang.Integer"
	FieldTypes["mediumtext"] = "java.lang.String"
	FieldTypes["set"] = "java.lang.String"
	FieldTypes["smallint"] = "java.lang.Integer"
	FieldTypes["text"] = "java.lang.String"
	FieldTypes["time"] = "java.sql.Time"
	FieldTypes["timestamp"] = "java.sql.Timestamp"
	FieldTypes["tinyblob"] = "java.lang.byte[]"
	FieldTypes["tinyint"] = "java.lang.Integer"
	FieldTypes["tinytext"] = "java.lang.String"
	FieldTypes["varchar"] = "java.lang.String"

}

// InitSqlJdbcTypes 初始化jdbcType与sql类型对照表
func InitSqlJdbcTypes() {
	JdbcTypes["char"] = "CHAR"
	JdbcTypes["varchar"] = "VARCHAR"
	JdbcTypes["tinyint"] = "TINYINT"
	JdbcTypes["smallint"] = "SMALLINT"
	JdbcTypes["int"] = "INTEGER"
	JdbcTypes["float"] = "FLOAT"
	JdbcTypes["bigint"] = "BIGINT"
	JdbcTypes["double"] = "DOUBLE"
	JdbcTypes["bit"] = "BOOLEAN"
	JdbcTypes["date"] = "TIMESTAMP"
	JdbcTypes["datetime"] = "TIMESTAMP"
	JdbcTypes["time"] = "TIMESTAMP"
	JdbcTypes["text"] = "VARCHAR"
	JdbcTypes["longtext"] = "LONGVARCHAR"
}
