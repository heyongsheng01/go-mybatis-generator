package {{.Package}};
{{.Imports}}
import tk.mybatis.mapper.common.Mapper;
import tk.mybatis.mapper.common.MySqlMapper;

{{.Descriptions}}
public interface {{.Name}} extends Mapper<{{.Model}}>, MySqlMapper<{{.Model}}> {

}