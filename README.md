# ONDC Buyer & Seller App in Go

This is a dummy implementation of an ONDC (Open Network for Digital Commerce) application that supports both Buyer Application Platform (BAP) and Seller Application Platform (SAP) functionality in Go. It implements the basic ONDC protocol APIs for both buyer and seller applications.

## Table of Contents

- [Project Structure](#project-structure)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
  - [Postman Collections](#postman-collections)
  - [cURL Examples](#curl-examples)
- [Development](#development)
- [Contributing](#contributing)

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
├── postman-collections/ # Postman API testing collections
│   ├── README.md       # Detailed testing documentation
│   ├── buyer-action-apis.json    # Buyer action endpoints
│   ├── buyer-callback-apis.json  # Buyer callback endpoints
│   ├── seller-action-apis.json   # Seller action endpoints
│   └── seller-callback-apis.json # Seller callback endpoints
├── Dockerfile         # Docker build instructions
├── docker-compose.yml # Docker compose configuration
└── README.md
```

## Features

- Complete ONDC protocol implementation for both buyer and seller platforms
- **Buyer Application Platform (BAP) features:**
  - Action endpoints (search, select, init, confirm, status, cancel, update, track)
  - Callback endpoints (on_search, on_select, on_init, on_confirm, on_status, on_cancel, on_update, on_track)
- **Seller Application Platform (SAP) features:**
  - Request processing endpoints (search, select, init, confirm, status, cancel, update, track)
  - Callback forwarding endpoints (on_search, on_select, on_init, on_confirm, on_status, on_cancel, on_update, on_track)
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
# Server Configuration
PORT=3000
BAP_ID=marblex-bap-preprod.marblex.ai
BAP_URI=https://evolving-fawn-cleanly.ngrok-free.app/v1
BPP_URI=https://dev-automation.ondc.org/api-service/ONDC:TRV11/2.0.0/seller

# Authentication Configuration (ONDC Signature)
PRIVATE_KEY=your_base64_encoded_private_key_here
PUBLIC_KEY=your_base64_encoded_public_key_here
SUBSCRIBER_ID=buyer-app.ondc.org
UNIQUE_KEY_ID=207
```

**Generate ED25519 keys for authentication:**
```bash
# Generate private key
openssl genpkey -algorithm Ed25519 -out private_key.pem

# Generate public key  
openssl pkey -in private_key.pem -pubout -out public_key.pem

# Convert to base64 for environment variables
echo "PRIVATE_KEY=$(openssl pkey -in private_key.pem -raw -outform DER | base64)"
echo "PUBLIC_KEY=$(openssl pkey -pubin -in public_key.pem -raw -outform DER | base64)"
```

**Environment Variables Notes:**
- `BAP_ID` and `BAP_URI` are used for buyer application functionality
- `BPP_URI` is used for seller callback forwarding
- `PRIVATE_KEY` and `PUBLIC_KEY` are required for ONDC signature authentication
- `SUBSCRIBER_ID` and `UNIQUE_KEY_ID` identify your application in the ONDC network

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

### Buyer Endpoints

#### Action Endpoints:
- POST `/buyer/search` - Initiate a search request
- POST `/buyer/select` - Select an item from search results
- POST `/buyer/init` - Initialize an order
- POST `/buyer/confirm` - Confirm an order
- POST `/buyer/status` - Check order status
- POST `/buyer/cancel` - Cancel an order
- POST `/buyer/update` - Update an order
- POST `/buyer/track` - Track an order

#### Callback Endpoints:
- POST `/buyer/on_search` - Receive search results
- POST `/buyer/on_select` - Receive selection confirmation
- POST `/buyer/on_init` - Receive initialization confirmation
- POST `/buyer/on_confirm` - Receive order confirmation
- POST `/buyer/on_status` - Receive status update
- POST `/buyer/on_cancel` - Receive cancellation confirmation
- POST `/buyer/on_update` - Receive update confirmation
- POST `/buyer/on_track` - Receive tracking updates

### Seller Endpoints

#### Action Endpoints:
- POST `/seller/search` - Process search requests from buyers
- POST `/seller/select` - Process select requests from buyers
- POST `/seller/init` - Process initialization requests from buyers
- POST `/seller/confirm` - Process confirmation requests from buyers
- POST `/seller/status` - Process status requests from buyers
- POST `/seller/cancel` - Process cancellation requests from buyers
- POST `/seller/update` - Process update requests from buyers
- POST `/seller/track` - Process tracking requests from buyers

#### Callback Endpoints:
- POST `/seller/on_search` - Forward search results to buyer
- POST `/seller/on_select` - Forward selection confirmation to buyer
- POST `/seller/on_init` - Forward initialization confirmation to buyer
- POST `/seller/on_confirm` - Forward order confirmation to buyer
- POST `/seller/on_status` - Forward status update to buyer
- POST `/seller/on_cancel` - Forward cancellation confirmation to buyer
- POST `/seller/on_update` - Forward update confirmation to buyer
- POST `/seller/on_track` - Forward tracking updates to buyer

## Testing

### Postman Collections

We provide comprehensive Postman collections for testing all endpoints. Import them from the `postman-collections/` folder:

- **`buyer-action-apis.json`** - Buyer action endpoints (search, select, init, confirm, etc.)
- **`buyer-callback-apis.json`** - Buyer callback endpoints (on_search, on_select, etc.)
- **`seller-action-apis.json`** - Seller action endpoints for processing buyer requests
- **`seller-callback-apis.json`** - Seller callback endpoints for forwarding responses

📖 **For detailed usage instructions, sample workflows, and environment setup, see:**  
**[Postman Collections Documentation](./postman-collections/README.md)**

### cURL Examples

You can also test the endpoints using curl:

#### Buyer Endpoints

```bash
# Test buyer search callback
curl -X POST http://localhost:3000/v1/buyer/on_search \
  -H "Content-Type: application/json" \
  -d '{"context": {"transaction_id": "123"}, "message": {"catalog": {}}}'

# Test buyer search action
curl -X POST http://localhost:3000/v1/buyer/search \
  -H "Content-Type: application/json" \
  -d '{"context": {"transaction_id": "123"}, "message": {"intent": {}}}'
```

#### Seller Endpoints

```bash
# Test seller search processing
curl -X POST http://localhost:3000/v1/seller/search \
  -H "Content-Type: application/json" \
  -d '{"context": {"transaction_id": "123"}, "message": {"intent": {}}}'

# Test seller callback forwarding
curl -X POST http://localhost:3000/v1/seller/on_search \
  -H "Content-Type: application/json" \
  -d '{"context": {"transaction_id": "123"}, "message": {"catalog": {}}}'
```

## Development

### Authentication Features

This implementation includes **automatic ONDC authentication**:

✅ **Automatic header generation** - Auth headers are automatically added to all outgoing buyer action APIs  
✅ **Optional signature verification** - Incoming callback requests are verified if auth headers are present  
✅ **ONDC v1.1.0 compliant** - Uses ED25519 signatures with BLAKE-512 hashing  
✅ **Environment-based configuration** - Easy setup with environment variables  

For detailed authentication documentation, see: **[Authentication Documentation](./docs/authentication.md)**

### Production Considerations

For production deployment, consider implementing:

1. ✅ ~~Proper request signing using ONDC's cryptographic requirements~~ (Already implemented)
2. Enhanced error handling and monitoring
3. Complete request validation and schema enforcement
4. Database integration for storing transactions and audit logs
5. Rate limiting and security hardening
6. Comprehensive test coverage
7. Key rotation and secrets management
8. Health checks and observability

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request 