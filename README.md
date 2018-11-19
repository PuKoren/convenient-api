# convenient-api
[![Build Status](https://travis-ci.org/PuKoren/convenient-api.svg?branch=master)](https://travis-ci.org/PuKoren/convenient-api)

An open API to ease UX form fillings by providing default values based on user input.

This API can be useful to pre-fill some form user input, like birth date, country/city, sex or civility, company name, etc.


Currently exposed as open API on [convenient.pukogames.com](http://convenient.pukogames.com)

## Routes
### Data enrichment

[POST] /user/v1
```bash
curl -X POST \
  http://convenient.pukogames.com/v1/user/ \
  -H 'Content-Type: application/json' \
  -d '{ "firstname": "David" }'
``` 

Return enriched data json as
```json
{
    "firstname": "David",
    "lastname": "",
    "country": "FR",
    "birthyear": 1972,
    "sex": "M"
}
```

Several fields are accepted as data POST to enrich data further:
```json
{
    "firstname": "string",
    "lastname": "string",
    "country": "2 char iso code",
    "birthyear": "integer",
    "email": "spam@pukogames.com"
}
```

The email field can enrich the company name, address, firstname and lastname, as well as potential age also.


## DB used
The project currently uses MaxMind DB, as well as French Insee public name databases.
