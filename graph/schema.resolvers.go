package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"os"
	"regexp"
	"seim/graph/generated"
	"seim/graph/model"
	"seim/ml/client/pb"
)

func (r *mutationResolver) UpdateImage(ctx context.Context, input model.UpdateInput) (*model.Image, error) {

	//log.Printf("%v",input)
	re1 := regexp.MustCompile("data:image\\/jpeg;base64,")
	fileData := re1.ReplaceAllString(input.Content, "")
	//log.Printf("%v",fileData)
	data, _ := base64.StdEncoding.DecodeString(fileData) //[]byte
	file, _ := os.Create("resource/pic/encode_and_decord.jpg")
	defer file.Close()
	file.Write(data)

	res, err := r.MlClient.Service.ReadImage(context.Background(),&pb.UpdateInput{Kind: "1"})
	if err != nil {
		return nil, err
	}
	result := &model.Image{ Kind: res.Content}

	return result, nil

	//return nil, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
