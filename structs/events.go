package structs

type Event struct {
	Id     string      `json:"_id"`
	Source interface{} `json:"_source"`
}

type IndexName struct {
	Index string `json:"index"`
}
