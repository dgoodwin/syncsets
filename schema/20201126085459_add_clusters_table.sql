-- +goose Up
-- +goose StatementBegin
CREATE TABLE clusters (
	id serial NOT NULL PRIMARY KEY,
	data jsonb NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clusters;
-- +goose StatementEnd
