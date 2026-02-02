# LIS (Laboratory Information System)

Laboratory Information System menggunakan Go dengan MySQL (tanpa ORM/GORM).

## Struktur Project

```
lis/
├── cmd/
│   └── server/
│       └── main.go              # Entry point aplikasi
├── internal/
│   ├── domain/
│   │   ├── dto/                 # Data Transfer Objects
│   │   │   ├── patient.go
│   │   │   ├── patient_mapper.go
│   │   │   ├── work_order.go
│   │   │   ├── work_order_mapper.go
│   │   │   └── response.go
│   │   └── entitiy/             # Domain entities
│   │       ├── patient.go
│   │       └── work_order.go
│   ├── handler/                 # HTTP handlers
│   │   ├── patient_handler.go
│   │   └── work_order_handler.go
│   ├── usecase/                 # Business logic
│   │   ├── patient_usecase.go
│   │   ├── patient_usecase_impl.go
│   │   ├── work_order_usecase.go
│   │   └── work_order_usecase_impl.go
│   └── repository/              # Database operations
│       ├── mysql.go
│       ├── patient_repository.go
│       ├── patient_repository_impl.go
│       ├── work_order_repository.go
│       └── work_order_repository_impl.go
├── migrations/
│   └── 001_create_tables.sql   # Database migrations
├── examples/
│   ├── repository_example.go
│   └── usecase_example.go
├── go.mod
└── go.sum
```

## Setup Database

1. Buat database MySQL:

```sql
CREATE DATABASE lis_db;
```

2. Jalankan migration:

```bash
mysql -u root -p lis_db < migrations/001_create_tables.sql
```

## Dependencies

```bash
go get github.com/go-sql-driver/mysql
go get github.com/google/uuid
```

## Menjalankan Server

```bash
go run cmd/server/main.go
```

Server akan berjalan di `http://localhost:8080`

## API Endpoints

### Patients

#### Create Patient

```bash
POST /patients
Content-Type: application/json

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

#### Get Patient by ID

```bash
GET /patients?id=<patient_id>
```

#### Update Patient

```bash
PUT /patients?id=<patient_id>
Content-Type: application/json

{
  "first_name": "John",
  "last_name": "Doe",
  "birth_date": "1990-05-15T00:00:00Z",
  "sex": "male",
  "address": "Jl. Updated No. 456, Jakarta",
  "phone": "081298765432",
  "email": "john.doe@example.com"
}
```

#### Delete Patient

```bash
DELETE /patients?id=<patient_id>
```

#### Get All Patients

```bash
GET /patients
```

#### Search Patients

```bash
GET /patients?q=<search_query>
```

### Work Orders

#### Create Work Order (dengan Patient)

```bash
POST /work-orders
Content-Type: application/json

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

#### Get Work Order by No Order

```bash
GET /work-orders?no_order=<no_order>
```

#### Update Work Order

```bash
PUT /work-orders?no_order=<no_order>
Content-Type: application/json

{
  "test_code": ["HB", "LEUKOSIT", "TROMBOSIT"],
  "patient": {
    "first_name": "Jane",
    "last_name": "Smith",
    "birth_date": "1985-03-20T00:00:00Z",
    "sex": "female",
    "address": "Jl. Updated No. 789, Jakarta",
    "phone": "081298765432",
    "email": "jane.smith@example.com"
  },
  "analyst": "Dr. New Analyst",
  "doctor": "Dr. Smith"
}
```

#### Delete Work Order

```bash
DELETE /work-orders?no_order=<no_order>
```

#### Get All Work Orders

```bash
GET /work-orders
```

#### Get Work Orders by Doctor

```bash
GET /work-orders?doctor=<doctor_name>
```

#### Get Work Orders by Analyst

```bash
GET /work-orders?analyst=<analyst_name>
```

### Health Check

```bash
GET /health
```

## Response Format

### Success Response

```json
{
  "code": 200,
  "status": "success",
  "data": { ... }
}
```

### Error Response

```json
{
  "code": 400,
  "status": "error",
  "message": "Error description"
}
```

## Fitur Utama

- ✅ Native Go `database/sql` (tanpa GORM)
- ✅ Transaction management di usecase layer
- ✅ DTO pattern untuk request/response
- ✅ Clean Architecture (Handler → Usecase → Repository)
- ✅ Auto-generate UUID untuk ID
- ✅ Work Order otomatis create Patient
- ✅ Foreign key constraints
- ✅ Connection pooling
- ✅ Context support untuk timeout/cancellation

## Konfigurasi

Edit konfigurasi database di `cmd/server/main.go`:

```go
config := repository.MySQLConfig{
    Host:     "localhost",
    Port:     "3306",
    User:     "root",
    Password: "password",
    Database: "lis_db",
}
```
