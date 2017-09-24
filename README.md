# Install

....


# API DOC

## All Products

GET - /restaurants/1/products

### Response Body

```json
{
   "id":1,
   "restaurant_id":1,
   "title":"Pizza",
   "description":"Pizza boa",
   "image":"http://image.com",
   "price_cents":2000,
   "ingredient_groups":[
      {
         "id":1,
         "title":"Queijo",
         "basic":false,
         "product_id":1,
         "ingredients":[
            {
               "id":2,
               "name":"Muçarela",
               "price_cents":50
            },
            {
               "id":1,
               "name":"Gorgonzola",
               "price_cents":140
            }
         ]
      }
   ]
}
```

## Show Product

GET - /restaurants/1/products/1

### Response Body

```json
[
   {
      "id":2,
      "restaurant_id":1,
      "title":"Pizza Presunto",
      "description":"Pizza de preseunto",
      "image":"http://image.com",
      "price_cents":3000,
      "ingredient_groups":null
   },
   {
      "id":1,
      "restaurant_id":1,
      "title":"Pizza",
      "description":"Pizza boa",
      "image":"http://image.com",
      "price_cents":2000,
      "ingredient_groups":null
   }
]
```

## Create order

POST - /orders

### Payload:

```json
{
   "user_id":"123das",
   "restaurant_id":1,
   "products":[
      {
         "product_id":1,
         "quantity":2,
         "ingredients":[
            {
               "id":1
            },
            {
               "id":5
            }
         ]
      },
      {
         "product_id":2,
         "quantity":1
      }
   ]
}
```

### Response Body

```json
{
   "id":5,
   "user_id":"123das",
   "restaurant_id":1,
   "status":"requested",
   "products":[
      {
         "id":9,
         "product_id":1,
         "order_id":5,
         "quantity":2,
         "total_price_cents":4140,
         "ingredients":[
            {
               "id":1,
               "name":"Gorgonzola",
               "price_cents":140
            }
         ]
      },
      {
         "id":10,
         "product_id":2,
         "order_id":5,
         "quantity":1,
         "total_price_cents":3000,
         "ingredients":null
      }
   ]
}
```

## Update product's order quantity

PUT - /restaurant/{restaurant_id}/orders/{order_id}/products/{product_order_id}/quantity

#### Payload:

```
{
  "user_id": "123das",
  "quantity": 12
}
```