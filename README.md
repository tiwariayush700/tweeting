# User-Management


####ABOUT
This is a golang library which allows you to maintain your users.
All users are mapped to a specific role. Based on his/her role the apis are acccessbile.
For example, only admin users are allowed to fetch any users' details and can update the role of any user except his own.

######ACCESS FOR NORMAL USERS
Normal User can create tweets. Read their own tweets

######ACCESS FOR ADMIN USERS
ADMIN user can access all data and edit the tweets and user details
They can create an action which will be subject to approval from super admin

######ACCESS FOR SUPER ADMIN USERS
SUPER ADMIN users can view all the action logs created by admins and approve or reject the actions


#### Run Locally
```bash
1) Install golang
2) mkdir -p $HOME/go/{bin,src} 
3) Set following in .bash_profile 
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin
4) Clone tweeting
5) docker-compose up
6) go run /cmd/main.go -file=local.json
```

#### Import Locally
```
go get -u github.com/tiwariayush700/tweeting
```

### APIS CAN BE TESTED FROM THE
 `/apiTestLocal directory`