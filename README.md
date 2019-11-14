# Face Registry

This micro-service was designed to keep a registry of people whose faces can 
be recognized when a picture of a face is sent to it.

Each person can have multiple faces stored (to account for different angles, for example).

Communication with the server is done through *grpc*.

The *API* contains the following methods:
* **Recognize**
* **AddFace**
* **RemoveFace**
* **GetPeople**
* **AddPerson**
* **RemovePerson**

Check the proto files for more info.

## Start the service
Simply run `docker-compose up -d` and you're done!

## Notes
* Right now, the service takes a while to get started. (Around a minute)
I suspect this is related to the library I'm using. 
I'll see if there's a way to make it start faster.

* You can use the clients in the `cmd` folder to test the service. 
Bear in mind, like I said above, the service takes a while to get started...

* The name starts with `dom` because I made this as part of making a 
home management system. And "Dom" is Polish for "Home". I'm not Polish, though. :portugal:

## TODO
[ ] Write unit tests

[ ] Implement CI/CD
