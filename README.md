# assignment_demo_2023

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

This is a demo and template for backend assignment of 2023 TikTok Tech Immersion.

## Installation

Requirement:

- golang 1.18+
- docker

To install dependency tools:

```bash
make pre
```

## Run

```bash
docker-compose up -d
```

Check if it's running:

```bash
curl localhost:8080/ping
```

## API Documentation

### Get messages

#### Endpoint:  GET **{server}/api/pull**

#### Request payload parameters

| Parameter  | Description |  Extra Details |
| ----------- | ----------- | ----------- |
| chat      | format "<member1>:<member2>", e.g. "john:doe"  | Same results if you reverse the order of names |
| cursor  | Int type, denotes the starting position of the message's send time, inclusively, 0 by default  ||
| limit | int type, maximum number of messages returned per request, 10 by default || 
| reverse | bool type, if false, results will be sorted in ascending order by time ||

#### Example payload
```
{
    "chat": "d:a",
    "cursor": 0,
    "limit": 5,
    "reverse": false
}
```

#### Response data
| Parameter  | Description |  Extra Details |
| ----------- | ----------- | ----------- |
| messages   | Array of messages between two users ||
| has_more | boolean type, indicates if there are more messages || 
| next_cursor | int type, indicates the starting point of the next messages if has_more is true ||

#### Sample Response
```
{
    "messages": [
        {
            "chat": "d:a",
            "text": "abcd",
            "sender": "a",
            "send_time": 1686579323151951513
        },
        {
            "chat": "d:a",
            "text": "abcd",
            "sender": "a",
            "send_time": 1686579906117379339
        }
    ],
    "has_more": true,
    "next_cursor": 2
}
```


### Send messages

#### Endpoint:  POST **{server}/api/send**

#### Request payload parameters

| Parameter  | Description |  Extra Details |
| ----------- | ----------- | ----------- |
| sender      | Sender name/identifier | |
| text | String type, content of message || 
| chat | String type, format "<member1>:<member2>", e.g. "john:doe"  ||

#### Example payload
```
{

    "sender": "a",
    "text": "abcde",
    "chat": "a:d"
}
```

#### Response is empty, 200 status code if successful, 500 otherwise

