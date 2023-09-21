
# movie-xsis : Developed by Lutfi M

- Golang (gorilla mux)
- MySQL

## Application Manual

1. Set-up database (docs/movie.sql) :
   or Please make DB with name "movie"
2. How to Run :

   ```
   go run . or gor run main.go
   ```

## Endpoints List

[POST]      `localhost:7000/api/movie` : Add Movie </br>
[GET]       `localhost:7000/api/movie` : All Movie </br>
[GET]       `localhost:7000/api/movie/{ID}` : Movie by Id </br>
[PATCH]     `localhost:7000/api/movie/{ID}` : Edit Movie by Id </br>
[DELETE]    `localhost:7000/api/movie/{ID}` : Delete Movie by Id </br>
