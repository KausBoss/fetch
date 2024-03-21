# Receipt Processor Web Service

This is a web service for processing receipts and calculating points based on specified rules. It provides two endpoints, one to process receipts and other to retrieve points awarded for each receipt.

## Running the Application with Docker

To run this application using Docker, follow these steps:

1. **Clone the Repository**:

```git clone https://github.com/fetch-rewards/receipt-processor-challenge.git```


2. **Build the Docker Image**:

```docker build -t receipt-processor .```


3. **Run the Docker Container**:

```docker run -p 8080:8080 receipt-processor```


4. **Access the Endpoints**:
- The application will be running at `http://localhost:8080`.
- Use the following endpoints:
    - `/receipts/process`: POST endpoint to process receipts. Provide a JSON payload of the receipt details.
    - `/receipts/{id}/points`: GET endpoint to retrieve points awarded for a specific receipt by its ID.

## API Endpoints

### Process Receipts:
- **Endpoint**: `/receipts/process`
- **Method**: POST
- **Payload**: JSON containing receipt details (Retailer, PurchaseDate, PurchaseTime, Items, Total)
- **Response**: JSON object containing the ID generated for the receipt


### Get Points:
- **Endpoint**: `/receipts/{id}/points`
- **Method**: GET
- **Response**: JSON object containing the number of points awarded for the receipt with the specified ID

## Dependencies
- `github.com/google/uuid`: Used for generating UUIDs.
- `github.com/gorilla/mux`: Used for routing HTTP requests.

