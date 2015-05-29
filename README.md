![minigun](http://img2.wikia.nocookie.net/__cb20111124065924/mafiawars/images/b/b4/Gatlin_Minigun.png) 

## Minigun 

Simple tool to load your api with tons of requests 

```
go run minigun.go -file sampleRequest.xml -url http://my-server.com -rounds 100 
```

compile it for Windows and mail your friends  

```
GOOS=windows GOARCH=386 go build -o minigun.exe minigun.go 
```
