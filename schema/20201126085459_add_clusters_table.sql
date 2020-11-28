-- +goose Up
-- +goose StatementBegin
CREATE TABLE clusters (
	id serial NOT NULL PRIMARY KEY,
	name varchar(255) NOT NULL,
	namespace varchar(255) NOT NULL,
	data jsonb NOT NULL,
	unique (name, namespace)
);
CREATE TABLE syncsets (
	id serial NOT NULL PRIMARY KEY,
	name varchar(255) NOT NULL,
	namespace varchar(255) NOT NULL,
	data jsonb NOT NULL,
	unique (name, namespace)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clusters;
DROP TABLE syncsets;
-- +goose StatementEnd
