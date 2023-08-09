

## init your go lang project

```
go mod init github.com/NubeIO/module-contrib-demo
```

## add a 3rd party dependency
```
go get github.com/NubeDev/bom-api
```

## rename module name as required

```golang
const name = "module-contrib-demo"
```

* [app.go](pkg/app.go)
* [main.go](main.go)
* [build.bash](build.bash)

### see naming rules
https://nubeio.github.io/rubix-ce-docs/docs/api-docs/modules/


## add an api

[api.go](pkg/api.go)


```golang
type helloWorld struct {
	A              string    `json:"a"`
	B              int       `json:"b"`
	C              bool      `json:"c"`
	TimeDateFormat string    `json:"time_date_format"`
	TimeDate       time.Time `json:"time_date"`
}

func (inst *Module) Get(path string) ([]byte, error) {
	if path == "ping" {
		return json.Marshal(helloWorld{
			A:              "ping",
			B:              0,
			C:              false,
			TimeDateFormat: time.Now().Format(time.Stamp),
			TimeDate:       time.Now().UTC(),
		})
	}
	
	return nil, errors.New(path)
}

```