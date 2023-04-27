# Pawsitively Purrfect

## Introduction
**_Pawsitively Purrfect_** is a project aimed at providing an online platform for pet adoption. This platform is designed to help pet lovers find their perfect furry companions while also providing a way for animal shelters and rescue organizations to showcase their available pets.

The "Pawsitively Purrfect" platform provides a ~~user-friendly~~ interface for browsing pets based on different search criteria such as pet type, breed, location, and more. Users can view detailed information about each pet and contact the pet's owner or the organization responsible for its care.

In addition to helping pets find their forever homes, "Pawsitively Purrfect" also provides a way for animal shelters to manage their available pets and adoption applications. With this platform, they can easily add new pets, update pet information, and track adoption applications.

## Architecture
### System Components


### System Design


### Data Flow

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
