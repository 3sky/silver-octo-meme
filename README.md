# Envy proxy with ZeroMQ

Solution tested on DigitalOcean Droplet. In case of re-creation VM follow this [article](https://blog.3sky.dev/article/cloud-init-intro/).

## Introduction

Let's implement some basic [ZeroMQ](https://zeromq.org/) publisher and subscriber in Golang.
Utilize [Envoy](https://www.envoyproxy.io) as a proxy. 

## Implmentation

![arch](https://excalidraw.com/#json=-24QxmqIttlH6wb4_7Oqt,BjLW7_-WRsN8ZRBM-DApFQ)

## Demo

1. Run environment(droplet)

    ```bash
    docker-compose up
    ```
1. Open another terminal's windows(workstation)

1. Grab droplet's IP(workstation)
   
   ```bash
   export IP=$(tf output -raw public_ip)
   ```

1. Watch logs(workstation)

1. Run CURL command(workstation)

    ```bash
    curl -i --request POST   --url http://"$IP":8080/   --header 'Content-Type: application/json'   --data '{
        "title": "test",
        "content": "this is test message"
    }' 
    ```
