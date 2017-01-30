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
| matchers  | true       | matchers list          | Array   |
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

### Get Silence 

* **URL** /api/v1/silence/{sid}

* **Method** GET

* **Body**

  None

* **Response**

GET /api/v1/silence/2a966b4b-23b2-441f-9af5-e08b0a73aa06

* ***success 200***
```json
{  
   "status":"success",
   "data":{  
      "id":"2a966b4b-23b2-441f-9af5-e08b0a73aa06",
      "matchers":[  
         {  
            "name":"service",
            "value":"dash",
            "isRegex":false
         }
      ],
      "startsAt":"2017-01-30T06:37:14.95697622Z",
      "endsAt":"2017-01-31T04:15:00Z",
      "updatedAt":"2017-01-30T06:37:14.95697622Z",
      "createdBy":"rsamal@us.ibm.com",
      "comment":"Maintenance window"
   }
}

```

* ***NotFound 404***

Error getting silence:

### Delete Silence 

* **URL** /api/v1/silence/{sid}

* **Method** DELETE

* **Body**

  None

* **Response**

GET /api/v1/silence/2a966b4b-23b2-441f-9af5-e08b0a73aa06

* ***success 200***
```json
{"status":"success"}
```

* ***BadRequest 400***
```json
{  
   "status":"error",
   "errorType":"bad_data",
   "error":"not found"
}
```

## Silencer/Maintenance API in batch only with hostname and ip


### Add Silence in batch

* **URL** /api/v1/silences/batch

* **Method** POST

* **Body**

| Parameter | Required   | Description            | Type    |
|-----------|------------|------------------------|---------|
| hosts     | true       | list of hosts          | Array   |
| name      | true       | name of the host 	  | string  |
| ip        | true       | IP of the host         | string  | 
| startsAt  | false      | Start time in rfc3339  | string  |
| endsAt 	| false      | End time in rfc3339 	  | string  |
| createdBy | false      | Creater email          | string  |
| comment | false        | comment                | String  |


* **Response**

POST /api/v1/silences/batch

```json
{ 
  "hosts": [
    {"name" : "node0", "ip" :"172.30.1.0"},
    {"name" : "node1", "ip" : "172.30.1.1"},
    {"name" : "node2", "ip" : "172.30.1.2"}
  ],  

  "endsAt" :   "2017-01-30T10:15:00.000-08:00",
  "createdBy" : "rsamal@us.ibm.com",
  "comment" : "some comment"
  
}


```

* ***success 200***
```json
[  
   {  
      "id":"72f28c51-e649-4d15-bc42-4c77c68a446b",
      "name":"node0",
      "ip":"172.30.1.0"
   },
   {  
      "id":"36534064-c8fb-492e-9018-557eed309a6f",
      "name":"node1",
      "ip":"172.30.1.1"
   },
   {  
      "id":"87c80dbd-8924-4653-8f40-52a651415c40",
      "name":"node2",
      "ip":"172.30.1.2"
   }
]
```
* ***BadRequest 400***
``` json
{  
   "status":"error",
   "errorType":"bad_data",
   "error":"Invalid IP address 172.a.1.2"
}
```

* ***ServerError 500***
``` json
{  
   "status":"error",
   "errorType":"server_error",
   "error":"new silence must not start in the past"
}
```


### Delete Silence in batch

* **URL** /api/v1/silences/batch

* **Method** DELETE

* **Body**

| Parameter | Required   | Description            | Type    |
|-----------|------------|------------------------|---------|
|           | true       | list of IDs            | Array   |
| id        | true       | silence ID       	  | string  |

* **Response**

DELETE /api/v1/silences/batch

```json 
[  {  
      "id":"097b2ac5-142b-46d0-903a-4956d6bec768"
   },
   {  
      "id":"a9a6dc53-2997-4cd1-bb93-05635d0801db"
   },
   {  
      "id":"bb992c1e-c4af-425e-a610-9e9a16881bee"
   }
]

```

* ***success 200***
```json
{  
   "status":"Success",
   "message":"Deleted Silencers"
}
```
* ***BadRequest 400***
``` json
{  
   "status":"error",
   "errorType":"bad_data",
   "error":"bb992c1e-c4af-425e-a610-9e9a16881bee11 not found"
}
```

* ***ServerError 500***
``` json
{  
   "status":"error",
   "errorType":"server_error",
   "error":"error in deletion"
}
```

### Add Silence in batch

* ***Call the Delete batch API***
* ***Call create batch API with new updated data***


## Alert API
