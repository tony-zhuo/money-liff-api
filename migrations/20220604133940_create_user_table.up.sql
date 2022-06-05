CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    line_id VARCHAR (40) UNIQUE NOT NULL ,
    name VARCHAR (20) NOT NULL ,
    avatar_url TEXT DEFAULT NULL ,
    created_at timestamp without time zone DEFAULT NULL ,
    updated_at timestamp without time zone DEFAULT NULL ,
    deleted_at timestamp without time zone DEFAULT NULL
)