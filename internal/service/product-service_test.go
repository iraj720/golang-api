package service

import (
	"snapp/internal/entity"
	"testing"
)

func TestHelloName(t *testing.T) {
	var havePoint float32 = 5
	var haveVotes uint64 = 4
	var wantPoint float32 = 6
	var wantVotes uint64 = 5

	var newPoint float32 = 10.0
	beforeProduct := entity.Product{ID: 1, Name: "TV", Point: havePoint, VotesCount: haveVotes}
	
	newProduct := voteToProduct(beforeProduct, newPoint)
	if newProduct.Point != wantPoint {
		t.Fatalf(`voteToProduct(beforeProduct, newPoint) = point is : %f   want point : %f`, newProduct.Point, wantPoint)
	}

	if newProduct.VotesCount != wantVotes {
		t.Fatalf(`voteToProduct(beforeProduct, newPoint) = votes_count is : %d   want votes : %d`, newProduct.VotesCount, wantVotes)
	}
}
