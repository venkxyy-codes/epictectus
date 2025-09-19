# Epictectus Backend - Product & Tech Flow

## Overview

Epictectus is a Sri Yatra backend system for payment processing, user management, and 
CRM integration. This README provides a product-tech flow for the main APIs, their 
responsibilities, and how data moves through the system. New developers can use this 
guide to quickly understand and extend the platform.

---

## Main Product Flows & APIs

### 1. Payment Gateway Flow

**Endpoint:**  
`POST /payment-gateway/standard-link`  
**Handler:** `PgHandler.CreateStandardPaymentLink`

**Purpose:**  
Creates a payment link (currently via Razorpay) and optionally syncs payment info with CRM (Leadsquared).

**Flow:**
- Receives payment link creation request with customer and payment details.
- Validates request and payment provider.
- Calls `PaymentGatewayService.CreateStandardPaymentLinkRazorpay`.
- On success, returns payment link and triggers CRM sync.

**Tech Stack:**  
- Handler: `handler/payment-gateway.go`  
- Service: `service/payment-gateway/payment-gateway.go`  
- Domain/Contract: `domain/payment-gateway.go`, `contract/payment-gateway.go`  
- CRM Integration: `service/crm/crm.go`

---

### 2. User Management Flow

**Endpoints:**  
- `POST /user/signup`  
- `POST /user/login`  
**Handler:** `UserHandler.SignUpUser`, `UserHandler.LoginUser`

**Purpose:**  
Handles user registration and authentication.

**Flow:**
- Receives user signup/login request.
- Validates input.
- Calls `UserService.CreateUser` or `UserService.LoginUser`.
- Returns success or error response.

**Tech Stack:**  
- Handler: `handler/user.go`  
- Service: `service/user/user.go`  
- Domain/Contract: `domain/user.go`, `contract/user.go`  
- Repo: `repo/user.go`

---

### 3. Webhook Processing Flow

**Endpoint:**  
`POST /webhook/leadsquared-activity`  
**Handler:** `WebhookHandler.ProcessLeadsquaredActivityWebhook`

**Purpose:**  
Processes incoming CRM (Leadsquared) activity webhooks.

**Flow:**
- Receives webhook payload.
- Validates and parses request.
- Calls `WebhookProcessorService.HandleLeadsquaredWebhook`.
- Returns success response.

**Tech Stack:**  
- Handler: `handler/webhook.go`  
- Service: `service/webhook_processor/webhook_processor.go`  
- Contract: `contract/webhooks.go`

---

## Architecture & Directory Structure

- **handler/**: API entry points (Gin HTTP handlers)
- **service/**: Business logic and integrations
- **repo/**: Data persistence and retrieval
- **domain/**: Core business models
- **contract/**: API request/response models
- **view/**: Response formatting
- **config/**: Configuration management
- **clients/**: External service clients (e.g., HTTP)
- **utils/**: Utility functions
- **error/**: Error definitions

---

## Setup & Integration

1. **Configuration:**  
   - Set up payment gateway and CRM credentials in `env`.
   - Example config: `.env`

2. **Run the Server:**  
   - Use `run_local.sh` or `run.sh` to start the backend.

3. **API Usage:**  
   - Use the documented endpoints for payment, user, and webhook flows.
   - Refer to contract files for request/response formats.

---

## Extending the System

- Add new payment providers by extending `PaymentGatewayService` and updating the handler.
- Integrate new CRM systems via the CRM service and contract.
- Add new user features in the user service and handler.

---

## Error Handling

- All APIs validate input and return structured error responses.
- Errors are defined in `error/error.go` and rendered via `utils`.

---

## Contact & Contribution

For questions or contributions, please refer to the codebase structure above and follow
best practices for modular development.

