package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}
