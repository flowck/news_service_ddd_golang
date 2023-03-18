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

// Topic is an object representing the database table.
type Topic struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *topicR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L topicL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TopicColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var TopicTableColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "topics.id",
	Name:      "topics.name",
	CreatedAt: "topics.created_at",
	UpdatedAt: "topics.updated_at",
}

// Generated where

var TopicWhere = struct {
	ID        whereHelperstring
	Name      whereHelperstring
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"topics\".\"id\""},
	Name:      whereHelperstring{field: "\"topics\".\"name\""},
	CreatedAt: whereHelpernull_Time{field: "\"topics\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"topics\".\"updated_at\""},
}

// TopicRels is where relationship names are stored.
var TopicRels = struct {
	NewsArticlesTopics string
}{
	NewsArticlesTopics: "NewsArticlesTopics",
}

// topicR is where relationships are stored.
type topicR struct {
	NewsArticlesTopics NewsArticlesTopicSlice `boil:"NewsArticlesTopics" json:"NewsArticlesTopics" toml:"NewsArticlesTopics" yaml:"NewsArticlesTopics"`
}

// NewStruct creates a new relationship struct
func (*topicR) NewStruct() *topicR {
	return &topicR{}
}

func (r *topicR) GetNewsArticlesTopics() NewsArticlesTopicSlice {
	if r == nil {
		return nil
	}
	return r.NewsArticlesTopics
}

// topicL is where Load methods for each relationship are stored.
type topicL struct{}

var (
	topicAllColumns            = []string{"id", "name", "created_at", "updated_at"}
	topicColumnsWithoutDefault = []string{"name"}
	topicColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	topicPrimaryKeyColumns     = []string{"id"}
	topicGeneratedColumns      = []string{}
)

type (
	// TopicSlice is an alias for a slice of pointers to Topic.
	// This should almost always be used instead of []Topic.
	TopicSlice []*Topic
	// TopicHook is the signature for custom Topic hook methods
	TopicHook func(context.Context, boil.ContextExecutor, *Topic) error

	topicQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	topicType                 = reflect.TypeOf(&Topic{})
	topicMapping              = queries.MakeStructMapping(topicType)
	topicPrimaryKeyMapping, _ = queries.BindMapping(topicType, topicMapping, topicPrimaryKeyColumns)
	topicInsertCacheMut       sync.RWMutex
	topicInsertCache          = make(map[string]insertCache)
	topicUpdateCacheMut       sync.RWMutex
	topicUpdateCache          = make(map[string]updateCache)
	topicUpsertCacheMut       sync.RWMutex
	topicUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var topicAfterSelectHooks []TopicHook

var topicBeforeInsertHooks []TopicHook
var topicAfterInsertHooks []TopicHook

var topicBeforeUpdateHooks []TopicHook
var topicAfterUpdateHooks []TopicHook

var topicBeforeDeleteHooks []TopicHook
var topicAfterDeleteHooks []TopicHook

var topicBeforeUpsertHooks []TopicHook
var topicAfterUpsertHooks []TopicHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Topic) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topicAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Topic) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topicBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Topic) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topicAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Topic) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topicBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Topic) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topicAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Topic) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topicBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Topic) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topicAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Topic) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topicBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Topic) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topicAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTopicHook registers your hook function for all future operations.
func AddTopicHook(hookPoint boil.HookPoint, topicHook TopicHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		topicAfterSelectHooks = append(topicAfterSelectHooks, topicHook)
	case boil.BeforeInsertHook:
		topicBeforeInsertHooks = append(topicBeforeInsertHooks, topicHook)
	case boil.AfterInsertHook:
		topicAfterInsertHooks = append(topicAfterInsertHooks, topicHook)
	case boil.BeforeUpdateHook:
		topicBeforeUpdateHooks = append(topicBeforeUpdateHooks, topicHook)
	case boil.AfterUpdateHook:
		topicAfterUpdateHooks = append(topicAfterUpdateHooks, topicHook)
	case boil.BeforeDeleteHook:
		topicBeforeDeleteHooks = append(topicBeforeDeleteHooks, topicHook)
	case boil.AfterDeleteHook:
		topicAfterDeleteHooks = append(topicAfterDeleteHooks, topicHook)
	case boil.BeforeUpsertHook:
		topicBeforeUpsertHooks = append(topicBeforeUpsertHooks, topicHook)
	case boil.AfterUpsertHook:
		topicAfterUpsertHooks = append(topicAfterUpsertHooks, topicHook)
	}
}

// One returns a single topic record from the query.
func (q topicQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Topic, error) {
	o := &Topic{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for topics")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Topic records from the query.
func (q topicQuery) All(ctx context.Context, exec boil.ContextExecutor) (TopicSlice, error) {
	var o []*Topic

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Topic slice")
	}

	if len(topicAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Topic records in the query.
func (q topicQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count topics rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q topicQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if topics exists")
	}

	return count > 0, nil
}

// NewsArticlesTopics retrieves all the news_articles_topic's NewsArticlesTopics with an executor.
func (o *Topic) NewsArticlesTopics(mods ...qm.QueryMod) newsArticlesTopicQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"news_articles_topics\".\"topic_id\"=?", o.ID),
	)

	return NewsArticlesTopics(queryMods...)
}

// LoadNewsArticlesTopics allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (topicL) LoadNewsArticlesTopics(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTopic interface{}, mods queries.Applicator) error {
	var slice []*Topic
	var object *Topic

	if singular {
		var ok bool
		object, ok = maybeTopic.(*Topic)
		if !ok {
			object = new(Topic)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTopic)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTopic))
			}
		}
	} else {
		s, ok := maybeTopic.(*[]*Topic)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTopic)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTopic))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &topicR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &topicR{}
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
		qm.WhereIn(`news_articles_topics.topic_id in ?`, args...),
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
			foreign.R.Topic = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TopicID {
				local.R.NewsArticlesTopics = append(local.R.NewsArticlesTopics, foreign)
				if foreign.R == nil {
					foreign.R = &newsArticlesTopicR{}
				}
				foreign.R.Topic = local
				break
			}
		}
	}

	return nil
}

// AddNewsArticlesTopics adds the given related objects to the existing relationships
// of the topic, optionally inserting them as new records.
// Appends related to o.R.NewsArticlesTopics.
// Sets related.R.Topic appropriately.
func (o *Topic) AddNewsArticlesTopics(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*NewsArticlesTopic) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TopicID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"news_articles_topics\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"topic_id"}),
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

			rel.TopicID = o.ID
		}
	}

	if o.R == nil {
		o.R = &topicR{
			NewsArticlesTopics: related,
		}
	} else {
		o.R.NewsArticlesTopics = append(o.R.NewsArticlesTopics, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &newsArticlesTopicR{
				Topic: o,
			}
		} else {
			rel.R.Topic = o
		}
	}
	return nil
}

// Topics retrieves all the records using an executor.
func Topics(mods ...qm.QueryMod) topicQuery {
	mods = append(mods, qm.From("\"topics\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"topics\".*"})
	}

	return topicQuery{q}
}

// FindTopic retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTopic(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Topic, error) {
	topicObj := &Topic{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"topics\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, topicObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from topics")
	}

	if err = topicObj.doAfterSelectHooks(ctx, exec); err != nil {
		return topicObj, err
	}

	return topicObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Topic) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no topics provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(topicColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	topicInsertCacheMut.RLock()
	cache, cached := topicInsertCache[key]
	topicInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			topicAllColumns,
			topicColumnsWithDefault,
			topicColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(topicType, topicMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(topicType, topicMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"topics\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"topics\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into topics")
	}

	if !cached {
		topicInsertCacheMut.Lock()
		topicInsertCache[key] = cache
		topicInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Topic.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Topic) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	topicUpdateCacheMut.RLock()
	cache, cached := topicUpdateCache[key]
	topicUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			topicAllColumns,
			topicPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update topics, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"topics\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, topicPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(topicType, topicMapping, append(wl, topicPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update topics row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for topics")
	}

	if !cached {
		topicUpdateCacheMut.Lock()
		topicUpdateCache[key] = cache
		topicUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q topicQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for topics")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for topics")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TopicSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), topicPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"topics\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, topicPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in topic slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all topic")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Topic) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no topics provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(topicColumnsWithDefault, o)

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

	topicUpsertCacheMut.RLock()
	cache, cached := topicUpsertCache[key]
	topicUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			topicAllColumns,
			topicColumnsWithDefault,
			topicColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			topicAllColumns,
			topicPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert topics, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(topicPrimaryKeyColumns))
			copy(conflict, topicPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"topics\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(topicType, topicMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(topicType, topicMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert topics")
	}

	if !cached {
		topicUpsertCacheMut.Lock()
		topicUpsertCache[key] = cache
		topicUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Topic record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Topic) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Topic provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), topicPrimaryKeyMapping)
	sql := "DELETE FROM \"topics\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from topics")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for topics")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q topicQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no topicQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from topics")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for topics")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TopicSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(topicBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), topicPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"topics\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, topicPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from topic slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for topics")
	}

	if len(topicAfterDeleteHooks) != 0 {
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
func (o *Topic) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTopic(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TopicSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TopicSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), topicPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"topics\".* FROM \"topics\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, topicPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TopicSlice")
	}

	*o = slice

	return nil
}

// TopicExists checks if the Topic row exists.
func TopicExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"topics\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if topics exists")
	}

	return exists, nil
}

// Exists checks if the Topic row exists.
func (o *Topic) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TopicExists(ctx, exec, o.ID)
}