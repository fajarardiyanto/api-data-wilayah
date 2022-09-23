-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS data_wilayah_regencies(
     id bigint not null primary key ,
     district_id bigint not null ,
     name varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS data_wilayah_regencies;
-- +goose StatementEnd