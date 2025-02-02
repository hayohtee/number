# number
number is a REST API that takes a number and returns interesting mathematical properties about it, along with a fun fact.

## How to run
1. Ensure the GO SDK is installed
2. Clone the repository
   ```bash
   git clone https://github.com/hayohtee/number.git
   ```
3. Change into the directory
   ```bash
   cd number
   ```
4. Build the project
   ```bash
   go build -o api ./...
   ```
5. Run the program\
   The server should start on http://localhost:4000, send a GET request to the endpoint to retrieve the information
   ```bash
   ./api
   ```

## API Documentation
Endpoint URL
```bash
GET http://localhost:4000/api/classify-number?number=371
```
Response
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,  // sum of its digits
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371" //gotten from the numbers API
}
```
