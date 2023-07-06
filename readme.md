The goal of this project is to build a thin web server that allows
me to quickly map API endpoints to SQL queries. Using a config file
that looks like

```json
// config.json

{
  endpoint: "/orders/{id}",
  methods: ["get"],
  queryBy: "id",
  query: "select id, total_price, item_name, placed_at, status from orders",
},
{
  endpoint: "/orders/",
  methods: ["get"],
  pagination: true, // enables pagination URL params
  filterBy: ["store_name"], // allows fuzzy searching on these columns
  sortBy: ["placed_at"], // allows sorting by this value
  query: "select id, total_price, item_name from orders",
},
{
  endpoint: "/orders",
  method: ["post"],
  inputFields: ["total_price", "item"],
},
{
  endpoint: "/orders",
  method: ["put"],
  inputFields: ["total_price"],
},
...
```

Upon opening, the server loads up all of the configured endpoints
and creates a handler for each one

```golang
for _, endpointConfig := range endpoints {
  endpoint := endpointConfig["endpoint"]

  // Can do some other stuff such as pagination, filtering, sorting

  http.HandleFunc(endpoint, funcHandler)
}
```

`method: ["get"]`







