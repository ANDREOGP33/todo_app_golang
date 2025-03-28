package domain

type Tarefa struct {
	Nome        string `form:"nome" json:"nome"`
	Descricao   string `form:"descricao" json:"descricao"`
	Concluida   bool   `form:"concluida" json:"concluida"`
	DataCriacao string `json:"data_criacao"`
}
