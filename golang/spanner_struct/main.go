package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type Singer struct {
	SingerId  string
	FirstName string
	LastName  string
	LikeNum   int64
	Albums    []*Album
}

type Album struct {
	SingerId    string
	AlbumId     int64
	Name        string
	SalesAmount int64
}

type SingerAndAlbum struct {
	SingerId    string
	FirstName   string
	LastName    string
	LikeNum     int64
	AlbumId     int64
	Name        string
	SalesAmount int64
}

func prepare(args []string) *spanner.Client {
	//fmt.Println(args)
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	client := prepare(os.Args)
	ctx := context.Background()
	ro := client.ReadOnlyTransaction()
	defer ro.Close()

	fmt.Println("===== FLATTEN")
	for _, model := range SelectFlatten(ctx, ro) {
		fmt.Println(model.FirstName, model.LastName, model.AlbumId, model.Name, model.SalesAmount)
	}

	fmt.Println("")
	fmt.Println("===== ARRAY<STRUCT>")
	for _, model := range SelectWithArrayStruct(ctx, ro) {
		fmt.Println(model.FirstName, model.LastName)
		for _, album := range model.Albums {
			fmt.Println("    ", album.AlbumId, album.Name, album.SalesAmount)
		}
	}
}

func SelectWithArrayStruct(ctx context.Context, ro *spanner.ReadOnlyTransaction) []*Singer {
	stmt := spanner.Statement{
		SQL: `SELECT s.*, ARRAY(SELECT As STRUCT a.* FROM Albums a WHERE a.SingerId = s.SingerId) FROM Singers s`,
	}
	iter := ro.Query(ctx, stmt)
	var models []*Singer
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}

		model := &Singer{}
		if err := row.Columns(&model.SingerId, &model.FirstName, &model.LastName, &model.LikeNum, &model.Albums); err != nil {
			panic(err)
		}
		models = append(models, model)
	}
	return models
}

func SelectFlatten(ctx context.Context, ro *spanner.ReadOnlyTransaction) []*SingerAndAlbum {
	stmt := spanner.Statement{
		SQL: `SELECT * FROM Singers s JOIN Albums a USING(SingerId)`,
	}
	iter := ro.Query(ctx, stmt)
	var models []*SingerAndAlbum
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}

		model := &SingerAndAlbum{}
		if err := row.Columns(&model.SingerId, &model.FirstName, &model.LastName, &model.LikeNum, &model.Name, &model.SalesAmount, &model.AlbumId); err != nil {
			panic(err)
		}
		models = append(models, model)
	}
	return models
}
