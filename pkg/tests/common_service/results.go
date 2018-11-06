package commontestresults

type FileParseResult struct {
	ID      		string  	`json:"id"`
	Title   		string      `json:"title"`
	Filename		string		`json:"filename"`
	URL     		string      `json:"url"`
	Size    		float64     `json:"size,string"`
	ParseTime    	int32	    `json:"parse_time,string"`
	ParseMethod 	string  	`json:"parse_method"`
	Expansion		string		`json:"expansion"`
}

type FileParseResultService interface {
	Create(r *FileParseResult) error
	GetById(id string) (r *FileParseResult, err error)
}

