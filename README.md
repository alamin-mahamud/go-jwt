# jwt authentication example in golang

POST http://localhost:12345/authenticate

``` json
{
	"username": "alamin-mahamud",
  	"password": "simple-password"
}
```

``` json
// Response
{
	"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsYW1pbi1tYWhhbXVkIn0.3pDO8lruVO2GAnABjknpMZK03XzsVktVBkNBmAbz-8I"
}

```
