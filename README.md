# gobeepme

[![Build Status](https://api.travis-ci.org/daveshanley/gobeepme.svg)](https://travis-ci.org/daveshanley/gobeepme) [![GoDoc](https://godoc.org/github.com/daveshanley/gobeepme?status.svg)](http://godoc.org/github.com/daveshanley/gobeepme)

A simple console app/library/service to allow you to quickly ping and locate your iOS device.

## Update (October 2021)

This code is pretty old, the UI code is practically fossilized. It does still work however. 

## What exactly is it though?

If you're like me, you are always putting your iPhone face down on a couch or table somewhere and forgetting where you put it.
 Especially if it's face-down on a dark surface (and you have a dark colored iPhone). It always results
in having to login to iCloud on my laptop and sending a sound to my iPhone. This is annoying - so I wanted something
simpler. I wanted to simply ask my [Amazon Echo](http://amazon.com/echo) where my phone was and simply have it
beep. That would require some kind of hosted service, so I built this!

## What does it do?

### Well it runs in 2 different ways...

* Runs as an interactive console application that you can step through.
* Runs as an http service over TLS with very simple JSON API.

The service also provides a simple UI that uses with the API.

## Building

Check out the code.

```console
git clone https://github.com/daveshanley/gobeepme.git
```

The project is now using go modules, so just type:

```console
go build gobeepme.go
```

## Running gobeepme

### Console experience

To run the console, simply run the `gobeepme` executable from a console. You will be guided from there. You can also supply
a number of flags to avoid typing them in. The flags are: 

    Usage of ./gobeepme:
      -msg string
            Message to be sent to iOS device (default "Beep Beep!")
      -name string
            Name of the iOS device you want to beep
      -user string
            Your iCloud ID / AppleID (normally an email)
      -pass string
            Pretty sure this is self explanatory
      -port int
            (service only) Port to run https service on (default 9443)
      -service
            Run as https service
      -cert string
            (service only) certificate to use
      -key string
            (service only) private server key
      
### Service experience

To run the service you will need an SSL cert/private key. If you don't have this already (most likely you don't) then you can 
generate a self signed cert using openssl by issuing the following command. 

```bash
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem
```

[Click here to read more on Stack Overflow](http://stackoverflow.com/questions/10175812/how-to-create-a-self-signed-certificate-with-openssl)

Or if you want to run gobeep me as a full stateless service in the cloud, then you will need an actual valid certificate. You can either pay
for one of these, or you can use [Let's Encrypt](https://letsencrypt.org/) for completely free and valid certs (with a short lifetime).

### Starting the service

Simply pass in the `-service` flag, your key and your cert location, and an optional port (defaults to 9443)

```bash
./gobeepme -service -cert=cert.pem -key=key.pem
```

You should then see a message stating: 

```bash
Starting beepme as a service on port 9443
```

# Connect to your Amazon Echo

Pretty simple really. The Echo supports [IFTTT](https://ifttt.com/) (If This The That). You simply need to add the IFTTT channel to your Echo and use a simple recipe 
to trigger an IFTTT Maker event when you speak a trigger word. To make this simple, I have created a *[gobeepme sample recipe](https://ifttt.com/recipes/378582-gobeepme-sample)*
it's the same one that I also use daily. 

The service request for a beep is dead simple.
```json
{"apple_id": "your_id","password":"your_passwd", "name":"device_name","message":"Beep Beep!"}
``` 
The service endpoint is `/beep`, requires the data to be a POST and the content-type needs to be `application/json`



