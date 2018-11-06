package mongo

import (
	"github.com/JuniorDT/opendata-searcher-data-service/pkg/tests/common_service"
	"gopkg.in/mgo.v2"
)

type CSVParseResultService struct {
	collection *mgo.Collection
}

func NewTestResultService(session *Session, dbName string, collectionName string) *CSVParseResultService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(CSVParseResultIndex())
	return &CSVParseResultService {collection}
}

func(r *CSVParseResultService) Create(t *commontestresults.CSVParseResult) error {
	testResult := newCSVParseResultModel(t)
	return r.collection.Insert(&testResult)
}

func(r *CSVParseResultService) GetById(id string) (*commontestresults.CSVParseResult, error) {
	resultModel := CSVParseResultModel{}
	err := r.collection.FindId(id).One(&resultModel)
	return resultModel.toDefaultStructure(), err
}