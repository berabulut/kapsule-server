# MongoDB Installation and Configuration on Ubuntu Instance 


## Installation

Follow official [documentation](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/).


## Configuration 

 - Create admin user

	```
	> use admin
	> db.createUser( 
		{ user: "username",
		  pwd: "password", 
		  roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ] },
		  { w: "majority" , wtimeout: 5000 } )
	```
- Create normal user

	```
	> use admin
	> db.createUser( 
		{ user: "username",
		  pwd: "password", 
		  roles: [  { role: "readWrite", db: "kapsule" } ] },
		  { w: "majority" , wtimeout: 5000 } )
	```

- Enable remote access ( /etc/mongod.conf )

	Change **bindIp** to **0.0.0.0**  

- Enable authorization ( /etc/mongod.conf )

	```
	security:
	  authorization: enabled
	```
- Restart MongoDB 

	`sudo systemctl restart mongod`

- Check MongoDB status

	`sudo systemctl status mongod`

- Access mongo shell with admin privileges when you need it

	`mongo -u "username" -p "password"`