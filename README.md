# blogir
Golang weblog platform

## Development environment setup
### Fedora 27

As root:
  * `dnf install golang mariadb-server`
  * `mysql -u root`
    * `CREATE DATABASE blogir;`
  
As user:
  * `mkdir -p ~/Development/go/src/github.com/sinner-`
  * `export GOPATH=$HOME/Development/go`
  * `export PATH=$PATH:$HOME/Development/go/bin`
  * `git clone git@github.com:sinner-/blogir.git $GOPATH/src/github.com/sinner-/blogir`
  * `cd $GOPATH/src/github.com/sinner-/blogir`
  * `go get`
  * `go install`
  * Set the BLOGIR_ADMIN_PASSHASH environment variable with a bcrypt hashed password. blogirrc has an example password of "password".
  * Optionally, edit blogirrc env vars and `source blogirrc`.
  * `blogir`
