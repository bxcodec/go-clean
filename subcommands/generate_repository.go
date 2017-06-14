package subcommands

import (
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/Masterminds/sprig"
)

type RepoGenerator struct {
	PackageShort string
	ModelName    string
	FetchParams  string
	Imports      map[string]string
	ImportsShort map[string]string
}

func (s *Subs) generateRepository() {
	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/repositoryInterface.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "repository/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}

	k := 1

	mapImport := make(map[string]string)
	mapImport["models"] = "github.com/bxcodec/gclean/models"
	mapImport["time"] = "time"
	mapImportShort := make(map[string]string)
	mapImportShort["models"] = "models"
	mapImportShort["time"] = "time"

	dataSend := &RepoGenerator{
		PackageShort: "models",
		ModelName:    "Article",
		FetchParams:  "cursor string , num int64",
		Imports:      mapImport,
		ImportsShort: mapImportShort,
	}
	f, err := os.Create(pathP + "sample" + strconv.Itoa(k) + ".go")
	if err != nil {
		fmt.Println("Erorr")
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "repositoryInterface.tpl", dataSend)

	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(0)
	}
}
