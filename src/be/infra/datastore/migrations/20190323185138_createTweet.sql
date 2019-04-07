
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tweets (
  id SERIAL,
  text TEXT NOT NULL,
  created_at timestamp not null default current_timestamp,
  update_at timestamp not null default current_timestamp on update current_timestamp,
  CONSTRAINT pk_tweets PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tweets;

