# http
HTTP module for the Modulus framework. It is based on the Chi router.

How to use:
1. Install mtools
```bash
go install github.com/go-modulus/mtools/cmd/mtools@latest
```
2. Init a project
```bash
mtools init
```
3. Add http module
```bash
mtools module install -m "modulus/chihttp"
```
or use just `mtools module install` and select `modulus/chihttp` module in the list.

Add a module option OverrideHttpRouter to the http module constructor inside the entrypoint of your application.

For example, in `/cmd/conosle/main.go`
```go
// DO NOT Remove. It will be edited by the `mtools module create` CLI command.
	modules := []*module.Module{
		http.NewModule(chihttp.OverrideHttpRouter),
...
}
```
