import os, json, time, pika

AMQP_URL = os.getenv("AMQP_URL", "amqp://guest:guest@rabbitmq:5672/")
QUEUE    = os.getenv("AMQP_QUEUE", "southpark_messages")

def connect():
    params = pika.URLParameters(AMQP_URL)
    return pika.BlockingConnection(params)

def start():
    while True:
        try:
            print(f"[consumer] connecting to {AMQP_URL} …", flush=True)
            conn = connect()
            ch = conn.channel()
            ch.queue_declare(queue=QUEUE, durable=True)
            print(f"[consumer] listening on queue: {QUEUE}", flush=True)

            def on_msg(ch_, method, props, body):
                try:
                    data = json.loads(body.decode("utf-8"))
                    author = data.get("author", "Unknown")
                    text   = data.get("body", "")
                    print(f"{author} says: {text}", flush=True)
                    ch_.basic_ack(delivery_tag=method.delivery_tag)
                except Exception as e:
                    print(f"[consumer] error parsing message: {e}", flush=True)
                    ch_.basic_nack(delivery_tag=method.delivery_tag, requeue=False)

            ch.basic_qos(prefetch_count=10)
            ch.basic_consume(queue=QUEUE, on_message_callback=on_msg, auto_ack=False)
            ch.start_consuming()
        except KeyboardInterrupt:
            break
        except Exception as e:
            print(f"[consumer] connection error: {e} — retrying in 3s", flush=True)
            time.sleep(3)

if __name__ == "__main__":
    start()
