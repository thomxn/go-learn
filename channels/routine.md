channels## Routines
The go scheduler runs on one CPU core a single routine until it finishes or makes a blocking call. 

By default go attempts to use only one CPU core.

If we override this, the go scheduler will assign a go routine to one thread on each logical core in a multi-core machine

**Concurrency is not parallelism**

Concurrency means that different modules can be loaded in same time(not really... but kind of same time) and switch between them on the go
Parallelism means that we can load and execute multiple things at the same time

Concurrency is more leaned on programming and parallelsm is more leaned on hardware + software

**Main routinne** is the main and default go routine

_Child routines are created using **go** keyword_

Just prefix the function call with **go()**

Eg: _go functionname()_

## Channels

Go routines communicate with each other via channels. Channels are typed ie. the type of data to be passed around a channel is predefined

Channels are two way communication mechanisms

**<-** is used to direct the data in the channel

Whenever we wait for a message in a channel **IT is a BLOCKING call**

Ising **range** keyword with channel means that we are specifying go routine to wait for the channel

Function literals are the JS, Python equivalent of lambda in go. 
The syntax is similar IIFE in JS; ie, function definition followed by parantheses

Make sure to properly scope the variables used in routines becuase due to concurrency, if teh all the routines use the same reference, integrity will be lost