package main


var Schema = `
	type User {
		id: ID!
		first_name: String!
		last_name: String!
		avatar: String!
	}


	input UserInput {
		first_name: String!
		last_name: String!
		avatar: String!
	}
	
	schema {
		query: Query
		mutation: Mutation
	}

	type Empty {}

	type Mutation {
		createUser(user: UserInput!) : Empty!
		updateUser(user: UserInput!): Empty!
		deleteUser(userId: ID!) : Empty!
	} 

	type Query {
#		Users(): [User]!
#		User(): UserInput!
		User(id: ID!): User!
#		User(): User
	}
`