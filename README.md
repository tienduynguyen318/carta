# Carta Coding Challenge

## Installation and running this solution

1. Install golang following this [link](https://golang.org/doc/install)
2. Run command build
3. Run the binary file with the arguments

## Assumption

1. There is no logging
2. The quantity column fit inside float64
3. The date column does not contain timezone
4. One row of bad data will stop the application (number of column different thang 6, cancel share larger than vested share)

## Improvement

1. Add logging
2. Handle error better: Instead of panic cancel option larger than vested option, exit the application
3. Tolerate faulty data: If the data is incorrect/malformed, simply continue to the next row
4. Persist the report: It is likely that the report will be reused and therefore we should persist it
