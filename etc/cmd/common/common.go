package common // import "go.pedge.io/pb/etc/cmd/common"

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"go.pedge.io/env"
)

// GenerateEnv is the environment for Generators.
type GenerateEnv struct {
	CSVFile    string `env:"CSV_FILE,required"`
	GoTmplFile string `env:"GO_TMPL_FILE"`
	GoFile     string `env:"GO_FILE"`
	GogoFile   string `env:"GOGO_FILE"`
	PbTmplFile string `env:"PB_TMPL_FILE"`
	PbFile     string `env:"PB_FILE"`
}

// TmplData is what is passed into the template.
type TmplData struct {
	// either proto, go, or gogo
	Type string
	// The data from the GenerateHelper
	Data interface{}
}

// Main is the main function for commands.
func Main(generateHelper GenerateHelper) {
	env.Main(
		func(generateEnvObj interface{}) error {
			return NewGenerator(generateHelper).Generate(generateEnvObj.(*GenerateEnv))
		},
		&GenerateEnv{},
	)
}

// GenerateHelper does stuff for Generators.
type GenerateHelper interface {
	TmplData(records [][]string) (interface{}, error)
	ExtraTmplFuncs() map[string]interface{}
}

// Generator generates stuff.
type Generator interface {
	Generate(generateEnv *GenerateEnv) error
}

// NewGenerator creates a new Generator for the given GenerateHelper.
func NewGenerator(generateHelper GenerateHelper) Generator {
	return newGenerator(generateHelper)
}

type generator struct {
	generateHelper GenerateHelper
}

func newGenerator(generateHelper GenerateHelper) *generator {
	return &generator{generateHelper}
}

func (g *generator) Generate(generateEnv *GenerateEnv) (retErr error) {
	csvFile, err := os.Open(generateEnv.CSVFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := csvFile.Close(); err != nil && retErr == nil {
			retErr = err
		}
	}()
	records, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return err
	}
	for i, record := range records {
		records[i] = cleanRecord(record)
	}
	internalTmplData, err := g.generateHelper.TmplData(records)
	if err != nil {
		return err
	}
	if (generateEnv.GoFile != "" || generateEnv.GogoFile != "") && generateEnv.GoTmplFile != "" {
		goTemplate, err := g.newTemplate(generateEnv.GoTmplFile)
		if err != nil {
			return err
		}
		if generateEnv.GoFile != "" {
			goFile, err := os.Create(generateEnv.GoFile)
			if err != nil {
				return err
			}
			defer func() {
				if err := goFile.Close(); err != nil && retErr == nil {
					retErr = err
				}
			}()
			if err := goTemplate.Execute(goFile, newTmplData("go", internalTmplData)); err != nil {
				return err
			}
		}
		if generateEnv.GogoFile != "" {
			gogoFile, err := os.Create(generateEnv.GogoFile)
			if err != nil {
				return err
			}
			defer func() {
				if err := gogoFile.Close(); err != nil && retErr == nil {
					retErr = err
				}
			}()
			if err := goTemplate.Execute(gogoFile, newTmplData("gogo", internalTmplData)); err != nil {
				return err
			}
		}
	}
	if generateEnv.PbFile != "" && generateEnv.PbTmplFile != "" {
		pbTemplate, err := g.newTemplate(generateEnv.PbTmplFile)
		if err != nil {
			return err
		}
		pbFile, err := os.Create(generateEnv.PbFile)
		if err != nil {
			return err
		}
		defer func() {
			if err := pbFile.Close(); err != nil && retErr == nil {
				retErr = err
			}
		}()
		if err := pbTemplate.Execute(pbFile, newTmplData("proto", internalTmplData)); err != nil {
			return err
		}
	}
	return nil
}

func (g *generator) newTemplate(file string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"add": func(i int, j int) int {
			return i + j
		},
	}
	for key, value := range g.generateHelper.ExtraTmplFuncs() {
		funcMap[key] = value
	}
	return template.New(filepath.Base(file)).Funcs(funcMap).ParseFiles(file)
}

func newTmplData(t string, internalTmplData interface{}) *TmplData {
	return &TmplData{
		Type: t,
		Data: internalTmplData,
	}
}

func cleanRecord(record []string) []string {
	for i, elem := range record {
		record[i] = strings.TrimSpace(elem)
	}
	return record
}
