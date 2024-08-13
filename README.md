# Pack Calculation Application
## Overview

This Golang application calculates the optimal number of packs required to fulfill an order, given a set of pack sizes. The application adheres to the following constraints:

Only whole packs can be shipped.
The minimum number of items should be shipped.
The minimum number of packs should be used.
The application includes a basic HTTP API and can include a simple UI for interacting with the calculation logic. Pack sizes are configurable for flexibility.

## Dependencies

Go
Gin framework
(Optional) HTML template engine
Installation

Clone the repository.
Install dependencies using go mod tidy.

## Configuration
Pack sizes are defined in a JSON configuration file (config.json by default). The file should contain a JSON array of pack sizes:

```JSON
[
  { "size": 5000 },
  { "size": 2000 },
  { "size": 1000 },
  { "size": 500 },
  { "size": 250 }
]
```

## Usage

### API:
Send a GET request to /packs?quantity=<order_quantity> to get the pack breakdown.
Example: curl http://localhost:8080/packs?quantity=1250
### UI:
Access the UI at http://localhost:8080
Enter the order quantity and submit the form.
### Testing
Unit tests are included in the test package. Run tests using go test -v.
