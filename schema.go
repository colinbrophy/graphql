package graphql


var Schema = `
	type User {
		id: ID!
		firstName: string!
		lastName: string!
	}

	input UserInput {
		firstName: string!
		lastName: string!
	}
	
	schema {
		query: Query
		mutation: Mutation
	}

	type Mutation {
		createUser(user: !): Person
		updateUser(crewNo: ID!, person: PersonInput!): Person
		deleteUser(crewNo: ID!): Person
	}

	type Query {
		User(): [User]
		Person(id: ID!): Person
	}
`