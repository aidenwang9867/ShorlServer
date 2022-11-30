package utils

import (
	"context"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/bigtable"
)

func InitDatabase() (context.Context, *bigtable.Client, error) {
	project := flag.String("project", "rice-comp-539-spring-2022", "The Google Cloud Platform project ID. Required.")
	instance := flag.String("instance", "rice-comp-539-shared-table", "The Google Cloud Bigtable instance ID. Required.")
	flag.Parse()

	for _, f := range []string{"project", "instance"} {
		if flag.Lookup(f).Value.String() == "" {
			log.Fatalf("The %s flag is required.", f)
		}
	}
	ctx := context.Background()

	client, err := bigtable.NewClient(ctx, *project, *instance)
	if err != nil {
		log.Fatalf("Could not create data operations client: %v", err)
		return nil, nil, err
	}
	return ctx, client, nil
}

func WriteToLinks(ctx context.Context, client *bigtable.Client, links []Link) error {
	tbl := client.Open(tableLinks)
	mutations := make([]*bigtable.Mutation, len(links))
	rowKeys := make([]string, len(links))

	for i, link := range links {
		mutations[i] = bigtable.NewMutation()
		mutations[i].Set(linkColumnFamilyName, shortLinkColumnName, bigtable.Now(), []byte(link.Short))
		mutations[i].Set(linkColumnFamilyName, longLinkColumnName, bigtable.Now(), []byte(link.Long))
		rowKeys[i] = link.Short
	}
	rowErrs, err := tbl.ApplyBulk(ctx, rowKeys, mutations)
	if rowErrs != nil {
		for _, rowErr := range rowErrs {
			log.Printf("Error writing row: %v", rowErr)
		}
		log.Fatalf("Could not write some rows")
	}
	if err != nil {
		log.Fatalf("Could not apply bulk row mutation: %v", err)
		return err
	}
	return nil
}

func WriteToUsers(ctx context.Context, client *bigtable.Client, users []User) error {
	tbl := client.Open(tableUsers)
	mutations := make([]*bigtable.Mutation, len(users))
	rowKeys := make([]string, len(users))

	for i, user := range users {
		mutations[i] = bigtable.NewMutation()
		mutations[i].Set(userColumnFamilyName, userIDColumnName, bigtable.Now(), []byte(user.UserID))
		mutations[i].Set(userColumnFamilyName, passwordHashColumnName, bigtable.Now(), []byte(user.PasswordHash))
		mutations[i].Set(userColumnFamilyName, passwordSaltColumnName, bigtable.Now(), []byte(user.PasswordSalt))
		rowKeys[i] = user.UserID
	}
	rowErrs, err := tbl.ApplyBulk(ctx, rowKeys, mutations)
	if rowErrs != nil {
		for _, rowErr := range rowErrs {
			log.Printf("Error writing row: %v", rowErr)
		}
		log.Fatalf("Could not write some rows")
	}
	if err != nil {
		log.Fatalf("Could not apply bulk row mutation: %v", err)
		return err
	}
	return nil
}

func WriteToTokens(ctx context.Context, client *bigtable.Client, tokens []Token) error {
	tbl := client.Open(tableTokens)
	mutations := make([]*bigtable.Mutation, len(tokens))
	rowKeys := make([]string, len(tokens))

	for i, token := range tokens {
		mutations[i] = bigtable.NewMutation()
		mutations[i].Set(tokenColumnFamilyName, tokenColumnName, bigtable.Now(), []byte(token.Token))
		mutations[i].Set(tokenColumnFamilyName, accessLimitPerMinColumnName, bigtable.Now(), []byte(token.AccessLimitPerMin))
		rowKeys[i] = token.Token
	}
	rowErrs, err := tbl.ApplyBulk(ctx, rowKeys, mutations)
	if rowErrs != nil {
		for _, rowErr := range rowErrs {
			log.Printf("Error writing row: %v", rowErr)
		}
		log.Fatalf("Could not write some rows")
	}
	if err != nil {
		log.Fatalf("Could not apply bulk row mutation: %v", err)
		return err
	}
	return nil
}

func GetLongLinkByShortLink(ctx context.Context, client *bigtable.Client, key string) string {
	tbl := client.Open(tableLinks)
	row, err := tbl.ReadRow(ctx, key, bigtable.RowFilter(bigtable.ColumnFilter(longLinkColumnName)))
	if err != nil {
		log.Fatalf("Could not read row with key %s: %v", key, err)
	}
	if row == nil {
		return ""
	}
	fmt.Println(row[longLinkColumnName])
	return string(row[longLinkColumnName][0].Value)
}
