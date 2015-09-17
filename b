# go build -a -installsuffix cgo -o main .
echo '••• Gerando myserver •••'
go build -a -installsuffix cgo -o myserver server/main.go
echo '••• Gerando myclient •••'
go build -a -installsuffix cgo -o myclient client/main.go
echo '••• Gerando mylocal •••'
go build -a -installsuffix cgo -o mylocal cmd/main.go 
echo '••• Compilou os fontes e gerou os executáveis mylocal, myclient e myserver •••'
