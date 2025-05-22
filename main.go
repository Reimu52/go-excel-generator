package main

import (
	_ "embed"
	"flag"
	"fmt"

	"os"

	"text/template"

	"gopkg.in/yaml.v3"
	"remu52.com/excelGenerator/config"
)

//go:embed schemas/config.yml
var configFile []byte

//go:embed templates/model.tmpl
var modelTmplFile []byte

//go:embed templates/excel.tmpl
var excelTmplFile []byte

func main() {
	// exportExcel()
	configPath := flag.String("c", "", "配置文件路径")
	modelTmplPath := flag.String("m", "", "model模板文件路径")
	excelTmplPath := flag.String("e", "", "excel模板文件路径")
	flag.Parse()
	var data []byte
	var err error
	if *configPath == "" {
		data = configFile
	} else {
		data, err = os.ReadFile(*configPath)
		if err != nil {
			panic(fmt.Errorf("读取文件失败: %v", err))
		}
	}
	var yamlConfig config.Config
	err = yaml.Unmarshal(data, &yamlConfig)
	if err != nil {
		panic(fmt.Errorf("解析yaml文件失败: %v", err))
	}

	if yamlConfig.Model != nil {
		var tmpl *template.Template
		var err error
		if *modelTmplPath == "" {
			tmpl, err = template.New("model").Parse(string(modelTmplFile))
		} else {
			tmpl, err = template.ParseFiles(*modelTmplPath)
			if err != nil {
				panic(fmt.Errorf("创建文件失败: %v", err))
			}
		}
		file, err := os.OpenFile(yamlConfig.Model.ModelName+".go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic(fmt.Errorf("创建文件失败: %v", err))
		}
		defer file.Close()
		tmpl.Execute(file, yamlConfig.Model)

	}

	if yamlConfig.Excel != nil {
		var tmpl *template.Template
		var err error
		if *excelTmplPath == "" {
			tmpl, err = template.New("excel").Parse(string(excelTmplFile))
		} else {
			tmpl, err = template.ParseFiles(*excelTmplPath)
			if err != nil {
				panic(fmt.Errorf("创建文件失败: %v", err))
			}
		}
		file, err := os.OpenFile(yamlConfig.Excel.Name+".go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic(fmt.Errorf("创建文件失败: %v", err))
		}
		defer file.Close()
		tmpl.Execute(file, yamlConfig.Excel)

	}
}
