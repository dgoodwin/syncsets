-- +goose Up
-- +goose StatementBegin
CREATE TABLE clusters (
	id serial NOT NULL PRIMARY KEY,
	data jsonb NOT NULL
);
CREATE TABLE syncsets (
	id serial NOT NULL PRIMARY KEY,
	data jsonb NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clusters;
DROP TABLE syncsets;
-- +goose StatementEnd
