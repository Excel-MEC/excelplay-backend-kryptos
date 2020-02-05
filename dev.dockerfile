FROM golang:1.12
WORKDIR /excelplay-backend-kryptos
COPY . .
WORKDIR /excelplay-backend-kryptos/src
CMD ["go","run","."]