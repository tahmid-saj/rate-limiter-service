# rate-limiter-service

Rate limiter service to limit requests using a sliding window (with logging) approach and pre-defined rules. Developed with Go / Gin, DynamoDB.

<br/>
<br/>

Sample rule:
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

<br/>
<br/>

Sample logged requests in sliding window:
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
