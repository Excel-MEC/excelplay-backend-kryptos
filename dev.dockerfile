FROM golang:1.13
EXPOSE 8080
WORKDIR /excelplay-backend-kryptos
COPY . .
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ENTRYPOINT CompileDaemon -directory="." -log-prefix=false -build="go build /excelplay-backend-kryptos/cmd/excelplay-backend-kryptos" -command="./excelplay-backend-kryptos"