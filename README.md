# User-Management


####ABOUT
This is a golang library which allows you to maintain your users.
All users are mapped to a specific role. Based on his/her role the apis are acccessbile.
For example, only admin users are allowed to fetch any users' details and can update the role of any user except his own.

#### Run Locally
```bash
1) Install golang
2) mkdir -p $HOME/go/{bin,src} 
3) Set following in .bash_profile 
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin
4) Clone user-management
5) docker-compose up
6) go run /cmd/main.go -file=local.json
```

#### Import Locally
```
go get -u github.com/tiwariayush700/user-management
```

### APIS CAN BE TESTED FROM THE
 `/apiTestLocal directory`