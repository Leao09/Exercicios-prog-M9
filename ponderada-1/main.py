import paho.mqtt.client as mqtt
import time
from faker import Faker
import json

def connect_client(broker_address="localhost", broker_port=9341):
    client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
    client.connect(broker_address, broker_port, 60)
    return client

def generate_sensor_data():
    fake = Faker()
    sensor_outputs = {
        "CO_ppm": fake.pyfloat(min_value=1, max_value=1000, right_digits=2),
        "NO2_ppm": fake.pyfloat(min_value=0.05, max_value=10, right_digits=2),
        "NH3_ppm": fake.pyfloat(min_value=1, max_value=300, right_digits=2),
    }
    return sensor_outputs

def publish_sensor_data(client, topic, sensor_data):
    payload = json.dumps(sensor_data)
    client.publish(topic, payload)
    print(f"[SENSOR] Publishing message: {payload}")

def main():
    broker_address = "localhost"
    broker_port = 9341
    topic = "sensor/mics6814"
    sleep_time = 1

    client = connect_client(broker_address, broker_port)
    
    try:
        while True:
            sensor_data = generate_sensor_data()
            publish_sensor_data(client, topic, sensor_data)
            time.sleep(sleep_time)
    except KeyboardInterrupt:
        print("[SENSOR] Publisher interrupted")
    except Exception as err:
        print(f"[SENSOR] An unexpected exception occurred: {str(err)}")
    finally:
        client.disconnect()

if __name__ == "__main__":
    main()