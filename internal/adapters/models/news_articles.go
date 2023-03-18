// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// NewsArticle is an object representing the database table.
type NewsArticle struct {
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title        string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Slug         string      `boil:"slug" json:"slug" toml:"slug" yaml:"slug"`
	Content      null.String `boil:"content" json:"content,omitempty" toml:"content" yaml:"content,omitempty"`
	Status       null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	PublishedAt  null.Time   `boil:"published_at" json:"published_at,omitempty" toml:"published_at" yaml:"published_at,omitempty"`
	LastEditedAt null.Time   `boil:"last_edited_at" json:"last_edited_at,omitempty" toml:"last_edited_at" yaml:"last_edited_at,omitempty"`

	R *newsArticleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L newsArticleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NewsArticleColumns = struct {
	ID           string
	Title        string
	Slug         string
	Content      string
	Status       string
	PublishedAt  string
	LastEditedAt string
}{
	ID:           "id",
	Title:        "title",
	Slug:         "slug",
	Content:      "content",
	Status:       "status",
	PublishedAt:  "published_at",
	LastEditedAt: "last_edited_at",
}

var NewsArticleTableColumns = struct {
	ID           string
	Title        string
	Slug         string
	Content      string
	Status       string
	PublishedAt  string
	LastEditedAt string
}{
	ID:           "news_articles.id",
	Title:        "news_articles.title",
	Slug:         "news_articles.slug",
	Content:      "news_articles.content",
	Status:       "news_articles.status",
	PublishedAt:  "news_articles.published_at",
	LastEditedAt: "news_articles.last_edited_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var NewsArticleWhere = struct {
	ID           whereHelperstring
	Title        whereHelperstring
	Slug         whereHelperstring
	Content      whereHelpernull_String
	Status       whereHelpernull_String
	PublishedAt  whereHelpernull_Time
	LastEditedAt whereHelpernull_Time
}{
	ID:           whereHelperstring{field: "\"news_articles\".\"id\""},
	Title:        whereHelperstring{field: "\"news_articles\".\"title\""},
	Slug:         whereHelperstring{field: "\"news_articles\".\"slug\""},
	Content:      whereHelpernull_String{field: "\"news_articles\".\"content\""},
	Status:       whereHelpernull_String{field: "\"news_articles\".\"status\""},
	PublishedAt:  whereHelpernull_Time{field: "\"news_articles\".\"published_at\""},
	LastEditedAt: whereHelpernull_Time{field: "\"news_articles\".\"last_edited_at\""},
}

// NewsArticleRels is where relationship names are stored.
var NewsArticleRels = struct {
	NewsArticlesTopics string
}{
	NewsArticlesTopics: "NewsArticlesTopics",
}

// newsArticleR is where relationships are stored.
type newsArticleR struct {
	NewsArticlesTopics NewsArticlesTopicSlice `boil:"NewsArticlesTopics" json:"NewsArticlesTopics" toml:"NewsArticlesTopics" yaml:"NewsArticlesTopics"`
}

// NewStruct creates a new relationship struct
func (*newsArticleR) NewStruct() *newsArticleR {
	return &newsArticleR{}
}

func (r *newsArticleR) GetNewsArticlesTopics() NewsArticlesTopicSlice {
	if r == nil {
		return nil
	}
	return r.NewsArticlesTopics
}

// newsArticleL is where Load methods for each relationship are stored.
type newsArticleL struct{}

var (
	newsArticleAllColumns            = []string{"id", "title", "slug", "content", "status", "published_at", "last_edited_at"}
	newsArticleColumnsWithoutDefault = []string{"title", "slug"}
	newsArticleColumnsWithDefault    = []string{"id", "content", "status", "published_at", "last_edited_at"}
	newsArticlePrimaryKeyColumns     = []string{"id"}
	newsArticleGeneratedColumns      = []string{}
)

type (
	// NewsArticleSlice is an alias for a slice of pointers to NewsArticle.
	// This should almost always be used instead of []NewsArticle.
	NewsArticleSlice []*NewsArticle
	// NewsArticleHook is the signature for custom NewsArticle hook methods
	NewsArticleHook func(context.Context, boil.ContextExecutor, *NewsArticle) error

	newsArticleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	newsArticleType                 = reflect.TypeOf(&NewsArticle{})
	newsArticleMapping              = queries.MakeStructMapping(newsArticleType)
	newsArticlePrimaryKeyMapping, _ = queries.BindMapping(newsArticleType, newsArticleMapping, newsArticlePrimaryKeyColumns)
	newsArticleInsertCacheMut       sync.RWMutex
	newsArticleInsertCache          = make(map[string]insertCache)
	newsArticleUpdateCacheMut       sync.RWMutex
	newsArticleUpdateCache          = make(map[string]updateCache)
	newsArticleUpsertCacheMut       sync.RWMutex
	newsArticleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var newsArticleAfterSelectHooks []NewsArticleHook

var newsArticleBeforeInsertHooks []NewsArticleHook
var newsArticleAfterInsertHooks []NewsArticleHook

var newsArticleBeforeUpdateHooks []NewsArticleHook
var newsArticleAfterUpdateHooks []NewsArticleHook

var newsArticleBeforeDeleteHooks []NewsArticleHook
var newsArticleAfterDeleteHooks []NewsArticleHook

var newsArticleBeforeUpsertHooks []NewsArticleHook
var newsArticleAfterUpsertHooks []NewsArticleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NewsArticle) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsArticleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NewsArticle) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsArticleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NewsArticle) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsArticleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NewsArticle) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsArticleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NewsArticle) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsArticleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NewsArticle) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsArticleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NewsArticle) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsArticleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NewsArticle) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsArticleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NewsArticle) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range newsArticleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNewsArticleHook registers your hook function for all future operations.
func AddNewsArticleHook(hookPoint boil.HookPoint, newsArticleHook NewsArticleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		newsArticleAfterSelectHooks = append(newsArticleAfterSelectHooks, newsArticleHook)
	case boil.BeforeInsertHook:
		newsArticleBeforeInsertHooks = append(newsArticleBeforeInsertHooks, newsArticleHook)
	case boil.AfterInsertHook:
		newsArticleAfterInsertHooks = append(newsArticleAfterInsertHooks, newsArticleHook)
	case boil.BeforeUpdateHook:
		newsArticleBeforeUpdateHooks = append(newsArticleBeforeUpdateHooks, newsArticleHook)
	case boil.AfterUpdateHook:
		newsArticleAfterUpdateHooks = append(newsArticleAfterUpdateHooks, newsArticleHook)
	case boil.BeforeDeleteHook:
		newsArticleBeforeDeleteHooks = append(newsArticleBeforeDeleteHooks, newsArticleHook)
	case boil.AfterDeleteHook:
		newsArticleAfterDeleteHooks = append(newsArticleAfterDeleteHooks, newsArticleHook)
	case boil.BeforeUpsertHook:
		newsArticleBeforeUpsertHooks = append(newsArticleBeforeUpsertHooks, newsArticleHook)
	case boil.AfterUpsertHook:
		newsArticleAfterUpsertHooks = append(newsArticleAfterUpsertHooks, newsArticleHook)
	}
}

// One returns a single newsArticle record from the query.
func (q newsArticleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NewsArticle, error) {
	o := &NewsArticle{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for news_articles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all NewsArticle records from the query.
func (q newsArticleQuery) All(ctx context.Context, exec boil.ContextExecutor) (NewsArticleSlice, error) {
	var o []*NewsArticle

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NewsArticle slice")
	}

	if len(newsArticleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all NewsArticle records in the query.
func (q newsArticleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count news_articles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q newsArticleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if news_articles exists")
	}

	return count > 0, nil
}

// NewsArticlesTopics retrieves all the news_articles_topic's NewsArticlesTopics with an executor.
func (o *NewsArticle) NewsArticlesTopics(mods ...qm.QueryMod) newsArticlesTopicQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"news_articles_topics\".\"news_articles_id\"=?", o.ID),
	)

	return NewsArticlesTopics(queryMods...)
}

// LoadNewsArticlesTopics allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (newsArticleL) LoadNewsArticlesTopics(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNewsArticle interface{}, mods queries.Applicator) error {
	var slice []*NewsArticle
	var object *NewsArticle

	if singular {
		var ok bool
		object, ok = maybeNewsArticle.(*NewsArticle)
		if !ok {
			object = new(NewsArticle)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeNewsArticle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeNewsArticle))
			}
		}
	} else {
		s, ok := maybeNewsArticle.(*[]*NewsArticle)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeNewsArticle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeNewsArticle))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &newsArticleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &newsArticleR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`news_articles_topics`),
		qm.WhereIn(`news_articles_topics.news_articles_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load news_articles_topics")
	}

	var resultSlice []*NewsArticlesTopic
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice news_articles_topics")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on news_articles_topics")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for news_articles_topics")
	}

	if len(newsArticlesTopicAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.NewsArticlesTopics = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &newsArticlesTopicR{}
			}
			foreign.R.NewsArticle = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.NewsArticlesID {
				local.R.NewsArticlesTopics = append(local.R.NewsArticlesTopics, foreign)
				if foreign.R == nil {
					foreign.R = &newsArticlesTopicR{}
				}
				foreign.R.NewsArticle = local
				break
			}
		}
	}

	return nil
}

// AddNewsArticlesTopics adds the given related objects to the existing relationships
// of the news_article, optionally inserting them as new records.
// Appends related to o.R.NewsArticlesTopics.
// Sets related.R.NewsArticle appropriately.
func (o *NewsArticle) AddNewsArticlesTopics(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*NewsArticlesTopic) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.NewsArticlesID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"news_articles_topics\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"news_articles_id"}),
				strmangle.WhereClause("\"", "\"", 2, newsArticlesTopicPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.NewsArticlesID = o.ID
		}
	}

	if o.R == nil {
		o.R = &newsArticleR{
			NewsArticlesTopics: related,
		}
	} else {
		o.R.NewsArticlesTopics = append(o.R.NewsArticlesTopics, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &newsArticlesTopicR{
				NewsArticle: o,
			}
		} else {
			rel.R.NewsArticle = o
		}
	}
	return nil
}

// NewsArticles retrieves all the records using an executor.
func NewsArticles(mods ...qm.QueryMod) newsArticleQuery {
	mods = append(mods, qm.From("\"news_articles\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"news_articles\".*"})
	}

	return newsArticleQuery{q}
}

// FindNewsArticle retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNewsArticle(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*NewsArticle, error) {
	newsArticleObj := &NewsArticle{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"news_articles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, newsArticleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from news_articles")
	}

	if err = newsArticleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return newsArticleObj, err
	}

	return newsArticleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NewsArticle) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no news_articles provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(newsArticleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	newsArticleInsertCacheMut.RLock()
	cache, cached := newsArticleInsertCache[key]
	newsArticleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			newsArticleAllColumns,
			newsArticleColumnsWithDefault,
			newsArticleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(newsArticleType, newsArticleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(newsArticleType, newsArticleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"news_articles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"news_articles\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into news_articles")
	}

	if !cached {
		newsArticleInsertCacheMut.Lock()
		newsArticleInsertCache[key] = cache
		newsArticleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the NewsArticle.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NewsArticle) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	newsArticleUpdateCacheMut.RLock()
	cache, cached := newsArticleUpdateCache[key]
	newsArticleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			newsArticleAllColumns,
			newsArticlePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update news_articles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"news_articles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, newsArticlePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(newsArticleType, newsArticleMapping, append(wl, newsArticlePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update news_articles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for news_articles")
	}

	if !cached {
		newsArticleUpdateCacheMut.Lock()
		newsArticleUpdateCache[key] = cache
		newsArticleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q newsArticleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for news_articles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for news_articles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NewsArticleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), newsArticlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"news_articles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, newsArticlePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in newsArticle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all newsArticle")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NewsArticle) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no news_articles provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(newsArticleColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	newsArticleUpsertCacheMut.RLock()
	cache, cached := newsArticleUpsertCache[key]
	newsArticleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			newsArticleAllColumns,
			newsArticleColumnsWithDefault,
			newsArticleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			newsArticleAllColumns,
			newsArticlePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert news_articles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(newsArticlePrimaryKeyColumns))
			copy(conflict, newsArticlePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"news_articles\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(newsArticleType, newsArticleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(newsArticleType, newsArticleMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert news_articles")
	}

	if !cached {
		newsArticleUpsertCacheMut.Lock()
		newsArticleUpsertCache[key] = cache
		newsArticleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single NewsArticle record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NewsArticle) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NewsArticle provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), newsArticlePrimaryKeyMapping)
	sql := "DELETE FROM \"news_articles\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from news_articles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for news_articles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q newsArticleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no newsArticleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from news_articles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for news_articles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NewsArticleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(newsArticleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), newsArticlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"news_articles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, newsArticlePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from newsArticle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for news_articles")
	}

	if len(newsArticleAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *NewsArticle) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNewsArticle(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NewsArticleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NewsArticleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), newsArticlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"news_articles\".* FROM \"news_articles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, newsArticlePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NewsArticleSlice")
	}

	*o = slice

	return nil
}

// NewsArticleExists checks if the NewsArticle row exists.
func NewsArticleExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"news_articles\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if news_articles exists")
	}

	return exists, nil
}

// Exists checks if the NewsArticle row exists.
func (o *NewsArticle) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return NewsArticleExists(ctx, exec, o.ID)
}
