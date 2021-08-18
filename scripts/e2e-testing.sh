#!/bin/bash
main() {
  # Variables.
  WEB_URL=http://localhost:8080
  API_URL=http://localhost:8080/api/v1

  # Tests for web frontend.
  echo "Running e2e tests for the web frontend..."
  curl -X GET $WEB_URL

  # Give line breaks.
  echo

  # Tests for API backend.
  echo "Running e2e tests for the API backend (successful tests)..."
  curl -X GET $API_URL
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Kaede Kimura","address":"Tokyo"}' $API_URL

  # Another line break.
  echo

  # Expected to fail tests.
  echo "Running e2e tests for API backend (failure tests)..."
  curl -X POST -d '{"name":"Kaede Kimura"}' $API_URL
  curl -X POST -H "Content-Type: application/json" $API_URL
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Kaede Kimura"}' $API_URL
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Kaede Kimura","address":"Kyoto",badfomathere}' $API_URL
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Kaede Kimura","address":12345}' $API_URL
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Kaede Kimura","address":12345' $API_URL
  curl -X POST -H "Content-Type: application/json" -d '[{"name":"Kaede Kimura","address":"12345"},{"name":"Mai Sakurajima","address":"Fujisawa"}]' $API_URL
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Mai Sakurajima","mockAttribute":"Fujisawa"}' $API_URL
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Mai Sakurajima","address":"Fujisawa"}{"name":"Kamisato Ayaka","address":"Fukuoka"}' $API_URL
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Mai Sakurajima","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa","address":"Fujisawa"}' $API_URL
  curl -X PUT -H "Content-Type: application/json" -d '{"name":"Sayu Ogiwara","address":"Hokkaido"}' $API_URL
  curl -X GET $API_URL/404

  # Done tests.
  echo "Finished testing!"
}

# Execute main.
main
