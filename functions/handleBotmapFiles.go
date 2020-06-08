package functions

import (
	"bufio"
	"go-websocket-connection/domain"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ChargeBotmaps()  map[int]domain.BotmapDomain{

	lines := make(map[int]domain.BotmapDomain)

	f, err := os.Open("./botmaps/index.botmap")
	defer f.Close()

	check(err)

	var rd = bufio.NewReader(f)
	var cont = 0

	for  {
		line, err := rd.ReadString('\n')


		cont = cont + 1

		newLine := processLine(line, cont)

		lines[cont] = newLine

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
	}
	//jsonString, err := json.Marshal(lines)

	return lines
}

func GetMessage(index int) string {
	messages := ChargeBotmaps()
	return messages[index].Content
}
