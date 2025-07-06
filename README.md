# jiggler
![KrankyBearBeret](https://github.com/user-attachments/assets/95aef02f-a72c-4b82-aad2-5d2e4b30315a)

This is a simple mouse jiggler.

It requires at least one parameter -jiggle or -j flag usage followed by a time
in minutes from 1 to 60, default is 10 if a number less than 1 or greater
than 60 are provided

A second parameter -pixels or -p allows user choice of the number of pixels to
move the mouse. The default is 1, if less than 1 or greater than 100 are entered.
A pixel move one right, one down, one left one up will reposition the mouse faster
than larger numbers and be less invasive with normal use.

Example use: (single - or double dash both work)
* ./jiggler -jiggle 15 -pixel 1 &
* ./jiggler -j 10 -p 1 &
* ./jiggler --j 10 --p 5 &
* ./jiggler -help | --help | -h | --h

Daemonization / backgrounding is not provided, use your operating system to pass
the parameters you want and cause this application to run as a background process.



# To-do / known problems
- No problems I am aware of
- One possible addition, one extra flag to add to allow for command line checking for
updates. e.g. jiggler -checkupdate | jiggler -cu


# License
This is 100% free for anyone to use or misuse any way you like with no warranty as
to suitability or anything else, other than it has no viruses when I compile and
commit to git. But you should always check and scan anything you download from the
internet for viruses anyway. Don't be reckless.

Do NOT use jiggler when it will violate your security policies in public spaces.
This is intended for use only in secure, remote, work from home type environments.

In other words, I take no responsibility for how you use this, protect yourself. 
To be more blunt, **don't be stupid!**
