# Distributed Denial of Service
The act of sending many requests to a service in order to deny legit traffic.
You can use it to load test your webservers or simulate DDOS attacks for chaos testing and such.
This tool is created as an experiment into concurrency with Go.

**Warning: do not use this outside of testing your own internal systems**

As always, but repeated for good measure.
```
We trust you have received the usual lecture from the local System
Administrator. It usually boils down to these three things:

    #1) Respect the privacy of others.
    #2) Think before you type.
    #3) With great power comes great responsibility.
```

I accept no liability for your own actions. You have been warned.

## Preview
Here I setup a local server with a basic HTML page using `python -m http.server`.
After starting the server I immediately trigger the DDOS using this tool.
It issues 15k requests.
As the requests are being sent the end client browser is denied access to the site.
Once the DDOS ends the client loads the page in Firefox.

![POC or GTFO](https://github.com/nicholasericksen/poc-videos/blob/main/gifs/DDOS-Example-GIF.gif)

## Usage
Build the program locally with
```
go build ddos.go
```

Then run the program with 

```
./ddos -url=http://localhost:9876 -qty=15 -log=true
```

You can specify the quantity, logging mode for more details, and the url of the web service.

## Known Limitations
* Currently this only supports making `GET` request to the specified URL.
* Only supports HTTP based services
* Limited statistical summary in the tool, although log analysis could be performed.
* Technically not Ddos since there is no terraform component yet.
See the issues for roadmap items and future improvements.
