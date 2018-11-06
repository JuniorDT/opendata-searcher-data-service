package mongo

import (
	"github.com/JuniorDT/opendata-searcher-data-service/pkg/tests/common_service"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// this structure of result of test parse CSV for insert in Mongodb
type CSVParseResultModel struct {
	ID      	bson.ObjectId  	`bson:"_id,omitempty" json:"id"`
	Title   	string      	`bson:"title" json:"title"`
	URL     	string      	`bson:"url" json:"url"`
	Size    	float64     	`bson:"size" json:"size,string"`
	Time    	int32	    	`bson:"time" json:"time,string"`
	ParseMethod string  		`bson:"parse_method" json:"parse_method"`
}

// index for searching by URL in Mongodb
func CSVParseResultIndex() mgo.Index {
	return mgo.Index{
		Key: 		[]string{"url"},
		Unique: 	true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// convert CSV parse result from CSVParseResult to CSVParseResultModel
func newCSVParseResultModel(r *commontestresults.CSVParseResult) *CSVParseResultModel {
	return &CSVParseResultModel{
		Title: 		 r.Title,
		URL: 		 r.URL,
		Size: 		 r.Size,
		Time: 		 r.Time,
		ParseMethod: r.ParseMethod,
	}
}

// convert from CSVParseResultModel to CSVParseResult
func(r *CSVParseResultModel) toDefaultStructure() *commontestresults.CSVParseResult {
	return &commontestresults.CSVParseResult{
		ID: 		 r.ID.Hex(),
		Title: 		 r.Title,
		URL: 		 r.URL,
		Size: 		 r.Size,
		Time: 		 r.Time,
		ParseMethod: r.ParseMethod,
	}
}