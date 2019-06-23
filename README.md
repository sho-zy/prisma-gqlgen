# prisma-gqlgen
## Set up
- Prisma Client
  > `git clone https://github.com/sho-zy/prisma-gqlgen.git`  
  > `dep ensure`  
  > `docker-compose up -d`  
  > `prisma deploy`  

- Prisma App
  > `dep ensure -update`  
  > `go run ./scripts/gqlgen.go`  
  > `go run ./server`  

## URLs
- Prisma Admin  
  http://localhost:4466

- GraphQL Server  
  http://localhost:4000

## Commands
- Update and deploy the Prisma Client.  
  ( execute when you changed `./prisma/datamodel.prisma` )  
  > `prisma deploy`

  after, you can use the Prisma Admin.

- Start the GraphQL server. 
  > `go run ./server`  

  after, you can use the GraphQL server.

- Update the GraphQL schema.  
  ( execute when you changed `./server/schema.graphql` )  
  > `go run ./scripts/gqlgen.go`

## Cases
- Case when you add tables.
  1. Add table definition to the following modules.  
   `./prisma/datamodel.prisma`  
   `./server/schema.graphql`  
   `./gqlgen.yml`  
  2. execute the following commands.  
        > `rm ./tmp/resolver.go`  
        > `prisma deploy`  
        > `go run ./scripts/gqlgen.go`  
  3. edit the resolver ( `./server/resolver.go` ) referring to `/tmp/resolver.go`.  
  4. execute the following commands.  
        > `go run ./server`  
