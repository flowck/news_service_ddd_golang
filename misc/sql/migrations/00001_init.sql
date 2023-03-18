-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE NEWS_ARTICLES_STATUS as ENUM ('draft', 'published', 'scheduled');

CREATE TABLE news_articles(
    id UUID DEFAULT gen_random_uuid(),
    title VARCHAR(250) NOT NULL,
    slug VARCHAR(250) NOT NULL,
    content TEXT,
    status NEWS_ARTICLES_STATUS default 'draft',
    published_at TIMESTAMPTZ default now(),
    last_edited_at TIMESTAMPTZ default now(),

    CONSTRAINT pk_news_articles_id PRIMARY KEY(id)
);

CREATE TABLE topics(
    id UUID DEFAULT gen_random_uuid(),
    name VARCHAR(250) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),

    CONSTRAINT pk_topics_id PRIMARY KEY(id)
);

CREATE TABLE news_articles_topics(
    topic_id UUID NOT NULL,
    news_article_id UUID NOT NULL,

    CONSTRAINT pk_news_topics_id PRIMARY KEY(news_article_id, topic_id),

    CONSTRAINT pk_news_topics_news_id FOREIGN KEY (news_article_id)
        REFERENCES news_articles(id)
        ON DELETE CASCADE,

    CONSTRAINT pk_news_topics_topic_id FOREIGN KEY (topic_id)
        REFERENCES topics(id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE news_articles_topics;
DROP TABLE news_articles;
DROP TABLE topics;

DROP EXTENSION IF EXISTS "uuid-ossp";
DROP TYPE NEWS_ARTICLES_STATUS;
-- +goose StatementEnd
