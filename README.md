# Exoplanet Management Service

This is a microservice for managing different types of exoplanets. The service provides functionality to add, list, retrieve, update, delete exoplanets, and estimate fuel costs for trips.

## Features

- Add a new exoplanet
- List all available exoplanets
- Get information about a specific exoplanet by its unique ID
- Update the details of an existing exoplanet
- Delete an exoplanet from the catalog
- Estimate fuel costs for a trip to a particular exoplanet

## Requirements

- Gin web framework

## Project Structure 

```
exoplanet-service/
├── main.go
├── common/
│   └── constant.go
├── handler/
│   └── exoplanet.go
├── model/
│   └── exoplanet.go
├── repository/
│   |── exoplanet.go
|   └── store.go
├── service/
│   └── fuel.go
├── go.mod
└── go.sum
```
## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/mnsgoyal/exoplanet-service.git
    cd exoplanet-service
    ```

2. Install the dependencies:

    ```sh
    go mod tidy
    ```

3. Run the application:

    ```sh
    go run main.go
    ```

    The service will start on port 8080 by default.

## API Endpoints

### Add an Exoplanet

- **URL:** `/exoplanets`
- **Method:** `POST`
- **Request Body:**

    ```json
    {
        "name": "Exoplanet Name",
        "description": "Description of the exoplanet",
        "distance": 123,
        "radius": 0.15,
        "mass": 0.11,  
        "type": "Terrestrial" 
    }
    ```

- **Response:**

    ```json
    {
        "id": 1,
        "name": "Exoplanet Name",
        "description": "Description of the exoplanet",
        "distance": 123,
        "radius": 0.15,
        "mass": 0.11,
        "type": "Terrestrial"
    }
    ```

### List Exoplanets

- **URL:** `/exoplanets`
- **Method:** `GET`
- **Response:**

    ```json
    [
        {
            "id": 1,
            "name": "Exoplanet Name",
            "description": "Description of the exoplanet",
            "distance": 123,
            "radius": 0.15,
            "mass": 0.11,
            "type": "Terrestrial"
        }
    ]
    ```

### Get Exoplanet by ID

- **URL:** `/exoplanets/:id`
- **Method:** `GET`
- **Response:**

    ```json
    {
        "id": 1,
        "name": "Exoplanet Name",
        "description": "Description of the exoplanet",
        "distance": 123,
        "radius": 0.15,
        "mass": 0.11,
        "type": "Terrestrial"
    }
    ```

### Update Exoplanet

- **URL:** `/exoplanets/:id`
- **Method:** `PUT`
- **Request Body:**

    ```json
    {
        "name": "Updated Exoplanet Name",
        "description": "Updated description",
        "distance": 123,
        "radius": 0.16,
        "mass": 0.11,
        "type": "Terrestrial"
    }
    ```

- **Response:**

    ```json
    {
        "id": 1,
        "name": "Updated Exoplanet Name",
        "description": "Updated description",
        "distance": 123,
        "radius": 0.16,
        "mass": 0.11,
        "type": "Terrestrial"
    }
    ```

### Delete Exoplanet

- **URL:** `/exoplanets/:id`
- **Method:** `DELETE`
- **Response:** `204 No Content`

### Fuel Estimation

- **URL:** `/exoplanets/fuel/:id`
- **Method:** `GET`
- **Query Parameter:**

    - `crewCapacity`: Number of crew members

- **Response:**

    ```json
    {
        "fuelCost": 20079583.71
    }
    ```

