# Simple shop api demo using go

 https://gin-gonic.com/
 
 ### to exec
 ```bash
 go run main.go
 ```
 
 ## Methods
 
 - For products:
 ```
 POST: api/products -json ProductRequest -> Creates product (returns product response)
 GET:     api/products -> List product responses
 GET:     api/products/{productId} -> Returns product response
 DELETE:     api/products/{productId} -> Delete product and returns 1 or 0
 ```
 
 - For orders:
 ```
 POST: api/orders -json OrderRequest -> Creates order (returns order response)
 GET:     api/orders -> List order responses
 GET:     api/orders/{orderId} -> Returns order response
 DELETE:     api/orders/{orderId} -> Delete order and returns 1 or 0
 ```
 
 
 ## Requests
  - Product
 ```json
 {
      "name": "product name (string)" ,
      "price": "product price in PEN (number)",
      "description": "product desctiption (string)",
      "branch": "product brand (string)",
      "stars": "product stars (number)",
      "ProductCode": "string)",
      "imageUrl": "url (string)"
}
```
 - Order
 ```json
 {
    "created_by": "Username (string)",
    "created_at": "timestamp (2022-08-29T00:00:00.494Z)",
    "items": ["productid", "productid"]
}

```
## Requests
  - Product
 ```json
 {
      "id": "product id (string)",
      "name": "product name (string)" ,
      "price": "product price in PEN (number)",
      "description": "product desctiption (string)",
      "branch": "product brand (string)",
      "stars": "product stars (number)",
      "ProductCode": "string)",
      "imageUrl": "url (string)"
}
```
 - Order
 ```json
    {
        "id": "order id (string)",
        "created_by": "username (string)",
        "created_at": "timestamp (2022-08-29T00:00:00.494Z)",
        "invoice": "total price (number)",
        "items": [
            {
                "id": "productid (string)",
                "name": "Phone 13",
                "price": "product price in PEN (number)",
                "description": "product desctiption (string)",
                "branch": "product brand (string)",
                "stars": "product stars (number)",
                "ProductCode": "string)",
                "imageUrl": "url (string)"
            }
        ]
    }

```
