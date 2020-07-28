package sql2struct

import (
	"fmt"
	"os"
	"text/template"
	"tour/internal/word"
)

/*
//模板渲染 基本结构由一个go结构体 和其所属的tablename方法组成
/*
生成结果大致如下
type 大驼峰的表名 struct {
	//注释
	字段名 字段类型
	...	...
}
func (model 大驼峰的表名) TableName() string {
	return "表名称"
}

*/


const strcutTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}
func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

//转换后的go结构体信息
type StructColumn struct {
	Name string
	Type string
	Tag string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: strcutTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn  {
	tplColumns := make([]*StructColumn,0,len(tbColumns))
	for _,column := range tbColumns {
		tplColumns = append(tplColumns,&StructColumn{
			Name: column.ColumnName,
			Type: DBTypeToStructType[column.DataType],
			Tag: fmt.Sprintf("`json:"+"%s"+"`",column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

//对模板自定义函数和模块对象进行处理
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {

	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	//最后调用excute进行渲染
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}