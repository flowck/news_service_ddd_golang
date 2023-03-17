-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE NEWS_STATUS as ENUM ('draft', 'published', 'scheduled');

CREATE TABLE news(
    id UUID DEFAULT gen_random_uuid(),
    title VARCHAR(250),
    slug VARCHAR(250),
    content TEXT,
    status VARCHAR(4),
    publishedAt TIMESTAMPTZ default now(),
    lastEditedAt TIMESTAMPTZ default now(),

    CONSTRAINT pk_news_id PRIMARY KEY(id)
);

CREATE TABLE topics(
    id UUID DEFAULT gen_random_uuid(),
    name VARCHAR(250) NOT NULL,
    createdAt TIMESTAMPTZ DEFAULT now(),
    updatedAt TIMESTAMPTZ DEFAULT now(),

    CONSTRAINT pk_topics_id PRIMARY KEY(id)
);

CREATE TABLE news_topics(
    id SERIAL,
    news_id UUID NOT NULL,
    topic_id UUID NOT NULL,

    CONSTRAINT pk_news_topics_id PRIMARY KEY(id),
    CONSTRAINT pk_news_topics_news_id FOREIGN KEY (news_id) REFERENCES news(id)
        ON DELETE CASCADE,
    CONSTRAINT pk_news_topics_topic_id FOREIGN KEY (topic_id) REFERENCES topics(id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE news_topics;
DROP TABLE news;
DROP TABLE topics;

DROP EXTENSION IF EXISTS "uuid-ossp";
DROP TYPE NEWS_STATUS;
-- +goose StatementEnd
