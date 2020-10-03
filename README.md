# healthcheck-cmd

A tiny utility tool that runs a command when an HTTP healthcheck call is made.

## Usage

```
./healthcheck-cmd -command 'docker inspect my-container --format "{{json .State.Status }}" | grep "running"' 

curl -v http://localhost:4000 # returns 200 if my-container is running, 500 if not.
```

| Variable                         | Description                                                                                                                                                                                                                                          |
|----------------------------------|---------------------------------------------------------------------------------------|
| `-command`                       | ***Required*** Command to run when healthcheck is called.                             |
| `-listen`                        | Listen on port, default `:4000`                                                       |
| `-path`                          | Listen on path, deafult `/`                                                           |
| `-success`                       | Returned HTTP status code if command succeeds. Default `200`.                         |
| `-error`                         | Returned HTTP status code if command fails. Default `500`.                            |


