# pawsitively-purrfect
Pawsitively Purrfect


Project Structures:
```
- .github/
- cmd/
- configs/
    - .pawsitively-purrfect.yaml
- controller/
    # handler functions implementations
- infra/
    # nosql implementations
    - mongo/
    - arango/

    # sql implementations
    - postgres/

    # DB interfaces
    - sql/
        - DB interface
    - nosql
        - DB interface
- routers
    - middlewares/
- repo
    # contains all the repositories for all the entities
- service
    # graphql schemas
- libs/

- templates/

- Dockerfile
- Makefile
```
