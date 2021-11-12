# Gola CLI

Simple web app CLI for Go lang

The Gola CLI is based on the Fiber, Mongo, Validator and Godotenv library.

It uses the [Fiber](https://docs.gofiber.io/) web framework to manage the web app
and the [Mongo](https://docs.mongodb.com/drivers/go/current/fundamentals/connection/) package to handle the MongoDB connection
and the [Validator](https://github.com/go-playground/validator) to handle the validation of dtos
and [Godotenv](https://pkg.go.dev/github.com/joho/godotenv) for the configuration.


It's very opinionated on how the code should be structured.

#### Prerequisite

Node.js v12+


#### Installation

```bash
npm install gola -g
```

### Usage 
To create a new project
```bash
gola init [project_name]
```

To add new resource in existing project

```bash
gola resource [resource_name]
```

To run a project

```bash
gola run
```

To build a project
```bash
gola build
```


#### Code Structure

```
myapp (folder) - App folder
│
│
└───.env (file) - Environment config file
│
│
└───main.go (file) - Entry file
│
│ 
└───global (folder) - Package for global util functions
│   │
│   │__ global.go (file) 
│    
│
└───server (folder) - Package for core web server files
│   │
│   │__ server.go (file)
│
│
└───[resource_name] (folder) - Resource package. e.g. (user, project, content)
│   │
│   │__ controller.go (file) - All routing for resource goes here
│   │
│   │__ service.go (file) - Business logic to handle all route goes here 
│   │
│   │__ dto.go (file) - Dto validation object goes here
│
│
└───bin (folder) - Executable file would be placed here after build command
│   │
│   │__ myapp (file) - Executable
│  
│
└───go.mod (file) - Go module registry like package.json for node.js
│  
│
│
└───go.sum (file)  - Go module registry like package-lock.json for node.js


```

Example 

```bash
$ gola init myapp
$ cd myapp
$ gola resource user
$ gola run
```
After ```gola run``` command, follow this link to access the application

[http://localhost:3000/users](http://localhost:3000/users)


Work in progress
