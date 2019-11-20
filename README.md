# Face Registry [![CircleCI](https://circleci.com/gh/diogox/dom-face-registry.svg?style=svg)](https://circleci.com/gh/diogox/dom-face-registry)

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

## Testing
There are clients available for manual testing in the `cmd/clients` directory. 
You can run them like so:

```shell script
go run ./cmd/clients/addPerson.go
```

This particular client will give you a `person id` you will need to pass as an 
argument when adding a face:

```shell script
go run ./cmd/clients/addFace.go 77d71383-31c9-4339-9a61-feed335bba9f
```

Here, `77d71383-31c9-4339-9a61-feed335bba9f` is the `person id` we got from the previous 
command.

Other clients, like `removePerson` and `removeFace` also need IDs as arguments.

To recognize a face, you'd use the `recognize` client. 

You can use your own images by replacing the ones in the `data` directory.
Replace `addedFace.jpg` and `recognizable.jpg` with your own pictures and 
check the results!

If you find that the recognizer is not working the way it should, you 
can make it more, or less, "sensitive" by meddling with the threshold
in `config.yaml`. A lower threshold makes false positives more likely.
