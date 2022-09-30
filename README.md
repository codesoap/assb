Adam Stern's static breath hold table implemented in a
simple command line tool. He mentioned this table in [a
video](https://www.youtube.com/watch?v=eXBZUee4EpY&t=224).

Start holding your breath and start `assb`. A timer will start running.
When you want to breathe again, press return. It will be displayed how
long you held your breath and the timer will now show how long you have
left until the next breath hold starts. This continues until eight
breath holds have been completed.

Adam Stern recommends **not** to do this more than three times a week.

# Examples
A running session could look like this:

```console
$ assb
#1: Held breath for 1:47 min, breathed for 2:13 min.
#2: Held breath for 1:45 min, breathed for 2:15 min.
#3: Held breath for 1:52 min, breathed for 2:08 min.
#4: Held breath for 1:55 min, breathed for 2:05 min.
#5: Holding breath since 1:33 min... press return to breathe.
```
