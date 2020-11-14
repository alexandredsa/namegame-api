# NAMEGAME-API

## About


## Socket

### Events

### Consume

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
    [
        {
            "user": {
                "name": "RenataFan"
                "state": "READY"
            },
            "score": 200
        },
        {
            "user": {
                "name": "CraqueNeto",
                "state: "PENDING"
            },
            "score": 190
        },
    ]
}
```

### Produce

- `ROOM_CREATE`
```
{
    "id": "6dcb9970-53ca-493f-876a-61dca6d4de46",  
    "username": "EdilsonCapetinha"
}
```

- `ROOM_JOIN`
```
{
    "username": "DenilsonShow",
    "room_code": "TJAKLL"
}
```

- `HUNCH_CREATE`
```
{
    "hunch": 31000
}
```

- `PLAYER_STATE_UPDATE`
```
{
    "state": "READY"
}
```


### Flow example: 

- Emit -> `ROOM_CREATE`
- `ROOM_STATE` <- Receive
- Emit -> `PLAYER_STATE_UPDATE`
- `SCOREBOARD` <- Receive - [_Multiple times_]
- `ROOM_STATE` <- Receive
- Emit -> `HUNCH_CREATE`
- `ROOM_STATE` <- Receive
- `SCOREBOARD` <- Receive