
# movie-xsis : Developed by Lutfi M

- Golang (gorilla mux)
- MySQL

## Application Manual

1. Please make DB with name "db_movie"
2. How to Run :

   ```
   go run . / go run main.go
   ```

## Endpoints List

[POST]      `localhost:7000/Movies` : Add Movie </br>
[GET]       `localhost:7000/Movies` : All Movie </br>
[GET]       `localhost:7000/Movies/{ID}` : Movie by Id </br>
[PATCH]     `localhost:7000/Movies/{ID}` : Edit Movie by Id </br>
[DELETE]    `localhost:7000/Movies/{ID}` : Delete Movie by Id </br>
