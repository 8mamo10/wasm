# wasm
build
```
GOOS=js GOARCH=wasm go build -o test.wasm test_wasm.go
```
make symlink
```
$ sudo ln -s /home/vagrant/git/wasm/test.wasm /var/www/html
$ sudo ln -s /home/vagrant/git/wasm/wasm.html /var/www/html
$ sudo ln -s /home/vagrant/git/wasm/wasm_exec.js /var/www/html
```
add mime type
```
$ tail -n1 /etc/mime.types
application/wasm wasm
$ sudo apache2ctl restart
```
