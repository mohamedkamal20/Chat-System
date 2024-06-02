CREATE KEYSPACE IF NOT EXISTS chat_system WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE chat_system;

CREATE TABLE IF NOT EXISTS users (
        user_id UUID ,
        email TEXT PRIMARY KEY,
        password TEXT,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
);

CREATE TABLE IF NOT EXISTS messages (
        message_id UUID PRIMARY KEY,
        sender Text,
        recipient Text,
        content TEXT,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
);