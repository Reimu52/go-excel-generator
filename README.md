# go-excel-generator
execl导出代码自动生成工具，通过yaml文件配置xlsx表头，需要导出的字段信息，跨行跨列。
代码模板在templates下。
go run main.go {-c config.yml} {-m model.tmpl} {-e excel.tmpl} 
go run main.go