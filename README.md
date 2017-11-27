# planetary-gear-designer
A tool that calculates possible sets of matching compound planetary gear output parameters based on a set of input gear parameters.


## Usage
To test the program, invoke `gear-designer`. This will execute the program with the following default input:
```
17 Sun gear teeth, 20 Planet gear teeth, a minimum of 500:1 gain ratio, and a maximum of 3000:1 gain ratio
```

There are 6 input:
- `sun1-start`: Defines the starting range of the number of teeth for the Sun gear in the first gear set.
- `sun1-end`: Defines the ending range of the number of teeth for the Sun gear in the first gear set
- `planet1-start`: Defines the starting range of the number of teeth for the Planet gear in the first gear set
- `planet1-end`: Defines the ending range of the number of teeth for the Planet gear in the first gear set
- `min-gain`: Defines the minimum gear ratio desired in the output gear design
- `max-gain`: Defines the maximum gear ratio desired in the output gear design


## Development
The Dockerfile should instantiate a full dev environment for running and testing this module.

To run the container, run `make start`. This will create a new Docker container with all dependencies installed and the source code compiled, ready to use.

Before code commit, run `make test`. This will create a new Docker container with all dependencies installed and run the unit test (To Be Written).
