## Repository: [TrimmedMeanUseExample](https://github.com/BellowAverage/TrimmedMeanUseExample)

### README.md

```markdown
# TrimmedMeanUseExample

An example Go application demonstrating how to use the `TrimmedMeanPackage` to compute trimmed means for datasets of integers and floats.

## Overview

This application:

- Generates random samples of at least 100 integers and 100 floats.
- Computes the symmetric trimmed mean using the `TrimmedMeanPackage` with a 5% trim from both ends.
- Outputs the results to the console.

## Prerequisites

- Go installed on your machine (version 1.13 or higher).
- Git installed for cloning repositories.

## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/BellowAverage/TrimmedMeanUseExample.git
   cd TrimmedMeanUseExample
   ```

2. **Get the `TrimmedMeanPackage`**

   ```bash
   go get github.com/BellowAverage/TrimmedMeanPackage
   ```

## Usage

Run the application using the following command:

```bash
go run main.go
```

This will:

- Generate random datasets.
- Compute the trimmed means.
- Display the results in the console.

## Sample Output

```
Trimmed mean of integers: 499.85
Trimmed mean of floats: 501.23
```

*Note: The actual output will vary due to the use of random data.*

## Code Explanation

### main.go

```go
package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"

    "github.com/BellowAverage/TrimmedMeanPackage/trimmedmean"
)

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Generate a sample of 100 random integers between 0 and 1000
    intData := make([]int, 100)
    for i := range intData {
        intData[i] = rand.Intn(1000)
    }

    // Compute the symmetric trimmed mean for integers with 5% trimming
    meanInt, err := trimmedmean.TrimmedMeanInt(intData, 0.05)
    if err != nil {
        log.Fatalf("Error computing trimmed mean for integers: %v", err)
    }
    fmt.Printf("Trimmed mean of integers: %.2f\n", meanInt)

    // Generate a sample of 100 random floats between 0 and 1000
    floatData := make([]float64, 100)
    for i := range floatData {
        floatData[i] = rand.Float64() * 1000
    }

    // Compute the symmetric trimmed mean for floats with 5% trimming
    meanFloat, err := trimmedmean.TrimmedMeanFloat(floatData, 0.05)
    if err != nil {
        log.Fatalf("Error computing trimmed mean for floats: %v", err)
    }
    fmt.Printf("Trimmed mean of floats: %.2f\n", meanFloat)
}
```

### Explanation

- **Random Data Generation**

  - Uses `math/rand` to generate random integers and floats.
  - Seeds the random number generator with the current time to ensure different results on each run.

- **Computing Trimmed Means**

  - Calls `TrimmedMeanInt` and `TrimmedMeanFloat` with a trimming proportion of `0.05` for symmetric trimming.
  - Handles any errors returned by the functions.

- **Output**

  - Prints the computed trimmed means to the console with two decimal places.

## Understanding Trimmed Mean

A trimmed mean is calculated after discarding a certain percentage of the smallest and largest values from a dataset. This method reduces the impact of outliers on the mean, providing a more robust measure of central tendency.

In this example, a 5% trim removes the lowest 5% and the highest 5% of the data before calculating the mean.

## Customizing the Trimming Proportion

You can adjust the trimming proportion by changing the value passed to the functions. For example, to use a 10% trim:

```go
meanInt, err := trimmedmean.TrimmedMeanInt(intData, 0.10)
meanFloat, err := trimmedmean.TrimmedMeanFloat(floatData, 0.10)
```

## Error Handling

Ensure to handle errors returned by the trimmed mean functions to prevent unexpected crashes.