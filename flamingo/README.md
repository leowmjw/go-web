# Flamingo Project Structure

## Details

https://docs.flamingo.me/2.%20Flamingo%20Core/1.%20Flamingo%20Basics/2.%20Flamingo%20Project%20Structure.html

##  Project structure
A typical Flamingo project looks like this:
```
projectName (Project Root)
│   main.go (The entry for your project)
│   README.md
│   Dockerfile
│   Makefile
│   (Jenkinsfile or other CI config)
│   go.mod
│   go.sum
│
└───config (Your project configuration)
│   └───config.yml (Main project config)
│   └───config_dev.yml (additional configs - e.g. this one is loaded for CONTEXT dev)
│   └───routes.yml (Routing config)
│   └───SUBFOLDER (Optional additional configuration context)
│          └───config.yml (Additional config for this context)       
│   
└───src (Project specific modules live here)
│   └───myModule (a module - see module structure below)
│
└───frontend (Frontend templates - if "flamingo-carotene" is used)
│   └───src (main frontend source / structure by atomic design)
│   │    └───atom (see flamingo-carotene)
│   │    └───molecule
│   │    └───...
│   └───dist (not part of VCS - will have frontend build result)
```

## Training

https://docs.flamingo.me/6.%20Trainings/Workshop%20August.html
