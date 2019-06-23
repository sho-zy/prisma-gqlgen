package main

import (
	context "context"
	prisma "workspace/prisma-gqlgen/prisma-client"
)

// Resolver struct
type Resolver struct {
	Prisma *prisma.Client
}

// Mutation MutationResolver
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Post PostResolver
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}

// Query QueryResolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// User UserResolver
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*prisma.User, error) {
	return r.Prisma.CreateUser(prisma.UserCreateInput{
		Name: name,
	}).Exec(ctx)
}
func (r *mutationResolver) CreateDraft(ctx context.Context, title string, userID string) (*prisma.Post, error) {
	return r.Prisma.CreatePost(prisma.PostCreateInput{
		Title: title,
		Author: &prisma.UserCreateOneWithoutPostsInput{
			Connect: &prisma.UserWhereUniqueInput{ID: &userID},
		},
	}).Exec(ctx)
}
func (r *mutationResolver) Publish(ctx context.Context, postID string) (*prisma.Post, error) {
	published := true
	return r.Prisma.UpdatePost(prisma.PostUpdateParams{
		Where: prisma.PostWhereUniqueInput{ID: &postID},
		Data:  prisma.PostUpdateInput{Published: &published},
	}).Exec(ctx)
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (*prisma.User, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &obj.ID}).Author().Exec(ctx)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) PublishedPosts(ctx context.Context) ([]prisma.Post, error) {
	published := true
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{Published: &published},
	}).Exec(ctx)
}
func (r *queryResolver) Post(ctx context.Context, postID string) (*prisma.Post, error) {
	return r.Prisma.Post(prisma.PostWhereUniqueInput{ID: &postID}).Exec(ctx)
}
func (r *queryResolver) PostsByUser(ctx context.Context, userID string) ([]prisma.Post, error) {
	return r.Prisma.Posts(&prisma.PostsParams{
		Where: &prisma.PostWhereInput{
			Author: &prisma.UserWhereInput{
				ID: &userID,
			}},
	}).Exec(ctx)
}
func (r *queryResolver) Users(ctx context.Context) ([]prisma.User, error) {
	return r.Prisma.Users(&prisma.UsersParams{
		Where: &prisma.UserWhereInput{},
	}).Exec(ctx)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &obj.ID}).Posts(nil).Exec(ctx)
}
