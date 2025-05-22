package config

type Config struct {
	Model     *Model     `yaml:"model"`
	Excel     *Excel     `yaml:"excel"`
	ResultMap *ResultMap `yaml:"resultMap"`
}

type Model struct {
	PackageName string `yaml:"packageName"`
	ModelName   string `yaml:"modelName"`
	Fields      []struct {
		Name    string `yaml:"name"`
		Type    string `yaml:"type"`
		Tag     string `yaml:"tag"`
		Comment string `yaml:"comment"`
	} `yaml:"fields"`
}

type Excel struct {
	PackageName string `yaml:"packageName"`
	SheetName   string `yaml:"sheetName"`
	Name        string `yaml:"name"`
	Model       struct {
		PackageName  string     `yaml:"packageName"`
		Name         string     `yaml:"name"`
		Headers      [][]string `yaml:"headers"`
		MergeHeaders []string   `yaml:"mergeHeaders"`
		Fields       []string   `yaml:"fields"`
	} `yaml:"model"`
}

type ResultData struct {
	ModelPackage string    `yaml:"modelPackage"`
	ModelName        string    `yaml:"modelName"`
	KeyColumns  string    `yaml:"keyColumns"`
	Source      []string  `yaml:"source"`
	Target      []string  `yaml:"target"`
	ResultMap   ResultMap `yaml:"resultMap"`
}

type ResultMap struct {
	PackageName string       `yaml:"packageName"`
	ModelPackage string    `yaml:"modelPackage"`
	ModelName        string    `yaml:"modelName"`
	ResultData  []ResultData `yaml:"resultData"`
}
