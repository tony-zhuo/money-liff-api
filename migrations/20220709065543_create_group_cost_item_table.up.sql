CREATE TABLE group_cost_items (
    id SERIAL PRIMARY KEY ,
    group_id int NOT NULL ,
    uuid uuid NOT NULL ,
    name varchar(20) NOT NULL ,
    total_amount int NOT NULL ,
    payer_id int NOT NULL ,
    creator_id int NOT NULL ,
    pay_at timestamp without time zone ,
    remark text ,
    created_at timestamp without time zone DEFAULT NULL ,
    updated_at timestamp without time zone DEFAULT NULL ,
    deleted_at timestamp without time zone DEFAULT NULL ,

    FOREIGN KEY (group_id) REFERENCES groups (id) ,
    FOREIGN KEY (payer_id) REFERENCES users (id) ,
    FOREIGN KEY (creator_id) REFERENCES users (id)
)