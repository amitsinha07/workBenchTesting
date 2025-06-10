# ONDC Buyer & Seller App in Go

This is a dummy implementation of an ONDC (Open Network for Digital Commerce) application that supports both Buyer Application Platform (BAP) and Seller Application Platform (SAP) functionality in Go. It implements the basic ONDC protocol APIs for both buyer and seller applications.

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
PORT=3000
BAP_ID=marblex-bap-preprod.marblex.ai
BAP_URI=https://evolving-fawn-cleanly.ngrok-free.app/v1
BPP_URI=https://dev-automation.ondc.org/api-service/ONDC:TRV11/2.0.0/seller
```

Note: 
- `BAP_ID` and `BAP_URI` are used for buyer application functionality
- `BPP_URI` is used for seller callback forwarding

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

See `postman-collections/README.md` for detailed usage instructions.

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