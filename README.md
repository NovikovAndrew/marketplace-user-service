# User service marketplace 

### Migration 

I use golang-migrate, for install macOS use command 

````
brew install golang-migrate
````

to create schema use command

````
 migrate create -ext sql -dir db/postgresql/migration -seq init_schema
````

for migrate up use command 

````
migrate -path db/postgresql/migration -database "postgresql://root:root@localhost:5432/marketplace?sslmode=disable" -verbose up
````

check file Makefile for more migration commands

### Install mockgen

make sure mockgen is install

```
ls -l ~/go/bin
```

for check mockgen user command 

````
which mockgen
````

for set mockgen use your go to your shell in my case

```
vi ~/.zshrch
```

after insert the path go/bin

```
export PATH=$PATH:~/go/bin
```

and

```
source ~/.zshrc
```

check file Makefile for more migration commands

### gPRC client for testing

```
brew tap ktr0731/evans
brew install evans
```

Docker image

```
docker run --rm -v "$(pwd):/mount:ro" \
    ghcr.io/ktr0731/evans:latest \
      --path ./proto/files \
      --proto file-name.proto \
      --host example.com \
      --port 50051 \
      repl
```

if you want to start evans use command 

```
evans -r repl
```

specific host and port

```
evans --tls --host example.com --port -r repl
```

check Makefile for more migration commands