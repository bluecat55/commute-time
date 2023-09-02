
# Welcome to commute-time!

A simple productivity tool that shows you the commute time

Wanted to know the time to go to/return from office at the run of a simple command? Follow the simple steps mentioned below

## Prerequisites
Need a Google maps key with directions enabled

## Steps

 1. Clone the repo
 2. Rename .env.example to .env and enter your home address, office address and google maps API key
 3. go build
 4. Configure the executable in your path
 5. Run the command specifying the mode as below. An array will be logged showing the duration for all possible routes

## Modes
Whether you want to get the depart time (to office) or the return time (from office)

 1. Pass r or return for return mode like so `commute-time r`
 2. Pass any other character for depart eg `commute-time d`
 3. Mode defaults to return when no args are given (as it is more tedious to type towards the end of the day)

## Disclaimer
Though this is a genuine use case for me, its still a light hearted project


