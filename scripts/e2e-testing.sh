#!/bin/bash
main() {
  # Variables.
  API_URL=http://localhost:8080/api/v1

  # Tests for API backend.
  echo "Running e2e inference test for the API backend..."
  curl -X POST -H "Content-Type: application/json" -d \
    '{
      "symptoms": [
        {
          "symptomId": "S1",
          "weight": 0.2
        },
        {
          "symptomId": "S2",
          "weight": 0.2
        },
        {
          "symptomId": "S3",
          "weight": 0.2
        },
        {
          "symptomId": "S4",
          "weight": 0.4
        },
        {
          "symptomId": "S5",
          "weight": 0.2
        },
        {
          "symptomId": "S6",
          "weight": 0.4
        },
        {
          "symptomId": "S7",
          "weight": 0.8
        },
        {
          "symptomId": "S8",
          "weight": 0.2
        },
        {
          "symptomId": "S9",
          "weight": 0.2
        },
        {
          "symptomId": "S10",
          "weight": 0.4
        },
        {
          "symptomId": "S11",
          "weight": 0.2
        },
        {
          "symptomId": "S12",
          "weight": 0.2
        },
        {
          "symptomId": "S13",
          "weight": 1
        }
      ]
    }' \
  http://localhost:8080/api/v1

  # Another line break.
  echo

  # Done tests.
  echo "Finished testing!"
}

# Execute main.
main