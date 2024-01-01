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

![](https://raw.githubusercontent.com/nicholasericksen/poc-videos/raw/d4e42fd9b4e1164ad77d6a0795bc8376ae262ac4/videos/DDOS_Example.mov)

https://raw.githubusercontent.com/nicholasericksen/poc-videos/raw/d4e42fd9b4e1164ad77d6a0795bc8376ae262ac4/videos/DDOS_Example.mov

<video width="320" height="240" controls>
  <source src="https://raw.githubusercontent.com/nicholasericksen/poc-videos/raw/d4e42fd9b4e1164ad77d6a0795bc8376ae262ac4/videos/DDOS_Example.mov" type="video/mp4"
  </video>

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

See the issues for roadmap items and future improvements.
