package domain

type BotmapDomain struct {
	Index         int `json:"index"`
	Content       string	`json:"content"`
	PositiveIndex int	`json:"positiveFlow"`
	NegativeIndex int	`json:"negativeFlow"`
	TypeMessage   string	`json:"type"`
}
