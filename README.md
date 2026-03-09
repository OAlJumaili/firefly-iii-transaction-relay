# firefly-iii-transaction-relay
I created this Go API Endpoint to accept transaction SMS messages from my bank, parse and extract the information in them to form a Firefly III transaction JSON payload and post it to the API and avoid manually inputting transactions into Firefly.

I say SMS messages, but it can accept literally any form of text message. you can customize what "key phrases" are used to extract the details of the transaction via environment variables, `.env.example` provides an empty template of all the necessary environment variables.

## Installation
You can either run this via Docker, or as an executable.

### Docker
```
docker run -p 8080:8080 judexgrim/firefly-iii-transaction-relay:latest
```

### Executable
Download the executable suitable for your machine's architecture from the latest release, and run it.