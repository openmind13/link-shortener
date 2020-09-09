# Link shortener

Link shortener api server

## Main api methods

Get new short url


'''
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"longurl": "https://avito.ru/moscow"}' \
  http://localhost:8080/create
'''

Response

'''
{"shorturl":"http://localhost:8080/68Ad2x9"}
'''

Get new custom url

'''
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
      "longurl": "https://start.avito.ru/tech",
      "shorturl": "trainee"}' \
  http://localhost:8080/createcustom
'''


Response

'''
{"shorturl":"http://localhost:8080/trainee"}
'''

GET

'''
curl --header "Content-Type: application/json" \
  --request GET \
  http://localhost:8080/{trainee}
'''

Redirect to longurl

## Testing

To test the entire application run:

'''
make test
'''