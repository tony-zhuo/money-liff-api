CREATE TABLE groups (
    id SERIAL PRIMARY KEY ,
    uuid uuid NOT NULL ,
    name VARCHAR (50) NOT NULL ,
    user_limit int NOT NULL ,
    image_url TEXT DEFAULT NULL ,
    admin_user_id int NOT NULL ,
    created_at timestamp without time zone DEFAULT NULL ,
    updated_at timestamp without time zone DEFAULT NULL ,
    deleted_at timestamp without time zone DEFAULT NULL ,
    FOREIGN KEY (admin_user_id) REFERENCES users (id)
)