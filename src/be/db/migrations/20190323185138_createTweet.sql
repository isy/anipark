
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tweets (
  id BIGINT NOT NULL AUTO_INCREMENT,
  text TEXT NOT NULL,
  PRIMARY KEY(id),
)

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tweets;

