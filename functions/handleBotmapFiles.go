package functions

import (
	"bufio"
	"go-websocket-connection/domain"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ChargeBotmaps()  map[int]domain.BotmapDomain{
	fileNames := GetAllNameFileBotmaps()
	lines := make(map[int]domain.BotmapDomain)

	var cont = 0

	for _, name := range fileNames{

		f, err := os.Open("./botmaps/" + name)
		defer f.Close()

		check(err)

		var rd = bufio.NewReader(f)

		for  {
			line, err := rd.ReadString('\n')
			cont = cont + 1
			newLine := processLine(line, cont)
			lines[cont] = newLine

			if err == io.EOF {
				break
			}
			check(err)
		}

	}
	//jsonString, err := json.Marshal(lines)
	return lines
}

/**
todo: Deixar diretório dinâmico
description: Essa função retorna todos os nomes dos arquivos botmap
 */
func GetAllNameFileBotmaps() []string {
	var fileNames []string
	files, err := ioutil.ReadDir("botmaps")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files{
		if strings.Index(f.Name(), ".botmap") > 0 {
			fileNames = append(fileNames, f.Name())
		}
	}

	return fileNames
}

func GetMessage(index int) string {
	messages := ChargeBotmaps()
	return messages[index].Content
}
