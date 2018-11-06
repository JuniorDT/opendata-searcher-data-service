package mongo

import (
	"github.com/JuniorDT/opendata-searcher-data-service/pkg/tests/common_service"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FileParseResultService struct {
	collection *mgo.Collection
}

// get collection, add index
func NewTestResultService(session *Session, dbName string, collectionName string) *FileParseResultService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(FileParseResultIndex())
	return &FileParseResultService {collection}
}

func(r *FileParseResultService) Create(t *commontestresults.FileParseResult) error {
	testResult := newFileParseResultModel(t)
	return r.collection.Insert(&testResult)
}

func(r *FileParseResultService) GetById(id string) (*commontestresults.FileParseResult, error) {
	resultModel := FileParseResultModel{}
	err := r.collection.FindId(bson.ObjectIdHex(id)).One(&resultModel)
	return resultModel.toFileParseResult(), err
}