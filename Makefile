.PHONY: build

build:
	set GOOS=linux
	set GOARCH=amd64
	set CGO_ENABLED=0
	sam build

wire:
	wire di/wire.go

uml:
	gouml init -o ./tmp/app.uml
	cat ./tmp/app.uml | gsed  '1i @startuml' | gsed '$a @enduml' > ./tmp/app.pu
	plantuml -tpng ./tmp/app.pu
