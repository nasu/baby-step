package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type Singer struct {
	SingerId   int64
	FirstName  string
	LastName   string
	BirthDate  time.Time
	SingerInfo []byte
	Albums     []*Album
}

type Album struct {
	SingerId    int64
	AlbumId     int64
	Name        string `spanner:"AlbumTitle"`
	SalesAmount int64
}

type SingerAndAlbum struct {
	SingerId    int64
	FirstName   string
	LastName    string
	AlbumId     int64
	Name        string `spanner:"AlbumTitle"`
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
		fmt.Println(model.FirstName, model.AlbumId, model.Name)
	}

	fmt.Println("")
	fmt.Println("===== ARRAY<STRUCT>")
	for _, model := range SelectWithArrayStruct(ctx, ro) {
		fmt.Println(model.FirstName)
		for _, album := range model.Albums {
			fmt.Println("    ", album.AlbumId, album.Name)
		}
	}
}

func SelectWithArrayStruct(ctx context.Context, ro *spanner.ReadOnlyTransaction) []*Singer {
	stmt := spanner.Statement{
		SQL: `SELECT s.SingerId, s.FirstName, ARRAY(SELECT As STRUCT a.AlbumId, a.AlbumTitle FROM Albums a WHERE a.SingerId = s.SingerId) FROM Singers s`,
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
		/* error
		   var s []interface{}{}
		   if err := row.Columns(&model.SingerId, &model.FirstName, &s); err != nil {
		*/
		if err := row.Columns(&model.SingerId, &model.FirstName, &model.Albums); err != nil {
			panic(err)
		}
		models = append(models, model)
	}
	return models
}

func SelectFlatten(ctx context.Context, ro *spanner.ReadOnlyTransaction) []*SingerAndAlbum {
	stmt := spanner.Statement{
		SQL: `SELECT s.SingerId, s.FirstName, a.AlbumTitle FROM Singers s JOIN Albums a USING(SingerId)`,
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
		if err := row.Columns(&model.SingerId, &model.FirstName, &model.Name); err != nil {
			panic(err)
		}
		models = append(models, model)
	}
	return models
}

func SelectWithArrayStructOnly(ctx context.Context, ro *spanner.ReadOnlyTransaction) []*Singer {
	stmt := spanner.Statement{
		SQL: `SELECT ARRAY(SELECT As STRUCT s.* FROM Singers s)`,
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

		if err := row.Columns(&models); err != nil {
			panic(err)
		}
	}
	return models
}

func SelectSimple(ctx context.Context, ro *spanner.ReadOnlyTransaction) []*Singer {
	stmt := spanner.Statement{
		SQL: `SELECT SingerId, FirstName, LastName, BirthDate, SingerInfo FROM Singers`,
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
		if err := row.Columns(&model.SingerId, &model.FirstName, &model.LastName, &model.BirthDate, &model.SingerInfo); err != nil {
			panic(err)
		}
		models = append(models, model)
	}
	return models
}
