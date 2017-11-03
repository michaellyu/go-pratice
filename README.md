Go Pratice
==========

### Start Docker

```shell
sudo docker-compose up -d
```

### Install Go Packages

```shell
cd ./go/src
./install_pkgs.sh
```
#### Or

```shell
./get github.com/someone/somepkg
```

### Run Code

```shell
./run helloworld.go
```

### Build Code

```shell
./build helloworld.go
```

### End

```shell
sudo docker-compose down
```
