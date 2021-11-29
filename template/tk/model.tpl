package {{.Package}};
{{.Imports}}
import java.io.Serializable;

{{.Descriptions}}
{{.Annotations}}
@Table(name = "{{.TableName}}")
public class {{.Name}} implements Serializable {

    private static final long serialVersionUID = 1L;
    {{.Fields}}
}