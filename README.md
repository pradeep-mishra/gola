# GOLA

Simple Web app framework for go lang

The Gola framework is based on Fiber, godotenv, and mongo library.

It's very opinionated on how code should be structured.

#### Prerequisite

Node.js v12+
Npm


#### Installation

```bash
npm install gola -g
```

### Usage 
To create a new project
```bash
gola init [project_name]
```

to add new controller in existing project

```bash
gola controller [controller_name]
```

to run the project

```bash
gola run
```

to build a project
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
└───[controller_name] (folder)
│   │
│   │__ controller.go (file)
│  
│
└───go.mod (file)
│  
│
│
└───go.sum (file) 


```


Work in progress
