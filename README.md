# ONDC Buyer App in Go

This is a dummy implementation of an ONDC (Open Network for Digital Commerce) Buyer Application Platform (BAP) in Go. It implements the basic ONDC protocol APIs for a buyer application.

## Project Structure

```
.
├── cmd/
│   └── api/            # Application entry points
│       └── main.go     # Main application entry point
├── internal/           # Private application code
│   ├── api/           
│   │   ├── handlers/   # Request handlers
│   │   ├── routes/     # Route definitions
│   │   └── server/     # Server setup
│   ├── models/         # Data models
│   └── config/         # Configuration
├── pkg/                # Public packages
├── Dockerfile         # Docker build instructions
├── docker-compose.yml # Docker compose configuration
└── README.md
```

## Features

- Basic ONDC protocol implementation
- Action endpoints (search, select, init, confirm, status, cancel)
- Callback endpoints (on_search, on_select, on_init, on_confirm, on_status, on_cancel)
- Environment-based configuration
- Request validation and error handling
- Structured logging
- Docker support for easy deployment

## Prerequisites

- Go 1.21 or higher (for local development)
- Git
- Docker and Docker Compose (for containerized deployment)

## Setup

### Local Development

1. Clone the repository:
```bash
git clone <repository-url>
cd ondc-buyer-app
```

2. Set up environment variables:
```env
PORT=3000
BAP_ID=marblex-bap-preprod.marblex.ai
BAP_URI=https://evolving-fawn-cleanly.ngrok-free.app/v1
BPP_URI=https://dev-automation.ondc.org/api-service/ONDC:TRV11/2.0.0/seller
```

3. Install dependencies:
```bash
go mod download
```

4. Run the application:
```bash
go run cmd/api/main.go
```

### Docker Deployment

1. Clone the repository:
```bash
git clone <repository-url>
cd ondc-buyer-app
```

2. Build and run with Docker Compose:
```bash
docker-compose up --build
```

The application will be available at `http://localhost:3000`.

To run in detached mode:
```bash
docker-compose up -d
```

To stop the containers:
```bash
docker-compose down
```

## API Endpoints

All endpoints are prefixed with `/v1`

### Action Endpoints:
- POST `/search` - Initiate a search request
- POST `/select` - Select an item from search results
- POST `/init` - Initialize an order
- POST `/confirm` - Confirm an order
- POST `/status` - Check order status
- POST `/cancel` - Cancel an order
- POST `/update` - Update an order

### Callback Endpoints:
- POST `/on_search` - Receive search results
- POST `/on_select` - Receive selection confirmation
- POST `/on_init` - Receive initialization confirmation
- POST `/on_confirm` - Receive order confirmation
- POST `/on_status` - Receive status update
- POST `/on_cancel` - Receive cancellation confirmation
- POST `/on_update` - Receive update confirmation

## Testing

You can test the endpoints using curl:

```bash
# Test search callback
curl -X POST http://localhost:3000/v1/on_search \
  -H "Content-Type: application/json" \
  -d '{"context": {"transaction_id": "123"}, "message": {"catalog": {}}}'
```

## Development

This is a basic implementation for demonstration purposes. For production use, you should:

1. Implement proper request signing using ONDC's cryptographic requirements
2. Add proper error handling and logging
3. Implement complete request validation
4. Add database integration for storing transactions
5. Implement proper security measures
6. Add tests

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request 