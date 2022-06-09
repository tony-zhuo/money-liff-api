CREATE TABLE user_groups (
    id SERIAL PRIMARY KEY ,
    group_id int NOT NULL ,
    user_id int NOT NULL ,
    FOREIGN KEY (group_id) REFERENCES groups (id) ,
    FOREIGN KEY (user_id) REFERENCES users (id)
)