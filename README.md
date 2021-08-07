# nuxt-gin-todolist
## Environment
* Docker: v20.10.7
* Go: v1.16.4
* Node.js: v14.17.1

## Command
```bash
# Set runtime enviroment
make setEnv

# Launch go server & db
make runbackend

# Launch web dev-server
make runweb

# Launch db
make rundb

# Stop db
make stopdb
```

## Development

### Develop gin server
```bash
# run first time or enviroment changed
make setEnv

# Launch go server & db
make runbackend
```

### Develop web dev-server

```bash
# run first time or enviroment changed
make setEnv

# Launch go server & db
make runweb
```
