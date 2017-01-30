# Alertmanager API Document v1 

## Silencer/Maintenance Window API

### Get Silences

* **URL** /api/v1/silences

* **Method** GET

* **Body Params**

None

* **Body**

None

* **Response**

GET /api/v1/silences

```json
{  
   "status":"success",
   "data":[  
      {  
         "id":"4008a369-6d9c-4a06-b783-88ebb07ace2a",
         "matchers":[  
            {  
               "name":"name",
               "value":"node0*",
               "isRegex":true
            },
            {  
               "name":"ip",
               "value":"192.67.0.1",
               "isRegex":false
            }
         ],
         "startsAt":"2017-01-30T06:09:22.797297Z",
         "endsAt":"2017-01-30T10:08:00Z",
         "updatedAt":"2017-01-30T06:09:22.797297Z",
         "createdBy":"rsamal@us.ibm.com",
         "comment":"testing"
      }
   ]
}

```

### Add Silence 

* **URL** /api/v1/silences

* **Method** POST

* **Body**

| Parameter | Required   | Description            | Type    |
|-----------|------------|------------------------|---------|
| matchers  | true       | matchers list          | string  |
| name      | true       | matcher name 		  | string  |
| value     | true       | matacher value         | string  | 
| isRegex   | false      | True for regex         | boolean |
| startsAt  | false      | Start time in rfc3339  | string  |
| endsAt 	| false      | End time in rfc3339 	  | string  |
| createdBy | false      | Creater email          | string  |
| comment | false        | comment                | String  |


* **Response**

POST /api/v1/silences

```json
{ 
  "matchers" : [
    {
      "name" : "service",
      "value" :"dash",
      "isRegex" : false
    }
  ],
    
  
   "endsAt" :   "2017-01-30T20:15:00.000-08:00",
  "createdBy" : "rsamal@us.ibm.com",
  "comment" : "Maintenance window"
  
}

```

* ***success 200***
```json
{  
   "status":"success",
   "data":{  
      "silenceId":"5947eba9-c472-42c4-88fe-ad9b8cd781b0"
   }
}
```
* ***BadRequest 400***
``` json
{  
   "status":"error",
   "errorType":"bad_data",
   "error":"invalid character '}' looking for beginning of object key string"
}
```

* ***ServerError 500***
``` json
{  
   "status":"error",
   "errorType":"server_error",
   "error":"unexpected ID in new silence"
}
```
