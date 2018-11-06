package commontestresults

type CSVParseResult struct {
	ID      string  	`json:"id"`
	Title   string      `json:"title"`
	URL     string      `json:"url"`
	Size    float64     `json:"size,string"`
	Time    int32	    `json:"time,string"`
	ParseMethod string  `json:"parse_method"`
}

type CSVParseResultService interface {
	Create(r *CSVParseResult) error
	GetById(id string) (r *CSVParseResult, err error)
}

