# Pawsitively Purrfect

## Introduction
**_Pawsitively Purrfect_** is a project aimed at providing an online platform for pet adoption. This platform is designed to help pet lovers find their perfect furry companions while also providing a way for animal shelters and rescue organizations to showcase their available pets.

The "Pawsitively Purrfect" platform provides a ~~user-friendly~~ interface for browsing pets based on different search criteria such as pet type, breed, location, and more. Users can view detailed information about each pet and contact the pet's owner or the organization responsible for its care.

In addition to helping pets find their forever homes, "Pawsitively Purrfect" also provides a way for animal shelters to manage their available pets and adoption applications. With this platform, they can easily add new pets, update pet information, and track adoption applications.

## Architecture
### System Components
The System consists of three major Components mainly.

1. Backend Server: The backend server is built using Golang and GraphQL, with the Flamego web framework. It consists of multiple layers, including:
    - Resolver: Handles incoming GraphQL queries and mutations, and maps them to specific service functions.
    - Service: Implements business logic and interacts with repositories.
    - Repository: Provides an abstraction layer between the service layer and the database layer.
2. Databases: The system uses both NoSQL and SQL databases. ArangoDB is used as a NoSQL database, and Postgres is used as a SQL database, implemented using gRPC.
3. Frontend: The frontend is built using Tailwind CSS and includes basic pages such as login, register, and profile pages.

### System Design

The system uses a layered architecture that separates concerns and ensures loose coupling between the different components. It includes the following layers:

- Frontend: The frontend layer provides the user interface for the system and the main interface is yet to be built.
- GraphQL API: The GraphQL API layer handles incoming requests from the frontend and translates them into queries and mutations that can be executed by the backend.
- Resolver: The resolver layer maps the incoming GraphQL requests to their corresponding service methods.
- Service: The service layer contains the business logic of the system and performs the necessary operations on the data.
- Repository: The repository layer provides an abstraction layer over the database and handles the storage and retrieval of data.
- Database: The database layer stores the data used by the system and is responsible for ensuring its consistency and integrity.

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

### Data Model
- Description of the data model used in the project
- Explanation of the different entities and their relationships
- Schema and sample data for each entity

### API Documentation
- Overview of the API
- Description of the different API endpoints
- Parameters, query strings, and headers required for each endpoint
- Sample requests and responses for each endpoint

### Codebase
- Overview of the codebase
- Description of the different modules and packages used in the project
- Explanation of the different functions and their roles in the system
- Explanation of any notable design patterns or frameworks used in the project

### Deployment
- Explanation of how to deploy the project to a production environment
- Description of the hosting and infrastructure requirements

### Maintenance and Support
- Explanation of how to maintain and support the project in a production environment
- Information on how to debug and troubleshoot common issues
- Contact information for support and feedback

### Conclusion
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
