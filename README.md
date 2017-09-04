# NOAA Active Weather Alert Utility

## this branch is nonsense for my weather website, ignore

run thing, pass CAP alert code, returns type, headline, issued/updated times, expiry, and description for area.

I don't know if unmarshal-ing Alerts into an array then looping through it is necessary -- I'm not sure if the NOAA spec does multiple alerts per event, if so I might get duplicates and I'll fix it later. Too lazy to read through noaa's verbose documentation and find out. I'd rather get dups than miss alerts anyway.

"DONT USE A MAKEFILE FOR GO" 

...but i'm stubborn about where i want things to be existing
...yes i know symlinks exist
...go away
