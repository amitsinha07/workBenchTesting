# ONDC Postman Collections

This folder contains separate Postman collections for testing different aspects of the ONDC (Open Network for Digital Commerce) implementation.

## Collections Overview

### 1. Buyer Action APIs (`buyer-action-apis.json`)
**Purpose**: Test buyer-initiated actions  
**Endpoints Covered**:
- POST `/v1/buyer/search` - Initiate a search request
- POST `/v1/buyer/select` - Select an item from search results
- POST `/v1/buyer/init` - Initialize an order
- POST `/v1/buyer/confirm` - Confirm an order
- POST `/v1/buyer/status` - Check order status
- POST `/v1/buyer/cancel` - Cancel an order
- POST `/v1/buyer/update` - Update an order
- POST `/v1/buyer/track` - Track an order

### 2. Buyer Callback APIs (`buyer-callback-apis.json`)
**Purpose**: Test buyer callback endpoints that receive responses  
**Endpoints Covered**:
- POST `/v1/buyer/on_search` - Receive search results
- POST `/v1/buyer/on_select` - Receive selection confirmation
- POST `/v1/buyer/on_init` - Receive initialization confirmation
- POST `/v1/buyer/on_confirm` - Receive order confirmation
- POST `/v1/buyer/on_status` - Receive status update
- POST `/v1/buyer/on_cancel` - Receive cancellation confirmation
- POST `/v1/buyer/on_update` - Receive update confirmation
- POST `/v1/buyer/on_track` - Receive tracking updates

### 3. Seller Action APIs (`seller-action-apis.json`)
**Purpose**: Test seller endpoints that process buyer requests  
**Endpoints Covered**:
- POST `/v1/seller/search` - Process search requests from buyers
- POST `/v1/seller/select` - Process select requests from buyers
- POST `/v1/seller/init` - Process initialization requests from buyers
- POST `/v1/seller/confirm` - Process confirmation requests from buyers
- POST `/v1/seller/status` - Process status requests from buyers
- POST `/v1/seller/cancel` - Process cancellation requests from buyers
- POST `/v1/seller/update` - Process update requests from buyers
- POST `/v1/seller/track` - Process tracking requests from buyers

### 4. Seller Callback APIs (`seller-callback-apis.json`)
**Purpose**: Test seller callback endpoints that forward responses to buyer  
**Endpoints Covered**:
- POST `/v1/seller/on_search` - Forward search results to buyer
- POST `/v1/seller/on_select` - Forward selection confirmation to buyer
- POST `/v1/seller/on_init` - Forward initialization confirmation to buyer
- POST `/v1/seller/on_confirm` - Forward order confirmation to buyer
- POST `/v1/seller/on_status` - Forward status update to buyer
- POST `/v1/seller/on_cancel` - Forward cancellation confirmation to buyer
- POST `/v1/seller/on_update` - Forward update confirmation to buyer
- POST `/v1/seller/on_track` - Forward tracking updates to buyer

## How to Use

1. **Import Collections**: Import any of the JSON files into Postman
2. **Set Environment Variables**: Configure the `base_url` variable in Postman:
   - For local testing: `http://localhost:3000`
   - For deployed instances: Your actual server URL
3. **Test Workflow**: 
   - Start with buyer action APIs to initiate requests
   - Use seller action APIs to simulate processing
   - Test callback APIs to verify response handling

## Sample Test Flow

1. **Search Flow**:
   ```
   Buyer Search → Seller Search Processing → Seller on_search → Buyer on_search
   ```

2. **Complete Order Flow**:
   ```
   Buyer Search → Buyer Select → Buyer Init → Buyer Confirm → Buyer Status
   ```

3. **Seller Response Flow**:
   ```
   Seller on_search → Seller on_select → Seller on_init → Seller on_confirm → Seller on_status
   ```

## Environment Variables

All collections use the following environment variable:
- `base_url`: The base URL of your ONDC application (default: `http://localhost:3000`)

## Sample Data

All requests include sample data for mobility services (taxi booking) as per ONDC transportation domain specifications. You can modify the request bodies according to your specific use case.

## Notes

- All endpoints follow ONDC protocol standards
- Request/response formats comply with ONDC v1.1.0 specifications
- Sample data includes proper context, message structure, and domain-specific fields
- Each collection can be used independently for testing specific functionality 