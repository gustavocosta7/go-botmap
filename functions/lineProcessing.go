package functions

import (
	"go-websocket-connection/domain"
)

func processLine(line string, cont int)  domain.BotmapDomain{

	var newLine domain.BotmapDomain

	newLine.Index = cont
	newLine.Content = line
	newLine.PositiveIndex = cont + 1
	newLine.NegativeIndex = cont - 1
	newLine.TypeMessage = "normal"

	return newLine

}
