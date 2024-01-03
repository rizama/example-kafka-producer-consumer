import { Kafka, Partitioners } from 'kafkajs';

const kafka = new Kafka({
    brokers: ['localhost:29092'],
    clientId: 'sam-producer',
});

const producer = kafka.producer({
    createPartitioner: Partitioners.DefaultPartitioner,
});

await producer.connect();

for (let index = 0; index < 10; index++) {
    await producer.send({
        topic: 'helloworld',
        messages: [
            {
                key: `${index}`,
                value: `Hellow ${index}`,
            },
        ],
    });
}

await producer.disconnect();
