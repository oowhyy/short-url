-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shorturl
(
 short text NOT NULL,
 long  text NOT NULL,
 CONSTRAINT PK_1 PRIMARY KEY ( short, long )
);

CREATE INDEX Index_1 ON shorturl
(
 short
);

CREATE INDEX Index_2 ON shorturl
(
 long
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shorturl;
-- +goose StatementEnd

