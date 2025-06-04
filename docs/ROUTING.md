# API Routing

All routes are prefixed with `/api`. Authentication is handled via JWT in the `Authorization` header (`Bearer <token>`). Role-based access is enforced using Gin middleware. 

| Method | Path                     | Auth | Roles                     | Description                                     |
| ------ | ------------------------ | ---- | ------------------------- | ----------------------------------------------- |
| POST   | `/api/register`          | No   | —                         | Register a new user (Customer, Admin, Worker)   |
| POST   | `/api/login`             | No   | —                         | Authenticate user and return a JWT              |
| POST   | `/api/pizzas`            | Yes  | Admin, Worker             | Create a new pizza                              |
| GET    | `/api/pizzas`            | Yes  | Customer, Admin, Worker   | List all pizzas                                 |
| GET    | `/api/pizzas/:id`        | Yes  | Customer, Admin, Worker   | Get details of a single pizza by its ID         |
| PUT    | `/api/pizzas/:id`        | Yes  | Admin, Worker             | Update an existing pizza                        |
| DELETE | `/api/pizzas/:id`        | Yes  | Admin, Worker             | Delete a pizza by ID                            |
| POST   | `/api/orders`            | Yes  | Customer                  | Place a new order                               |
| GET    | `/api/orders`            | Yes  | Customer                  | List orders belonging to the authenticated user |
| PATCH  | `/api/orders/:id/status` | Yes  | Worker, Admin             | Update the status of an order                   |
| GET    | `/api/stats`             | Yes  | Admin                     | Retrieve total number of orders statistic       |
| GET    | `/api/profile/history`   | Yes  | Customer                  | Get order history for the authenticated user    |

## Route Details

### 1. Register User

- **URL**: `/api/register`
- **Method**: `POST`
- **Auth**: None
- **Body**:
```json
  {
    "email": "user@example.com",
    "password": "securepassword",
    "role": "customer"  // one of "customer", "admin", "worker" | Will be removed later, this serve only for the MVP
  }
```

* **Success Response**: `HTTP 201 Created`
* **Error Responses**:

  * `400 Bad Request` (invalid payload)
  * `500 Internal Server Error` (creation failure)

---

### 2. Login User

* **URL**: `/api/login`
* **Method**: `POST`
* **Auth**: None
* **Body**:

  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
* **Success Response**: `HTTP 200 OK`

  ```json
  {
    "token": "<jwt-token>"
  }
  ```
* **Error Responses**:

  * `400 Bad Request` (invalid payload)
  * `401 Unauthorized` (invalid credentials)

---

### 3. Create Pizza

* **URL**: `/api/pizzas`
* **Method**: `POST`
* **Auth**: Bearer JWT
* **Roles**: Admin, Worker
* **Body**:

  ```json
  {
    "name": "Margherita",
    "description": "Tomato, mozzarella, basil",
    "price": 9.50
  }
  ```
* **Success Response**: `HTTP 201 Created`
* **Error Responses**:

  * `400 Bad Request` (invalid payload)
  * `401 Unauthorized` / `403 Forbidden` (missing/invalid token or insufficient role)
  * `500 Internal Server Error` (database error)

---

### 4. List All Pizzas

* **URL**: `/api/pizzas`
* **Method**: `GET`
* **Auth**: Bearer JWT
* **Roles**: Customer, Admin, Worker
* **Success Response**: `HTTP 200 OK`

  ```json
  [
    {
      "ID": 1,
      "CreatedAt": "2025-06-03T10:00:00Z",
      "UpdatedAt": "2025-06-03T10:00:00Z",
      "DeletedAt": null,
      "Name": "Margherita",
      "Description": "Tomato, mozzarella, basil",
      "Price": 9.50
    },
    ...
  ]
  ```
* **Error Responses**:

  * `401 Unauthorized` / `403 Forbidden`
  * `500 Internal Server Error`

---

### 5. Get Pizza by ID

* **URL**: `/api/pizzas/:id`
* **Method**: `GET`
* **Auth**: Bearer JWT
* **Roles**: Customer, Admin, Worker
* **URL Params**:

  * `id` (integer) – ID of the pizza to retrieve
* **Success Response**: `HTTP 200 OK`

  ```json
  {
    "ID": 1,
    "CreatedAt": "2025-06-03T10:00:00Z",
    "UpdatedAt": "2025-06-03T10:00:00Z",
    "DeletedAt": null,
    "Name": "Margherita",
    "Description": "Tomato, mozzarella, basil",
    "Price": 9.50
  }
  ```
* **Error Responses**:

  * `401 Unauthorized` / `403 Forbidden`
  * `404 Not Found` (no pizza with given ID)
  * `500 Internal Server Error`

---

### 6. Update Pizza

* **URL**: `/api/pizzas/:id`
* **Method**: `PUT`
* **Auth**: Bearer JWT
* **Roles**: Admin, Worker
* **URL Params**:

  * `id` (integer) – ID of the pizza to update
* **Body**:

  ```json
  {
    "name": "Margherita Special",
    "description": "Tomato, mozzarella, basil, extra cheese",
    "price": 11.00
  }
  ```
* **Success Response**: `HTTP 200 OK`
* **Error Responses**:

  * `400 Bad Request` (invalid payload)
  * `401 Unauthorized` / `403 Forbidden`
  * `404 Not Found`
  * `500 Internal Server Error`

---

### 7. Delete Pizza

* **URL**: `/api/pizzas/:id`
* **Method**: `DELETE`
* **Auth**: Bearer JWT
* **Roles**: Admin, Worker
* **URL Params**:

  * `id` (integer) – ID of pizza to delete
* **Success Response**: `HTTP 204 No Content`
* **Error Responses**:

  * `401 Unauthorized` / `403 Forbidden`
  * `404 Not Found`
  * `500 Internal Server Error`

---

### 8. Create Order

* **URL**: `/api/orders`
* **Method**: `POST`
* **Auth**: Bearer JWT
* **Roles**: Customer
* **Body**:

  ```json
  {
    "items": [
      { "pizza_id": 1, "quantity": 2 },
      { "pizza_id": 3, "quantity": 1 }
    ]
  }
  ```
* **Success Response**: `HTTP 201 Created`
* **Error Responses**:

  * `400 Bad Request` (invalid payload)
  * `401 Unauthorized` / `403 Forbidden`
  * `500 Internal Server Error`

---

### 9. List Customer Orders

* **URL**: `/api/orders`
* **Method**: `GET`
* **Auth**: Bearer JWT
* **Roles**: Customer
* **Success Response**: `HTTP 200 OK`

  ```json
  [
    {
      "ID": 10,
      "CreatedAt": "2025-06-03T11:15:00Z",
      "UpdatedAt": "2025-06-03T11:20:00Z",
      "DeletedAt": null,
      "UserID": 5,
      "Status": "preparing",
      "Items": [
        {
          "ID": 22,
          "CreatedAt": "2025-06-03T11:15:00Z",
          "UpdatedAt": "2025-06-03T11:15:00Z",
          "DeletedAt": null,
          "OrderID": 10,
          "PizzaID": 1,
          "Quantity": 2,
          "Pizza": {
            "ID": 1,
            "Name": "Margherita",
            "Price": 9.50
          }
        }
      ]
    },
    ...
  ]
  ```
* **Error Responses**:

  * `401 Unauthorized` / `403 Forbidden`
  * `500 Internal Server Error`

---

### 10. Update Order Status

* **URL**: `/api/orders/:id/status`
* **Method**: `PATCH`
* **Auth**: Bearer JWT
* **Roles**: Worker, Admin
* **URL Params**:

  * `id` (integer) – ID of the order to update
* **Body**:

  ```json
  {
    "status": "ready"  // one of "pending", "preparing", "ready", "done"
  }
  ```
* **Success Response**: `HTTP 200 OK`
* **Error Responses**:

  * `400 Bad Request` (invalid payload)
  * `401 Unauthorized` / `403 Forbidden`
  * `404 Not Found`
  * `500 Internal Server Error`

---

### 11. Get Total Orders Statistic

* **URL**: `/api/stats`
* **Method**: `GET`
* **Auth**: Bearer JWT
* **Roles**: Admin
* **Success Response**: `HTTP 200 OK`

  ```json
  {
    "total_orders": 1234
  }
  ```
* **Error Responses**:

  * `401 Unauthorized` / `403 Forbidden`
  * `500 Internal Server Error`

---

### 12. Get Order History (Customer Profile)

* **URL**: `/api/profile/history`
* **Method**: `GET`
* **Auth**: Bearer JWT
* **Roles**: Customer
* **Success Response**: `HTTP 200 OK`

  ```json
  [
    {
      "ID": 10,
      "CreatedAt": "2025-06-03T11:15:00Z",
      "UpdatedAt": "2025-06-03T11:20:00Z",
      "DeletedAt": null,
      "UserID": 5,
      "Status": "preparing",
      "Items": [ /* same as List Customer Orders */ ]
    },
    ...
  ]
  ```
* **Error Responses**:

  * `401 Unauthorized` / `403 Forbidden`
  * `500 Internal Server Error`
