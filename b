# go build -a -installsuffix cgo -o main .
go build -a -installsuffix cgo -o myserver server/main.go
go build -a -installsuffix cgo -o myclient client/main.go
go build -a -installsuffix cgo -o mylocal cmd/main.go 
echo '••• Compilou os fontes e gerou os executáveis mylocal, myclient e myserver •••'
