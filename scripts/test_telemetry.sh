#!/bin/bash
cat << "EOF"
   ,-.
     / \  `.  __..-,O
    :   \ --''_..-'.'
    |    . .-' `. '.
    :     .     .`.'
     \     `.  /  ..
      \      `.   ' .
       `,       `.   \
      ,|,`.        `-.\
     '.||  ``-...__..-`
      |  |
      |__|
      /||\
     //||\\
    // || \\
 __//__||__\\__
'--------------' developed by Gabriel Rocha, art by https://www.asciiart.eu/space/telescopes

EOF

URL="http://localhost:3000"

function send_gps() {
  echo "Sending GPS data..."
  curl -X POST "$URL/telemetry/gps" \
    -H "Content-Type: application/json" \
    -d '{
      "mac_address": "AB:CD:EF:12:34:56",
      "timestamp": "2025-06-17T14:00:00Z",
      "latitude": -28.4827,
      "longitude": -49.0144
    }'
  echo -e "\n[✓] Request sent to /telemetry/gps"
}

function send_gyroscope() {
  echo "Sending gyroscope data..."
  curl -X POST "$URL/telemetry/gyroscope" \
    -H "Content-Type: application/json" \
    -d '{
      "mac_address": "AB:CD:EF:12:34:56",
      "timestamp": "2025-06-17T14:00:00Z",
      "x": 0.123,
      "y": -0.456,
      "z": 1.789
    }'
  echo -e "\n[✓] Request sent to /telemetry/gyroscope"
}

while true; do
  echo ""
  echo "===== TELEMETRY TEST ====="
  echo "1. Send GPS data"
  echo "2. Send Gyroscope data"
  echo "0. Exit"
  echo "========================================"
  read -p "Choose an option: " option

  case $option in
    1) send_gps ;;
    2) send_gyroscope ;;
    0) echo "Exiting..."; exit 0 ;;
    *) echo "Invalid option." ;;
  esac
done
