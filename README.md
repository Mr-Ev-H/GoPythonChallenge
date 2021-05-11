# GoPyChallenge 

## Summary

Solutions to the [Python Challenge](https://pythonchallenge.org) in Go.

The Python Challenge is an online coding challenge that went viral for a brief period in the late 2000's. Structured into a set of numbered puzzles, each challene could be solved using the Python standard lbrary or commonly available packages. This is my attempt to produce solutions to the Python Challenge using Go. 

## Running Challenges

The application requires you pass a challenge number as its only argument.

    gopychallenge 15

* Where its been possible to generate the the answer directly as ascii, this will be joined to the url and printed to stdout.

      Next: http://www.pythonchallenge.com/pc/def/274877906944.html

* Where a challenge required following a series of steps, the steps will be written to stdout.

* Where the answer is saved to one or more files, the path to these will be printed to stdout. Answer files will be written to a local directory (./answers/)
e.g.

      Answer saved to: answers/chal16.png

* If the challenge ends in a clue, the challenge will output the final step in the challenge you are expected to use in order solve the problem. (17 outputs a html document containing the answer, 15 outputs somebodys date of birth).

Only challenge 3 has an alternate solution, which can be run with 3a. Unsurprisingly it produces the same result as 3, only slightly faster.

## Why Does This Exist?

A long time ago I taught myself python by solving as much of the python challenge as possible. Learning Go has been on my list for a while and I quite liked the symmetry learning Go the same way I learnt Python. Both languages have a robust standard library, both require little overhead to get going, and both let you write terse code with fast feedback loops. 

## Project Structure

The project is launched from the challenge.go file.

All challenges solved have a corresponding .go file in the project's challenges directory.

Some of the challenges require content hidden within from the challenge page, rather than parsing the html of the pages, I've included the text either as constants in the go file, or as local assets read at the start of the challenge. These can be found in the in the challenges/data directory.

A small amount of common functionality has been abstracted to the utility module.

## Future Work

I probably won't be putting much work into this project beyond solving as many challenges as I can. Once I've arrived at a solution I'll tidy up the code as much as possible and leave it be. Challenges that run without touching the pythonchallenge site all run in less than 1 second, so I'm happy that they're doing what they need to do.

I've not yet explored unit testing within Go yet so I may look to add testing support at some point which may require some refactoring.  

I will be writing my experience learning Go and general thougths on the Python Challenge at some point in the near future.

## Known issues:
 * Challenge 13 (XML RPC) will print an warning message to stdout about the format of the response sent by the PythonChallenge API. it doesnt impact solving the challenge and can be ignored. 

