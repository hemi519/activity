# Activity Application

The Activity Application is a Go program that fetches random activity data from an external API and stores it in a JSON file. It provides a way to periodically fetch and save activity data, allowing users to generate random activity suggestions.

## Features

- Fetches random activity data from an external API
- Saves the fetched data to a JSON file
- Supports periodic fetching and saving of data
- Configurable API URL and output file path
- Logging of application activities

## Prerequisites

- Go 1.16 or later installed
- Docker tool be installed
- Internet connection to fetch data from the API

## Installation

1. Clone the repository:
   git clone https://github.com/hemi519/activity.git
2. With Docker 
   - sh run.sh
   - docker exec -it activity_tracker /bin/sh (to check the logs or output.json in container)
3. Without Docker(go need to be installed on the mechine)
   go run cmd/main.

## Run Testcases

- sh run_testcase.sh 

## Config Usage

- apiURL - The URL of the external API to fetch activity data from.
- output - The path to the JSON file where the fetched data will be saved.
- fetchInterval - The interval at which data should be fetched and saved. Use Go duration format (e.g., "1h" for 1 hour, "30m" for 30 minutes, "15s" for 15 seconds).

## To Be Noted

- This is not a production ready code. It has features of dockerization, unit testcase and logging. This can be improved a lot in terms of writing testcases and logging for all functionality
