CREATE TABLE group_cost_item_participants (
    id SERIAL PRIMARY KEY ,
    group_cost_item_id int NOT NULL ,
    user_id int NOT NULL ,
    amount int NOT NULL ,
    FOREIGN KEY (group_cost_item_id) REFERENCES group_cost_items (id) ,
    FOREIGN KEY (user_id) REFERENCES users (id)
)