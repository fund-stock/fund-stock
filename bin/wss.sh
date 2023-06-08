#curl --include \
#  --no-buffer \
#  --header "Connection: Upgrade" \
#  --header "Upgrade: websocket" \
#  --header "Host: stream.binance.com:9443" \
#  --header "Origin: https://stream.binance.com:9443" \
#  --header "Sec-WebSocket-Key: NVwjmQUcWCenfWu98asDmg==" \
#  --header "Sec-WebSocket-Version: 13" \
#  https://stream.binance.com:9443/ws/ethusd@kline_1m
#
#
curl --include \
     --no-buffer \
     --header "Connection: Upgrade" \
     --header "Upgrade: websocket" \
     --header "Sec-WebSocket-Key: NVwjmQUcWCenfWu98asDmg==" \
     --header "Sec-WebSocket-Version: 13" \
     https://goclient.smarttradess.com/app/v2/trade/wsIndex.json?token=d278b72f18814eed9041b1d345ff2aef&userId=15DC8EF440C

#curl --include \
#     --no-buffer \
#     --header "Connection: Upgrade" \
#     --header "Upgrade: websocket" \
#     --header "Host: 8.215.38.97:8080" \
#     --header "Origin: http://8.215.38.97:8080" \
#     --header "Sec-WebSocket-Key: NVwjmQUcWCenfWu98asDmg==" \
#     --header "Sec-WebSocket-Version: 13" \
#     http://8.215.38.97:8080/v1/api/kline/ws_index



#curl --location --request POST 'https://api.adjust.com/public/v1/apps/yxs12pfewq/trackers'\
#--header 'Authorization: Token token={API_TOKEN}' \
#--header 'Content-Type: application/json' \
#--data-raw '{
#  "name": "Adroll"
#}'