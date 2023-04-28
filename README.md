# Pawsitively Purrfect

## Introduction
**_Pawsitively Purrfect_** is a project aimed at providing an online platform for pet adoption. This platform is designed to help pet lovers find their perfect furry companions while also providing a way for animal shelters and rescue organizations to showcase their available pets.

The "Pawsitively Purrfect" platform provides a ~~user-friendly~~ interface for browsing pets based on different search criteria such as pet type, breed, location, and more. Users can view detailed information about each pet and contact the pet's owner or the organization responsible for its care.

In addition to helping pets find their forever homes, "Pawsitively Purrfect" also provides a way for animal shelters to manage their available pets and adoption applications. With this platform, they can easily add new pets, update pet information, and track adoption applications.

## Architecture
### System Components
The System consists of three major Components mainly.

1. **_Backend Server_**: The backend server is built using Golang and GraphQL, with the Flamego web framework. It consists of multiple layers, including:
    - **_Resolver_**: Handles incoming GraphQL queries and mutations, and maps them to specific service functions.
    - **_Service_**: Implements business logic and interacts with repositories.
    - **_Repository_**: Provides an abstraction layer between the service layer and the database layer.
2. **_Databases_**: The system uses both NoSQL and SQL databases. ArangoDB is used as a NoSQL database, and Postgres is used as a SQL database, implemented using gRPC.
3. **_Frontend_**: The frontend is built using Tailwind CSS and includes basic pages such as login, register, and profile pages.

### System Design

The system uses a layered architecture that separates concerns and ensures loose coupling between the different components. It includes the following layers:

- **_Frontend_**: The frontend layer provides the user interface for the system and the main interface is yet to be built.
- **_GraphQL API_**: The GraphQL API layer handles incoming requests from the frontend and translates them into queries and mutations that can be executed by the backend.
- **_Resolver_**: The resolver layer maps the incoming GraphQL requests to their corresponding service methods.
- **_Service_**: The service layer contains the business logic of the system and performs the necessary operations on the data.
- **_Repository_**: The repository layer provides an abstraction layer over the database and handles the storage and retrieval of data.
- **_Database_**: The database layer stores the data used by the system and is responsible for ensuring its consistency and integrity.

This layered architecture enables the system to be easily extensible and maintainable, as changes to one layer do not affect the others.

### Data Flow
Data flows through our system as follows:

1. The frontend client sends a request to the backend server using GraphQL.
2. The request is received by the GraphQL API layer, which validates and parses the request.
3. The request is then passed to the resolver layer, which maps the request to the corresponding service methods.
4. The service layer contains the business logic of the system and performs the necessary operations on the data.
5. The repository layer provides an abstraction layer over the database and handles the storage and retrieval of data.
6. The database layer stores the data used by the system and is responsible for ensuring its consistency and integrity.
7. The requested data is retrieved from the database and returned to the service layer.
8. The service layer processes the data and returns the response to the resolver layer.
9. The resolver layer maps the response to the GraphQL schema and returns it to the GraphQL API layer.
10. The GraphQL API layer sends the response back to the frontend client.

## Data Model
The "Pawsitively Purrfect" project uses a relational data model to store information about users, pets, shelters, and pet adoptions. There's also option to use NoSQL type database instead of SQL type database.

### Entities and Relationships
There are four main entities in the data model: User, Pet, Shelter, and Pet Adoption. The entities have various attributes that define them, and have different relationships between them.

Here's a breakdown of the different entities and their relationships:

- User Entity:<br/>
    The User entity represents the users of the platform. User has various attributes such as their full name, email, username, location etc. Each User can adopt multiple Pets from a Shelter which is also owned by a User.

- Shelter Entity:<br/>
    The Shelter entity represents the shelters on the platform. Shelter has various attributes such as its name, description, contact information etc. Each shelter can have only one owner, represented by the OwnerID attribute. Each shelter can own multiple pets.

- Pet Entity:<br/>
    The Pet entity represents the pets on the platform. Pet has various attributes such as its name, type, breed, adoption status etc. The AdoptionStatus attribute represents whether the pet is available for adoption or has already been adopted. Each pet belongs to one shelter, represented by the ShelterID attribute. However, a pet can be adopted by a user, which is represented by a PetAdoption record.

- PetAdoption Entity:<br/>
    The PetAdoption entity represents the adoption of pets by users. Each PetAdoption record has a unique ID and is associated with a specific pet and a specific user.

### Schema and Sample Data
JSON Schemas for the entities used in the "Pawsitively Purrfect" project:

User:
```json
{
  "type": "object",
  "properties": {
    "id": {"type": "string"},
    "firstName": {"type": "string"},
    "lastName": {"type": "string"},
    "bio": {"type": "string"},
    "location": {"type": "string"},
    "avatar": {"type": "string"},
    "username": {"type": "string"},
    "email": {"type": "string"},
    "passwordHash": {"type": "string"},
    "isActive": {"type": "boolean"},
    "isAdmin": {"type": "boolean"},
    "createdUnix": {"type": "integer"},
    "updatedUnix": {"type": "integer"},
    "lastLoginUnix": {"type": "integer"}
  }
}
```

Shelter:
```json
{
  "type": "object",
  "properties": {
    "id": {"type": "string"},
    "name": {"type": "string"},
    "description": {"type": "string"},
    "website": {"type": "string"},
    "location": {"type": "string"},
    "contactInformation": {"type": "string"},
    "logo": {"type": "string"},
    "numberOfPets": {"type": "integer"},
    "ownerID": {"type": "string"}
  }
}
```

Pet:
```json
{
  "type": "object",
  "properties": {
    "id": {"type": "string"},
    "name": {"type": "string"},
    "type": {"type": "string"},
    "breed": {"type": "string"},
    "gender": {"type": "string"},
    "photo": {"type": "string"},
    "adoptionStatus": {"type": "string"},
    "shelterID": {"type": "string"}
  }
}
```

PetAdoption:
```json
{
  "type": "object",
  "properties": {
    "id": {"type": "string"},
    "petID": {"type": "string"},
    "userID": {"type": "string"}
  }
}
```

And here's an example of some sample data that could be inserted into these tables:

User:
| ID | FirstName | LastName | Bio | Location      | Avatar | Username | Email                | PasswordHash | IsActive | IsAdmin | CreatedUnix | UpdatedUnix | LastLoginUnix |
|----|-----------|----------|-----|---------------|--------|----------|----------------------|--------------|----------|---------|-------------|-------------|---------------|
| 1  | John      | Doe      | I love dogs! | New York, NY | john.jpg | johnd  | john@pawsshelter.com | 1234567890   | true     | false   | 1620187800  | 1620187800  | 1620187800    |
| 2  | Jane      | Smith    | I am a cat lover! | Los Angeles, CA | jane.jpg | janes  | jane@furryfriends.org | 0987654321   | true     | false   | 1620187800  | 1620187800  | 1620187800    |
| 3  | Bob       | Johnson  | I love all animals! | Seattle, WA | bob.jpg  | bobj   | bob@happytailsrescue.org | 1357908642 | true     | true    | 1620187800  | 1620187800  | 1620187800    |


Shelter:
| ID | Name  | Description | Website | Location | ContactInformation | Logo | NumberOfPets | OwnerID |
|----|-------|-------------|---------|----------|--------------------|------|--------------|---------|
| 1  | PAWS  | Animal rescue organization | www.pawsshelter.com | New York, NY | contact@pawsshelter.com | pawsshelterlogo.jpg | 100 | 1 |
| 2  | Furry Friends | Purrfect home for your furry friends | www.furryfriends.org | Los Angeles, CA | contact@furryfriends.org | furryfriendslogo.jpg | 75 | 2 |
| 3  | Happy Tails  | Helping pets find their forever homes | www.happytailsrescue.org | Seattle, WA | contact@happytailsrescue.org | happytailslogo.jpg | 50 | 3 |


Pet:
| ID | Name | Type | Breed | Gender | Photo | AdoptionStatus | ShelterID |
|----|------|------|-------|--------|-------|----------------|----------|
| 1  | Max  | Dog  | Labrador Retriever | Male   | maxphoto.jpg  | Available    | 1 |
| 2  | Bella | Cat  | Siamese           | Female | bellaphoto.jpg | Adopted        | 1 |
| 3  | Lucy | Dog   | Bulldog          | Female | lucyphoto.jpg  | Available    | 2 |
| 4  | Simba | Cat  | Maine Coon       | Male   | simbaphoto.jpg | Available    | 2 |
| 5  | Charlie | Dog  | Golden Retriever | Male  | charliephoto.jpg | Adopted     | 3 |


PetAdoption:
| ID | PetID | UserID |
|----|-------|--------|
| 1  | 2     | 1      |
| 2  | 5     | 2      |
| 3  | 3     | 1      |


## API Documentation

The Pawsitively Purrfect is a GraphQL server that provides a set of endpoints for managing users, pets, shelters, and pet adoptions. The API allows clients to query and mutate data using a single GraphQL endpoint.

### Endpoints
The Pawsitively Purrfect API has the following endpoints:

#### Rest APIs
The  server some Rest APIs to view the website documentation, login, register and view profile page from UI.

- `/docs` - Serves the GraphQL documentation
- `/user/login` - Serves Login Page
- `/user/register` - Serves Registration Page

- `/logout` - Logout page
- `/{name}` - User profile page

#### GraphQL API
- `/graphql` - The main endpoint for the GraphQL server.

- **_You can find the full GraphQL documentation in_** <a href="https://pawsitively-purrfect.tiiny.site" target="_blank">**_pawsitively-purrfect.tiiny.site_**</a>.

## Codebase
- Overview of the codebase
- Description of the different modules and packages used in the project
- Explanation of the different functions and their roles in the system
- Explanation of any notable design patterns or frameworks used in the project

## Deployment
- Explanation of how to deploy the project to a production environment
- Description of the hosting and infrastructure requirements

## Maintenance and Support
- Explanation of how to maintain and support the project in a production environment
- Information on how to debug and troubleshoot common issues
- Contact information for support and feedback

## Conclusion
- Summary of the project
- Future plans and enhancements

( Add support for adding must columns in sql db implementations.
It's needed if we want to update column with zero value. )

<details>
<summary>
Project Structures:
</summary>

```
- .github/
    - workflows/
        - ci.yml
        - release.yml
- api/
    - graphql/
        - resolvers/
            - resolver.go
            - pet.go
            - shelter.go
            - user.go
        - schema/
            - params.go
            - schema.go
    - http/
        - handlers/
            - graphql.go
            - user.go
        - middlewares/
            - auth.go
            - context.go
            - rate_limit.go
            - session_csrf.go
        - routes.go
- cmd/
    - config.go
    - root.go
    - serve.go
- configs/
    - .pawsitively-purrfect.yaml
    - config.go
- infra/
    - database/
        - nosql/
            - arangodb/
                - arangodb.go
            - mongodb/
            - mock/
                - mock.go
            - database.go
        - sql/
            - postgres/
            - database.go
    - logr/
        - logger.go
- models/
    - gqtypes/
        - pet.go
        - shelter.go
        - types.go
        - user.go
    - errors.go
    - pet.go
    - shelter.go
    - user.go
- pkg/
    - decode.go
    - hash.go
    - path.go
- repos/
    - pet/
        - pet.go
        - pet_mock.go
    - shelter/
        - shelter.go
        - shelter_mock.go
    - user/
        - user.go
        - user_mock.go
    - pet.go
    - shelter.go
    - user.go
- services/
    - all/
        - all.go
    - pet/
        - pet.go
    - shelter/
        - shelter.go
    - user/
        - user.go
    - pet.go
    - shelter.go
    - user.go
- templates/
    - login.tmpl
    - register.tmpl
    - profile.tmpl
    - templates.go
- main.go

- docker-compose.yml
- Dockerfile
- Makefile
- README.md
```
</details>
