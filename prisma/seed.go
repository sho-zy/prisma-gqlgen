package main

import (
	"context"

	prisma "workspace/prisma-gqlgen/prisma-client"
)

func main() {

	client := prisma.New(nil)
	ctx := context.Background()

	published := true

	email1 := "alice@prisma.io"
	name1 := "Alice"
	title1 := "Join us for Prisma Day 2019 in Berlin"
	client.CreateUser(prisma.UserCreateInput{
		Name:  name1,
		Email: &email1,
		Posts: &prisma.PostCreateManyWithoutAuthorInput{
			Create: []prisma.PostCreateWithoutAuthorInput{
				prisma.PostCreateWithoutAuthorInput{
					Title:     title1,
					Published: &published,
				},
			},
		},
	}).Exec(ctx)

	email2 := "bob@prisma.io"
	name2 := "Bob"
	title2 := "Subscribe to GraphQL Weekly for community news"
	title3 := "Follow Prisma on Twitter"
	client.CreateUser(prisma.UserCreateInput{
		Name:  name2,
		Email: &email2,
		Posts: &prisma.PostCreateManyWithoutAuthorInput{
			Create: []prisma.PostCreateWithoutAuthorInput{
				prisma.PostCreateWithoutAuthorInput{
					Title:     title2,
					Published: &published,
				},
				prisma.PostCreateWithoutAuthorInput{
					Title:     title3,
					Published: &published,
				},
			},
		},
	}).Exec(ctx)
}
