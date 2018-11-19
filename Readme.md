# Installation (Windows 10)
### Protocol Buffers v3 (used to generate gRPC service code)

- The simplest way to do this is to download pre-compiled binaries for your platform(protoc-<version>-<platform>.zip) from here: https://github.com/google/protobuf/releases.
- Unzip this file.
- Update the environment variable PATH to include the path to the protoc binary file. (edit the system environment variable).
    > variable and system - `C:\Program Files (x86)\protoc-3.6.1-win32\bin`

### GnuWin32 (make command)
-   Download from here (http://gnuwin32.sourceforge.net/packages/make.htm)
-   Update the environment variable PATH - `C:\Program Files (x86)\GnuWin32\bin`

### Install following libraries

-   `tutorial-part-1 : Basics of writing a gRPC based microservice`

    ```
    gRPC Package - go get -u google.golang.org/grpc
    Go-micro dependencies - go get -u github.com/golang/protobuf/protoc-gen-go
    ```

-   `tutorial-part-2 : Docker and go-micro`

    ```
    Go-micro - go get github.com/micro/micro
    Cross Compilation (linux) - go get golang.org/x/sys/unix
    ```

- `tutorial-part-3 : docker-compose and datastores`

    ```
    Object-relational mapping (ORM) library for golang. - `go get -u github.com/jinzhu/gorm`
    Go implementation of Universally Unique Identifier (UUID). - `go get github.com/satori/go.uuid`
    ```

# Microservice
### Why?
- Complexity - Splitting features into microservices allows you to split code into smaller chunks. It harks back to the old unix adage of 'doing one thing well'. There's a tendency with monoliths to allow domains to become tightly coupled with one another, and concerns to become blurred. This leads to riskier, more complex updates, potentially more bugs and more difficult integrations.

- Scale - In a monolith, certain areas of code may be used more frequently than others. With a monolith, you can only scale the entire codebase. So if your auth service is hit constantly, you need to scale the entire codebase to cope with the load for just your auth service.

With microservices, that separation allows you to scale individual services individually. Meaning more efficient horizontal scaling. Which works very nicely with cloud computing with multiple cores and regions etc.

### Issues
##### Communication
- Because microservices are split out into separate codebases, one important issue with microservices, is communication.
- In a monolith communication is not an issue, as you call code directly from elsewhere in your codebase. However, microservices don't have that ability, as they live in separate places. So you need a way in which these independent services can talk to one another with as little latency as possible.
- Why not use REST? (JSON or XML over http).
    > Process encode and decode between services has a potential overhead problems at scale.

# Acknowledgment
Credit goes to [Ewan Valentine](https://ewanvalentine.io/tag/golang/) for the awesome and insightful tutorial.
