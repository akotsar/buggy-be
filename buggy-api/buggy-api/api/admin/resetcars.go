package admin

import (
	"buggy/api/admin/seed"
	"buggy/api/requestcontext"
	"buggy/internal/data/datacommon"
	"buggy/internal/data/makedata"
	"buggy/internal/data/modeldata"
	"buggy/internal/data/votedata"
	"buggy/internal/httpresponses"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func resetCarsHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	votedata.DeleteAllVotes(context.Session)
	modeldata.DeleteAllModels(context.Session)
	makedata.DeleteAllMakes(context.Session)

	seedMakes := seed.GetMakes()

	for _, m := range seedMakes {
		makeID := datacommon.GenerateNewID()
		makeRecordID := makedata.GenerateMakeRecordID(makeID)
		var votes int
		for _, model := range m.Models {
			votes += model.Votes
		}
		makeRecord := makedata.MakeRecord{
			DynamoRecordKey: datacommon.DynamoRecordKey{
				RecordID:  makeRecordID,
				TypeAndID: makeRecordID,
			},
			ShardID:     datacommon.GenerateNewShardID(),
			Name:        m.Name,
			Image:       m.Image,
			Description: m.Description,
			Votes:       votes,
		}

		err := makedata.PutMake(context.Session, &makeRecord)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		modelRecords := make([]*modeldata.ModelRecord, 0, len(m.Models))
		var voteRecords []*votedata.VoteRecord
		for _, model := range m.Models {
			modelID := fmt.Sprintf("%s|%s", makeID, datacommon.GenerateNewID())
			modelRecordID := modeldata.GenerateModelRecordID(modelID)
			modelRecord := modeldata.ModelRecord{
				DynamoRecordKey: datacommon.DynamoRecordKey{
					RecordID:  modelRecordID,
					TypeAndID: modelRecordID,
				},
				ShardID:     datacommon.GenerateNewShardID(),
				Name:        model.Name,
				Image:       model.Image,
				Description: model.Description,
				EngineVol:   model.EngineVol,
				MaxSpeed:    model.MaxSpeed,
				Votes:       model.Votes,
			}

			modelRecords = append(modelRecords, &modelRecord)

			for _, vote := range model.Comments {
				voteRecordID := votedata.GenerateVoteRecordID(modelID, datacommon.GenerateNewID())
				voteRecord := votedata.VoteRecord{
					DynamoRecordKey: datacommon.DynamoRecordKey{
						RecordID:  modelRecordID,
						TypeAndID: voteRecordID,
					},
					DateVoted: vote.DatePosted,
					Comment:   vote.Comment,
				}

				voteRecords = append(voteRecords, &voteRecord)
			}
		}

		err = modeldata.PutModels(context.Session, modelRecords)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		err = votedata.PutVotes(context.Session, voteRecords)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	}

	return httpresponses.CreateAPIResponse(200, nil), nil
}
