#!/bin/bash
main() {
  # Variables.
  API_URL=http://localhost:8080/api/v1

  # Tests for API backend.
  echo "Running e2e inference test for the API backend..."
  curl -X POST -H "Content-Type: application/json" -d \
    '{
      "diseaseId": "D01",
      "locale": "en",
      "symptoms": [
        {
          "symptomId": "S1",
          "weight": 0.25
        },
        {
          "symptomId": "S2",
          "weight": 0.25
        },
        {
          "symptomId": "S3",
          "weight": 0.25
        },
        {
          "symptomId": "S4",
          "weight": 0.25
        },
        {
          "symptomId": "S5",
          "weight": 0.25
        },
        {
          "symptomId": "S6",
          "weight": 0.25
        },
        {
          "symptomId": "S7",
          "weight": 0.25
        },
        {
          "symptomId": "S8",
          "weight": 0
        },
        {
          "symptomId": "S9",
          "weight": 0
        },
        {
          "symptomId": "S10",
          "weight": 0
        },
        {
          "symptomId": "S11",
          "weight": 0
        },
        {
          "symptomId": "S12",
          "weight": 0.5
        },
        {
          "symptomId": "S13",
          "weight": 0.2
        }
      ]
    }' \
  $API_URL

  # Another line break.
  echo

  # Done tests.
  echo "Finished testing!"
}

# Execute main.
main