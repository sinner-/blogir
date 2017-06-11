# blogir
Golang weblog platform

## Development environment setup
### Fedora 24

As root:
  * `dnf install golang mariadb-server`
  * `mysql -u root`
    * `CREATE DATABASE blogir;`
    * `USE blogir;`
    * `CREATE TABLE posts(title VARCHAR(255) NOT NULL, body LONGTEXT NOT NULL)`;
  
As user:
  * `mkdir -p ~/Development/go/src/github.com/sinner-`
  * `export GOPATH=$HOME/Development/go`
  * `export PATH=$PATH:$HOME/Development/go/bin`
  * `git clone git@github.com:sinner-/blogir.git $GOPATH/src/github.com/sinner-/blogir`
  * `cd $GOPATH/src/github.com/sinner-/blogir`
  * `go get`
  * `go install`
  * Optionally, edit blogirrc env vars and `source blogirrc`
  * `blogir`
