FROM golang:1.13
EXPOSE 8080
WORKDIR /excelplay-backend-kryptos
COPY . .
RUN ["go", "build", "/excelplay-backend-kryptos/cmd/excelplay-backend-kryptos"]
ENTRYPOINT ./excelplay-backend-kryptos