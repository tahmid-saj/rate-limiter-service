# rate-limiter-service

Rate limiter service to limit requests using a sliding window (with logging) approach and pre-defined rules. Developed with Go / Gin, DynamoDB.

<br/>
<br/>

## Directory structure

The directory structure is as follows:

```
rate-limiter-service/
├── dynamodb/                # Code related to setting up and interacting with DynamoDB for request logs and rules
├── models/                  # Contains data models for requests and rate-limiting rules
├── routes/                  # API route definitions using Gin framework
├── sliding-window/           # Logic for the sliding window algorithm used to limit requests
├── utils/                   # Utility functions (e.g., helpers for logging, error handling)
├── .gitignore               # Specifies files to ignore in version control
├── README.md                # Project overview, setup, and usage instructions
├── go.mod                   # Go module dependencies
├── go.sum                   # Hashes for Go module dependencies
└── main.go                  # Entry point for the rate limiter service

```

<br/>
<br/>

## Overview

### Design

The high level workflow of the rate limiter can be found below. Similar services can be found <a href="https://whimsical.com/web-microservices-6uqvwWZtcBFsNJB2hepGy1">here</a> and below:

#### Rate limiter workflow

<img width="508" alt="image" src="https://github.com/user-attachments/assets/a343b5ff-ddf1-4a5e-81c8-c26aa71e570b">

#### Similar services

<img width="834" alt="image" src="https://github.com/user-attachments/assets/b54088e7-870c-46dd-9cf6-2e5ec27d9d5c">

### Examples

#### Sample rule
```
{
  "RuleName": {
    "S": "chat-sigma-api-chat-message"
  },
  "Limit": {
    "N": "1"
  },
  "ParamName": {
    "S": "send-message"
  },
  "WindowInterval": {
    "N": "1"
  },
  "WindowTime": {
    "S": "minute"
  }
}
```

#### Sample logged requests in sliding window
```
{
  "RequestID": {
    "S": "1044eb44-c478-4caa-ad8a-0ec43a4a8495"
  },
  "LogRequests": {
    "L": [
      {
        "M": {
          "ParamName": {
            "S": "send-message"
          },
          "RuleName": {
            "S": "chat-sigma-api-chat-message"
          },
          "Timestamp": {
            "S": "2024-10-13T22:06:19.1614674-04:00"
          }
        }
      },
      {
        "M": {
          "ParamName": {
            "S": "send-message"
          },
          "RuleName": {
            "S": "chat-sigma-api-chat-message"
          },
          "Timestamp": {
            "S": "2024-10-13T22:49:03.1318595-04:00"
          }
        }
      }
    ]
  }
}
```
