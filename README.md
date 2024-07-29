Pack Calculation Application
Overview
This Golang application calculates the optimal number of packs required to fulfill an order, given a set of pack sizes. The application adheres to the following constraints:

Only whole packs can be shipped.
The minimum number of items should be shipped.
The minimum number of packs should be used.
The application includes a basic HTTP API and a simple UI for interacting with the calculation logic. Pack sizes are configurable for flexibility.

Dependencies

Go
Gin framework
(Optional) HTML template engine
Installation

Clone the repository.
Install dependencies using go mod tidy.
Configuration
Pack sizes are defined in a JSON configuration file (config.json by default). The file should contain a JSON array of pack sizes:
