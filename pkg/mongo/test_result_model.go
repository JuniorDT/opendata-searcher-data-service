package mongo

import (
	"github.com/JuniorDT/opendata-searcher-data-service/pkg/services_test_results/common_service"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// this structure of result of test parse file for insert in Mongodb
type FileParseResultModel struct {
	ID      	bson.ObjectId  	`bson:"_id,omitempty" json:"id"`
	Title   	string      	`bson:"title,omitempty" json:"title"`
	Filename 	string			`bson:"filename" json:"filename"`
	URL     	string      	`bson:"url,omitempty" json:"url"`				// url for source file
	Size    	float64     	`bson:"size" json:"size,string"`				// size (KB)
	ParseTime   int32	    	`bson:"parse_time" json:"parse_time,string"`	// parsing time (sec)
	Expansion	string			`bson:"expansion" json:"expansion,string"`	// file expansion
	ParseMethod string  		`bson:"parse_method" json:"parse_method"`		// method for parse
}

// index for searching by URL in Mongodb
func FileParseResultIndex() mgo.Index {
	return mgo.Index{
		Key: 		[]string{"url"},
		Unique: 	true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// convert file test parse result from FileParseResult to FileParseResultModel
func newFileParseResultModel(r *commontestresults.FileParseResult) *FileParseResultModel {
	return &FileParseResultModel{
		Title: 		 	r.Title,
		Filename:	 	r.Filename,
		URL: 		 	r.URL,
		Size: 		 	r.Size,
		ParseTime:		r.ParseTime,
		Expansion:		r.Expansion,
		ParseMethod: 	r.ParseMethod,
	}
}

// convert from FileParseResultModel to FileParseResult
func(r *FileParseResultModel) toFileParseResult() *commontestresults.FileParseResult {
	return &commontestresults.FileParseResult{
		ID: 		 	r.ID.Hex(),
		Title: 		 	r.Title,
		Filename:	 	r.Filename,
		URL: 		 	r.URL,
		Size: 		 	r.Size,
		ParseTime:		r.ParseTime,
		Expansion:		r.Expansion,
		ParseMethod: 	r.ParseMethod,
	}
}