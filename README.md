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
└───.env (file)
│
│
└───main.go (file)
│
│ 
└───global (folder)
│   │
│   │__ db.go (file)
│    
│
└───server (folder)
│   │
│   │__ server.go (file)
│
│
└───[resource_name] (folder)
│   │
│   │__ controller.go (file)
│   │
│   │__ service.go (file)
│   │
│   │__ dto.go (file)
│
│
└───bin (folder)
│   │
│   │__ myapp (executable file)
│  
│
└───go.mod (file)
│  
│
│
└───go.sum (file) 


```


Work in progress
