# larks_go

This will be a basic template which we will use for golang based backend projects. 


# Project structure

```
project_name
│   README.md -> (This file :))
│   Routes    -> defines URL to interface methods (Something like what Play framework provides)
│   main.go   -> Starting point of app, will start backend go server, connect to db and load routes and map to necessary interface.
|   Procfile  -> To start both frontend and backend application servers. 
|
└───/bin
│   │   build.sh -> Builds both go app and react app to production binaries.
└───/db
|    │   db.go    -> Establishes connection to the database (only PG for now).
└───/docs
|    │   generate_docs.go  -> WIP (will generate documentation about interface methods)
└───/interfaces
|    │   project.go  -> Interface to handle all Project related requests
└───/managers
|    │   base_manager.go
|    |   projects_manager.go
└───/objects
|    │   base.go
|    |   project.go
└───/web
|    │   base.go
|    |   project.go
```

Internally Go's routing is handled by the awesome GIN framework, just tweaked to meet my preferences.
  
