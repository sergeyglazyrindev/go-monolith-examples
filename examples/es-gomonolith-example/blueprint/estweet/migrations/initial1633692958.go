package migrations

import (
	"context"
	"fmt"
	"github.com/sergeyglazyrindev/go-monolith/core"
)

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
	ID      string
}

type initial1633692958 struct {
}

func (m initial1633692958) GetName() string {
	return "estweet.1633692958"
}

func (m initial1633692958) GetID() int64 {
	return 1633692958
}

func (m initial1633692958) Up(database *core.ProjectDatabase) error {
	// Create a client
	client := core.NewProjectESClient()
	client.DeleteIndex("tweets").Do(context.Background())
	// Create an index
	_, err := client.CreateIndex("tweets").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	j := 0
	for {
		if j == 100 {
			break
		}
		tweet := Tweet{User: fmt.Sprintf("olivere-%d", j), Message: fmt.Sprintf("Take Five-%d", j)}
		_, err = client.Index().
			Index("tweets").
			Type("doc").
			// Id("1").
			BodyJson(tweet).
			Refresh("wait_for").
			Do(context.Background())

		j += 1
	}
	return nil
}

func (m initial1633692958) Down(database *core.ProjectDatabase) error {
	client := core.NewProjectESClient()
	client.DeleteIndex("tweets").Do(context.Background())
	return nil
}

func (m initial1633692958) Deps() []string {
	return make([]string, 0)
}
