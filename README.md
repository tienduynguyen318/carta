# Carta Coding Challenge

## Installation and running this solution

1. Install golang following this [link](https://golang.org/doc/install)
2. Run command build
3. Run the binary file with the arguments

## Assumption

The biggest assumption was to solely focus the main functionality implementation and not to focus on error handling. In a real world scenario, it would be best practice to not crash the application when encounter bad data

I also choose to approach the code by having a small domain layer. I find that I can drastically simplify the core logic and testing by introducing this small layer. Effectively the input reader and output writer forms an adapter to the main use-cases. An example of this is implemented in the run report actions. This also allows other input reader and output writer to execute the domain logic in the future from other source rather than only read from csv and output to stdout

Also, no logging, metrics or observability was implemented as part of the exercise

## Approach

### 1. Implement a run report use-case

First, I create the tests that I need to run by reading the csv file. Then, I implement the code, read the whole csv file, looping through the result and write to the stdout using fmt.Println(). That way at least I have a runable code and rerunable set of tests to make sure the code I written is correct. The test is evolved as I add more code.

### 2. Put in the iterator pattern

Second, I refactor the reading of the csv input into iterator pattern. That way, I don't have to store the whole csv file in memory

### 3. Refactor to light domain model

Third, I refactor the code into the domain model where all the logic will happen. The domain model will take in inputReader and outputWriter interface as argument. Using abstraction make it easy to switch out inputReader and outputWriter, as long as the interface stays the same. Abstraction also make it easy to switchout inputReader and outputReader to a stub for the purpose of testing

### 4. Refactor the test into factory pattern

Lastly, I remove the dependency of the test from csv file by swapping it with factory pattern which create the object in memory

## Future Improvement

1. Add logging
2. Handle error better: Instead of panic cancel option larger than vested option, exit the application
3. Tolerate faulty data: If the data is incorrect/malformed, simply continue to the next row
4. Persist the report: It is likely that the report will be reused and therefore we should persist it
