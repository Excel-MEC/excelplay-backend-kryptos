FROM golang:1.12
WORKDIR /excelplay-backend-kryptos
COPY . .
WORKDIR /excelplay-backend-kryptos/
CMD ["go","run","."]