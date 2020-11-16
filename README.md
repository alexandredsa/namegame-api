# NAMEGAME-API

## About

## REST

- `POST` - `/rooms/create`

```
Headers: {
    "FCM_USER_TOKEN": "dkQJJ-3BLbg:APA91bFtqfdcku1WMPX2CkuCtJO9EomUucrh-aFs3X3mMJj636MPR7jbkRY"
}
```
```
{
    "username": "EdilsonCapetinha"
}
```

- `Response`
```
{
  "room": {
    "code": "TJKLAL"
  },
  "scoreboard": [
    {
      "user": {
        "name": "RenataFan",
        "state": "READY"
      },
      "score": 0
    },
    {
      "user": {
        "name": "CraqueNeto",
        "state": "PENDING"
      },
      "score": 0
    }
  ]
}
```

- `POST` - `/rooms/:room_code/join`
```
Headers: {
    "FCM_USER_TOKEN": "dkQJJ-3BLbg:APA91bFtqfdcku1WMPX2CkuCtJO9EomUucrh-aFs3X3mMJj636MPR7jbkRY"
}
```
```
{
    "username": "DenilsonShow",
}
```

- `Response`
```
{
  "room": {
    "code": "TJKLAL"
  },
  "scoreboard": [
    {
      "user": {
        "name": "RenataFan",
        "state": "READY"
      },
      "score": 120
    },
    {
      "user": {
        "name": "CraqueNeto",
        "state": "PENDING"
      },
      "score": 230
    },
    {
      "user": {
        "name": "RonaldoCareca",
        "state": "WAITING"
      },
      "score": 0
    }
  ]
}
```


- `POST` - `/rooms/:room_code/hunches`
```
Headers: {
    "FCM_USER_TOKEN": "dkQJJ-3BLbg:APA91bFtqfdcku1WMPX2CkuCtJO9EomUucrh-aFs3X3mMJj636MPR7jbkRY"
}
```
```
{
    "hunch": 31000
}
```

- `PUT` - `/rooms/:room_code/players/me`
```
Headers: {
    "FCM_USER_TOKEN": "dkQJJ-3BLbg:APA91bFtqfdcku1WMPX2CkuCtJO9EomUucrh-aFs3X3mMJj636MPR7jbkRY"
}
```
```
{
    "state": "READY"
}
```


## Firebase Cloud Messaging

### Messages

- `ROOM_STATE`
```
{

	"code": "TJKLAL",
	"round": {
		"current": 1,
		"max": 10,
		"question": {
			"name": "Flávio",
		    "answer": 34300,
		},
		"winner": {
			"user": {
				"name": "CraqueNeto"
			},
			"hunch": 31000
		},
        "ends_at": 1605392362
	}
}
```

- `SCOREBOARD`
```
{
  "scoreboard": [
    {
      "user": {
        "name": "RenataFan",
        "state": "READY"
      },
      "score": 200
    },
    {
      "user": {
        "name": "CraqueNeto",
        "state": "PENDING"
      },
      "score": 190
    }
  ]
}
```