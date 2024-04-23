# GOLANG GRAPHQL MONGODB CRUD Project

## GraphQL
GraphQL is a query language for APIs and a runtime for executing those queries by using a type system that you define for your data. It provides a more efficient, powerful, and flexible alternative to REST for building APIs.

Here is an overview of GraphQL and its key concepts:

### Key Concepts:
1. Schema: The schema defines the structure of your API, including the types of data that can be queried and the operations (queries, mutations, and subscriptions) that can be performed.
The schema is written using the GraphQL Schema Definition Language (SDL).

2. Queries: Queries are read-only operations that allow clients to request data from the server.
Clients can specify the exact data they need, and the server responds with only the requested data.
This fine-grained control helps minimize over-fetching and under-fetching of data.

3. Mutations: Mutations are operations that allow clients to modify data on the server (e.g., creating, updating, or deleting records).
Mutations are similar to queries in structure, but they represent changes to the data.

4. Subscriptions: Subscriptions allow clients to listen for real-time updates from the server.
Subscriptions establish a persistent connection between the client and the server, enabling the server to push data updates to the client.

5. Types: GraphQL supports a variety of types, including object types, scalar types (e.g., Int, String), lists, and custom types.
The schema defines the available types and their relationships.

6. Resolvers: Resolvers are functions that determine how to fetch or compute the data for a specific field in a query or mutation.
They map queries and mutations to the underlying data sources (e.g., databases, APIs).

7. Directives: Directives are special instructions that can be added to queries and schema definitions to control behavior (e.g., @deprecated, @include, @skip).

8. Introspection: GraphQL provides an introspection system that allows clients to query the API's schema and discover available types and operations.

##### Do the stuff below to initialize your project

1. Create a new folder for the Project
`mkdir gql-yt`
2. Mod init your project, give it whatever name you like
`go mod init github.com/Abhishekagarwal1597/Go-GraphQL-Project`
3. Get gql gen for your project
`go get github.com/99designs/gqlgen`
4. Add gqlgen to tools.go
`printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go`
`echo // +build tools > tools.go
echo package tools >> tools.go
echo import _ "github.com/99designs/gqlgen" >> tools.go
`
5. Get all the dependencies
`go mod tidy`
6. Initialize your project
`go run github.com/99designs/gqlgen init`
7. After you've written the graphql schema, run this - `go run github.com/99designs/gqlgen generate`
8. After you've built the project, these are the queries to interact with the API - 

#### Get All Jobs

`query GetAllJobs{
  jobs{
    _id
    title
    description
    company
    url
  }
}`

=======================

#### Create Job

`mutation CreateJobListing($input: CreateJobListingInput!){
  createJobListing(input:$input){
    _id
    title
    description
    company
    url
  }
}`

{
  "input": {
    "title": "Software Development Engineer - I",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt",
    "company": "Google",
    "url": "www.google.com/"
  }
}`


=========================

#### Get Job By Id

`query GetJob($id: ID!){
job(id:$id){
_id
title
description
url
company
}
}`


`{
  "id": "66268b6aecb75083e3286931"
}`



=========================


#### Update Job By Id

`mutation UpdateJob($id: ID!,$input: UpdateJobListingInput!) {
  updateJobListing(id:$id,input:$input){
    title
    description
    _id
    company
    url
  }
}`


`{
  "id": "638051d3acc418c13197fdf6",
  "input": {
    "title": "Software Development Engineer - III"
  }
}`

=================================


#### Delete Job By Id

`mutation DeleteQuery($id: ID!) {
  deleteJobListing(id:$id){
    deletedJobId
  }
}`

`{
  "id": "638051d3acc418c13197fdf6"
}`
