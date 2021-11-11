# Gola

Simple Web app framework for go lang

The Gola framework is based on Fiber, godotenv, and mongo library.

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
myapp (folder)
│
│
└───.env (file) - Environment config file
│
│
└───main.go (file) - Main entry file
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
│   │__ controller.go (file)
│   │
│   │__ service.go (file)
│   │
│   │__ dto.go (file)
│
│
└───bin (folder) - Executable file would be placed here after build command
│   │
│   │__ myapp (executable file)
│  
│
└───go.mod (file) - Go module registry like package.json for node.js
│  
│
│
└───go.sum (file)  - Go module registry like package-lock.json for node.js


```


Work in progress
