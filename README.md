# example module to show the use case of rubix-os (ROS)

This module will fetch data from the `Australian government bureau of meteorology` and write it to a point 

## to build and run

to build and run rubix-os you can use the bash script

`bash build.bash <YOUR_ROS_PATH>`

example
```
bash build.bash code/go
```



## init your go lang project

```
go mod init github.com/NubeIO/module-contrib-demo-test-test
```

## add a 3rd party dependency
```
go get github.com/NubeDev/bom-api
```

## rename module name as required

```golang
const name = "module-contrib-demo-test-test"
```

## renaming helpers

rename module name
```
find /home/aidan/code/go/module-contrib-demo -type f -print0 | xargs -0 sed -i 's#const name = "module-contrib-demo"#const name = "module-contrib-demo-test"#g'
```

rename library name
```
find /home/aidan/code/go/module-contrib-demo -type f -print0 | xargs -0 sed -i 's#github.com/NubeIO/module-contrib-demo-test#github.com/NubeIO/module-contrib-demo-test#g'
```

rename bash script
```
find /home/aidan/code/go/module-contrib-demo -type f -print0 | xargs -0 sed -i 's#module-contrib-demo#module-contrib-demo-test#g'
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

