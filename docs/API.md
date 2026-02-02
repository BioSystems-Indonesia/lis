# LIS API Documentation

Base URL: `http://localhost:8080`

## Table of Contents

- [Authentication](#authentication)
- [Patients API](#patients-api)
- [Work Orders API](#work-orders-api)
- [Response Format](#response-format)
- [Error Codes](#error-codes)

---

## Authentication

Currently, the API does not require authentication. This should be implemented in production.

---

## Patients API

### Create Patient

Create a new patient record.

**Endpoint:** `POST /patients`

**Request Body:**

```json
{
  "first_name": "John",
  "last_name": "Doe",
  "birth_date": "1990-05-15T00:00:00Z",
  "sex": "male",
  "address": "Jl. Contoh No. 123, Jakarta",
  "phone": "081234567890",
  "email": "john.doe@example.com"
}
```

**Request Fields:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| first_name | string | Yes | Patient's first name |
| last_name | string | Yes | Patient's last name |
| birth_date | string (ISO 8601) | Yes | Patient's birth date |
| sex | string | Yes | Patient's gender: "male" or "female" |
| address | string | No | Patient's address |
| phone | string | No | Patient's phone number |
| email | string | No | Patient's email address |

**Success Response (201 Created):**

```json
{
  "code": 201,
  "status": "success",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "first_name": "John",
    "last_name": "Doe",
    "birth_date": "1990-05-15T00:00:00Z",
    "sex": "male",
    "address": "Jl. Contoh No. 123, Jakarta",
    "phone": "081234567890",
    "email": "john.doe@example.com"
  }
}
```

**cURL Example:**

```bash
curl -X POST http://localhost:8080/patients \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "birth_date": "1990-05-15T00:00:00Z",
    "sex": "male",
    "address": "Jl. Contoh No. 123, Jakarta",
    "phone": "081234567890",
    "email": "john.doe@example.com"
  }'
```

---

### Get Patient by ID

Retrieve a specific patient by ID.

**Endpoint:** `GET /patients?id={patient_id}`

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | string | Yes | Patient UUID |

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "first_name": "John",
    "last_name": "Doe",
    "birth_date": "1990-05-15T00:00:00Z",
    "sex": "male",
    "address": "Jl. Contoh No. 123, Jakarta",
    "phone": "081234567890",
    "email": "john.doe@example.com"
  }
}
```

**cURL Example:**

```bash
curl -X GET "http://localhost:8080/patients?id=550e8400-e29b-41d4-a716-446655440000"
```

---

### Update Patient

Update an existing patient record.

**Endpoint:** `PUT /patients?id={patient_id}`

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | string | Yes | Patient UUID |

**Request Body:**

```json
{
  "first_name": "John",
  "last_name": "Doe",
  "birth_date": "1990-05-15T00:00:00Z",
  "sex": "male",
  "address": "Jl. Updated Address No. 456, Jakarta",
  "phone": "081298765432",
  "email": "john.updated@example.com"
}
```

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "first_name": "John",
    "last_name": "Doe",
    "birth_date": "1990-05-15T00:00:00Z",
    "sex": "male",
    "address": "Jl. Updated Address No. 456, Jakarta",
    "phone": "081298765432",
    "email": "john.updated@example.com"
  }
}
```

**cURL Example:**

```bash
curl -X PUT "http://localhost:8080/patients?id=550e8400-e29b-41d4-a716-446655440000" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "birth_date": "1990-05-15T00:00:00Z",
    "sex": "male",
    "address": "Jl. Updated Address No. 456, Jakarta",
    "phone": "081298765432",
    "email": "john.updated@example.com"
  }'
```

---

### Delete Patient

Delete a patient record.

**Endpoint:** `DELETE /patients?id={patient_id}`

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | string | Yes | Patient UUID |

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": {
    "message": "Patient deleted successfully"
  }
}
```

**cURL Example:**

```bash
curl -X DELETE "http://localhost:8080/patients?id=550e8400-e29b-41d4-a716-446655440000"
```

---

### Get All Patients

Retrieve all patient records.

**Endpoint:** `GET /patients`

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "first_name": "John",
      "last_name": "Doe",
      "birth_date": "1990-05-15T00:00:00Z",
      "sex": "male",
      "address": "Jl. Contoh No. 123, Jakarta",
      "phone": "081234567890",
      "email": "john.doe@example.com"
    },
    {
      "id": "660e8400-e29b-41d4-a716-446655440001",
      "first_name": "Jane",
      "last_name": "Smith",
      "birth_date": "1985-03-20T00:00:00Z",
      "sex": "female",
      "address": "Jl. Sample No. 456, Jakarta",
      "phone": "081298765432",
      "email": "jane.smith@example.com"
    }
  ]
}
```

**cURL Example:**

```bash
curl -X GET http://localhost:8080/patients
```

---

### Search Patients

Search patients by name, phone, or email.

**Endpoint:** `GET /patients?q={search_query}`

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| q | string | Yes | Search query (searches in first_name, last_name, phone, email) |

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "first_name": "John",
      "last_name": "Doe",
      "birth_date": "1990-05-15T00:00:00Z",
      "sex": "male",
      "address": "Jl. Contoh No. 123, Jakarta",
      "phone": "081234567890",
      "email": "john.doe@example.com"
    }
  ]
}
```

**cURL Example:**

```bash
curl -X GET "http://localhost:8080/patients?q=John"
```

---

## Work Orders API

### Create Work Order

Create a new work order with patient information. This will automatically create a patient record if provided.

**Endpoint:** `POST /work-orders`

**Request Body:**

```json
{
  "no_order": "WO001",
  "test_code": ["HB", "LEUKOSIT", "ERITROSIT"],
  "patient": {
    "first_name": "Jane",
    "last_name": "Smith",
    "birth_date": "1985-03-20T00:00:00Z",
    "sex": "female",
    "address": "Jl. Sample No. 456, Jakarta",
    "phone": "081298765432",
    "email": "jane.smith@example.com"
  },
  "analyst": "Dr. Analyst",
  "doctor": "Dr. Smith"
}
```

**Request Fields:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| no_order | string | Yes | Work order number (unique) |
| test_code | array[string] | Yes | List of test codes |
| patient | object | Yes | Patient information (see Patient fields above) |
| analyst | string | Yes | Analyst name |
| doctor | string | Yes | Doctor name |

**Success Response (201 Created):**

```json
{
  "code": 201,
  "status": "success",
  "data": {
    "no_order": "WO001",
    "test_code": ["HB", "LEUKOSIT", "ERITROSIT"],
    "patient": {
      "id": "770e8400-e29b-41d4-a716-446655440002",
      "first_name": "Jane",
      "last_name": "Smith",
      "birth_date": "1985-03-20T00:00:00Z",
      "sex": "female",
      "address": "Jl. Sample No. 456, Jakarta",
      "phone": "081298765432",
      "email": "jane.smith@example.com"
    },
    "analyst": "Dr. Analyst",
    "doctor": "Dr. Smith"
  }
}
```

**cURL Example:**

```bash
curl -X POST http://localhost:8080/work-orders \
  -H "Content-Type: application/json" \
  -d '{
    "no_order": "WO001",
    "test_code": ["HB", "LEUKOSIT", "ERITROSIT"],
    "patient": {
      "first_name": "Jane",
      "last_name": "Smith",
      "birth_date": "1985-03-20T00:00:00Z",
      "sex": "female",
      "address": "Jl. Sample No. 456, Jakarta",
      "phone": "081298765432",
      "email": "jane.smith@example.com"
    },
    "analyst": "Dr. Analyst",
    "doctor": "Dr. Smith"
  }'
```

---

### Get Work Order by No Order

Retrieve a specific work order by order number.

**Endpoint:** `GET /work-orders?no_order={no_order}`

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| no_order | string | Yes | Work order number |

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": {
    "no_order": "WO001",
    "test_code": ["HB", "LEUKOSIT", "ERITROSIT"],
    "patient": {
      "id": "770e8400-e29b-41d4-a716-446655440002",
      "first_name": "Jane",
      "last_name": "Smith",
      "birth_date": "1985-03-20T00:00:00Z",
      "sex": "female",
      "address": "Jl. Sample No. 456, Jakarta",
      "phone": "081298765432",
      "email": "jane.smith@example.com"
    },
    "analyst": "Dr. Analyst",
    "doctor": "Dr. Smith"
  }
}
```

**cURL Example:**

```bash
curl -X GET "http://localhost:8080/work-orders?no_order=WO001"
```

---

### Update Work Order

Update an existing work order and its associated patient information.

**Endpoint:** `PUT /work-orders?no_order={no_order}`

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| no_order | string | Yes | Work order number |

**Request Body:**

```json
{
  "test_code": ["HB", "LEUKOSIT", "TROMBOSIT"],
  "patient": {
    "first_name": "Jane",
    "last_name": "Smith",
    "birth_date": "1985-03-20T00:00:00Z",
    "sex": "female",
    "address": "Jl. Updated Address No. 789, Jakarta",
    "phone": "081298765432",
    "email": "jane.smith@example.com"
  },
  "analyst": "Dr. New Analyst",
  "doctor": "Dr. Smith"
}
```

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": {
    "no_order": "WO001",
    "test_code": ["HB", "LEUKOSIT", "TROMBOSIT"],
    "patient": {
      "id": "770e8400-e29b-41d4-a716-446655440002",
      "first_name": "Jane",
      "last_name": "Smith",
      "birth_date": "1985-03-20T00:00:00Z",
      "sex": "female",
      "address": "Jl. Updated Address No. 789, Jakarta",
      "phone": "081298765432",
      "email": "jane.smith@example.com"
    },
    "analyst": "Dr. New Analyst",
    "doctor": "Dr. Smith"
  }
}
```

**cURL Example:**

```bash
curl -X PUT "http://localhost:8080/work-orders?no_order=WO001" \
  -H "Content-Type: application/json" \
  -d '{
    "test_code": ["HB", "LEUKOSIT", "TROMBOSIT"],
    "patient": {
      "first_name": "Jane",
      "last_name": "Smith",
      "birth_date": "1985-03-20T00:00:00Z",
      "sex": "female",
      "address": "Jl. Updated Address No. 789, Jakarta",
      "phone": "081298765432",
      "email": "jane.smith@example.com"
    },
    "analyst": "Dr. New Analyst",
    "doctor": "Dr. Smith"
  }'
```

---

### Delete Work Order

Delete a work order and its associated patient record.

**Endpoint:** `DELETE /work-orders?no_order={no_order}`

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| no_order | string | Yes | Work order number |

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": {
    "message": "Work order deleted successfully"
  }
}
```

**cURL Example:**

```bash
curl -X DELETE "http://localhost:8080/work-orders?no_order=WO001"
```

---

### Get All Work Orders

Retrieve all work orders with their patient information.

**Endpoint:** `GET /work-orders`

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": [
    {
      "no_order": "WO001",
      "test_code": ["HB", "LEUKOSIT", "ERITROSIT"],
      "patient": {
        "id": "770e8400-e29b-41d4-a716-446655440002",
        "first_name": "Jane",
        "last_name": "Smith",
        "birth_date": "1985-03-20T00:00:00Z",
        "sex": "female",
        "address": "Jl. Sample No. 456, Jakarta",
        "phone": "081298765432",
        "email": "jane.smith@example.com"
      },
      "analyst": "Dr. Analyst",
      "doctor": "Dr. Smith"
    }
  ]
}
```

**cURL Example:**

```bash
curl -X GET http://localhost:8080/work-orders
```

---

### Get Work Orders by Doctor

Retrieve all work orders for a specific doctor.

**Endpoint:** `GET /work-orders?doctor={doctor_name}`

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| doctor | string | Yes | Doctor name |

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": [
    {
      "no_order": "WO001",
      "test_code": ["HB", "LEUKOSIT", "ERITROSIT"],
      "patient": {
        "id": "770e8400-e29b-41d4-a716-446655440002",
        "first_name": "Jane",
        "last_name": "Smith",
        "birth_date": "1985-03-20T00:00:00Z",
        "sex": "female",
        "address": "Jl. Sample No. 456, Jakarta",
        "phone": "081298765432",
        "email": "jane.smith@example.com"
      },
      "analyst": "Dr. Analyst",
      "doctor": "Dr. Smith"
    }
  ]
}
```

**cURL Example:**

```bash
curl -X GET "http://localhost:8080/work-orders?doctor=Dr.%20Smith"
```

---

### Get Work Orders by Analyst

Retrieve all work orders for a specific analyst.

**Endpoint:** `GET /work-orders?analyst={analyst_name}`

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| analyst | string | Yes | Analyst name |

**Success Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "data": [
    {
      "no_order": "WO001",
      "test_code": ["HB", "LEUKOSIT", "ERITROSIT"],
      "patient": {
        "id": "770e8400-e29b-41d4-a716-446655440002",
        "first_name": "Jane",
        "last_name": "Smith",
        "birth_date": "1985-03-20T00:00:00Z",
        "sex": "female",
        "address": "Jl. Sample No. 456, Jakarta",
        "phone": "081298765432",
        "email": "jane.smith@example.com"
      },
      "analyst": "Dr. Analyst",
      "doctor": "Dr. Smith"
    }
  ]
}
```

**cURL Example:**

```bash
curl -X GET "http://localhost:8080/work-orders?analyst=Dr.%20Analyst"
```

---

## Response Format

### Success Response

All successful API responses follow this format:

```json
{
  "code": 200,
  "status": "success",
  "data": { ... }
}
```

| Field  | Type         | Description                              |
| ------ | ------------ | ---------------------------------------- |
| code   | integer      | HTTP status code                         |
| status | string       | Always "success" for successful requests |
| data   | object/array | Response data (varies by endpoint)       |

### Error Response

All error responses follow this format:

```json
{
  "code": 400,
  "status": "error",
  "message": "Error description"
}
```

| Field   | Type    | Description                        |
| ------- | ------- | ---------------------------------- |
| code    | integer | HTTP status code                   |
| status  | string  | Always "error" for failed requests |
| message | string  | Error description                  |

---

## Error Codes

| Code | Description                                      |
| ---- | ------------------------------------------------ |
| 200  | OK - Request successful                          |
| 201  | Created - Resource created successfully          |
| 400  | Bad Request - Invalid request body or parameters |
| 404  | Not Found - Resource not found                   |
| 405  | Method Not Allowed - HTTP method not supported   |
| 500  | Internal Server Error - Server error occurred    |

---

## Common Error Examples

### Invalid Request Body

```json
{
  "code": 400,
  "status": "error",
  "message": "Invalid request body"
}
```

### Resource Not Found

```json
{
  "code": 404,
  "status": "error",
  "message": "patient not found"
}
```

### Missing Required Parameter

```json
{
  "code": 400,
  "status": "error",
  "message": "id parameter is required"
}
```

### Method Not Allowed

```json
{
  "code": 405,
  "status": "error",
  "message": "Method not allowed"
}
```

---

## Testing with Postman

Import the following collection to test the API:

1. Create a new Collection in Postman
2. Add the base URL as a collection variable: `{{base_url}}` = `http://localhost:8080`
3. Create requests for each endpoint listed above

### Example Environment Variables

```json
{
  "base_url": "http://localhost:8080",
  "patient_id": "550e8400-e29b-41d4-a716-446655440000",
  "no_order": "WO001"
}
```

---

## Notes

- All date fields use ISO 8601 format: `YYYY-MM-DDTHH:mm:ssZ`
- Sex field accepts only: `"male"` or `"female"`
- Patient IDs are automatically generated UUIDs
- Creating a work order automatically creates an associated patient
- Deleting a work order also deletes the associated patient (cascade delete)
- All timestamps are in UTC
