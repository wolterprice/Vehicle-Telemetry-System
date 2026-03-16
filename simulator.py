import requests
import random
import time

API = "http://localhost:8080/telemetry"

while True:
    data = {
        "speed": random.uniform(40,120),
        "rpm": random.uniform(1500,5000),
        "temperature": random.uniform(80,100),
        "acceleration": random.uniform(-2,2)
    }

    requests.post(API,json=data)

    print("sent",data)

    time.sleep(1)