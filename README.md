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
mtools module install -m "modulus/chi-http"
```
or use just `mtools module install` and select `modulus/chi-http` module in the list.


If you want to use middlewares for your routes, you can use `mtools module install -m "modulus/chi-http/middleware"`,
