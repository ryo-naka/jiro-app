// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Favorite is an object representing the database table.
type Favorite struct {
	ID        string `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	ContentID string `boil:"content_id" json:"content_id" toml:"content_id" yaml:"content_id"`

	R *favoriteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L favoriteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FavoriteColumns = struct {
	ID        string
	UserID    string
	ContentID string
}{
	ID:        "id",
	UserID:    "user_id",
	ContentID: "content_id",
}

// Generated where

var FavoriteWhere = struct {
	ID        whereHelperstring
	UserID    whereHelperstring
	ContentID whereHelperstring
}{
	ID:        whereHelperstring{field: "`favorite`.`id`"},
	UserID:    whereHelperstring{field: "`favorite`.`user_id`"},
	ContentID: whereHelperstring{field: "`favorite`.`content_id`"},
}

// FavoriteRels is where relationship names are stored.
var FavoriteRels = struct {
	Content string
	User    string
}{
	Content: "Content",
	User:    "User",
}

// favoriteR is where relationships are stored.
type favoriteR struct {
	Content *Content `boil:"Content" json:"Content" toml:"Content" yaml:"Content"`
	User    *User    `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*favoriteR) NewStruct() *favoriteR {
	return &favoriteR{}
}

// favoriteL is where Load methods for each relationship are stored.
type favoriteL struct{}

var (
	favoriteAllColumns            = []string{"id", "user_id", "content_id"}
	favoriteColumnsWithoutDefault = []string{"id", "user_id", "content_id"}
	favoriteColumnsWithDefault    = []string{}
	favoritePrimaryKeyColumns     = []string{"id"}
)

type (
	// FavoriteSlice is an alias for a slice of pointers to Favorite.
	// This should generally be used opposed to []Favorite.
	FavoriteSlice []*Favorite
	// FavoriteHook is the signature for custom Favorite hook methods
	FavoriteHook func(context.Context, boil.ContextExecutor, *Favorite) error

	favoriteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	favoriteType                 = reflect.TypeOf(&Favorite{})
	favoriteMapping              = queries.MakeStructMapping(favoriteType)
	favoritePrimaryKeyMapping, _ = queries.BindMapping(favoriteType, favoriteMapping, favoritePrimaryKeyColumns)
	favoriteInsertCacheMut       sync.RWMutex
	favoriteInsertCache          = make(map[string]insertCache)
	favoriteUpdateCacheMut       sync.RWMutex
	favoriteUpdateCache          = make(map[string]updateCache)
	favoriteUpsertCacheMut       sync.RWMutex
	favoriteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var favoriteBeforeInsertHooks []FavoriteHook
var favoriteBeforeUpdateHooks []FavoriteHook
var favoriteBeforeDeleteHooks []FavoriteHook
var favoriteBeforeUpsertHooks []FavoriteHook

var favoriteAfterInsertHooks []FavoriteHook
var favoriteAfterSelectHooks []FavoriteHook
var favoriteAfterUpdateHooks []FavoriteHook
var favoriteAfterDeleteHooks []FavoriteHook
var favoriteAfterUpsertHooks []FavoriteHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Favorite) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range favoriteBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Favorite) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range favoriteBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Favorite) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range favoriteBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Favorite) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range favoriteBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Favorite) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range favoriteAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Favorite) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range favoriteAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Favorite) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range favoriteAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Favorite) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range favoriteAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Favorite) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range favoriteAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFavoriteHook registers your hook function for all future operations.
func AddFavoriteHook(hookPoint boil.HookPoint, favoriteHook FavoriteHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		favoriteBeforeInsertHooks = append(favoriteBeforeInsertHooks, favoriteHook)
	case boil.BeforeUpdateHook:
		favoriteBeforeUpdateHooks = append(favoriteBeforeUpdateHooks, favoriteHook)
	case boil.BeforeDeleteHook:
		favoriteBeforeDeleteHooks = append(favoriteBeforeDeleteHooks, favoriteHook)
	case boil.BeforeUpsertHook:
		favoriteBeforeUpsertHooks = append(favoriteBeforeUpsertHooks, favoriteHook)
	case boil.AfterInsertHook:
		favoriteAfterInsertHooks = append(favoriteAfterInsertHooks, favoriteHook)
	case boil.AfterSelectHook:
		favoriteAfterSelectHooks = append(favoriteAfterSelectHooks, favoriteHook)
	case boil.AfterUpdateHook:
		favoriteAfterUpdateHooks = append(favoriteAfterUpdateHooks, favoriteHook)
	case boil.AfterDeleteHook:
		favoriteAfterDeleteHooks = append(favoriteAfterDeleteHooks, favoriteHook)
	case boil.AfterUpsertHook:
		favoriteAfterUpsertHooks = append(favoriteAfterUpsertHooks, favoriteHook)
	}
}

// One returns a single favorite record from the query.
func (q favoriteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Favorite, error) {
	o := &Favorite{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for favorite")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Favorite records from the query.
func (q favoriteQuery) All(ctx context.Context, exec boil.ContextExecutor) (FavoriteSlice, error) {
	var o []*Favorite

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Favorite slice")
	}

	if len(favoriteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Favorite records in the query.
func (q favoriteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count favorite rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q favoriteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if favorite exists")
	}

	return count > 0, nil
}

// Content pointed to by the foreign key.
func (o *Favorite) Content(mods ...qm.QueryMod) contentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ContentID),
	}

	queryMods = append(queryMods, mods...)

	query := Contents(queryMods...)
	queries.SetFrom(query.Query, "`content`")

	return query
}

// User pointed to by the foreign key.
func (o *Favorite) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`user`")

	return query
}

// LoadContent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (favoriteL) LoadContent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFavorite interface{}, mods queries.Applicator) error {
	var slice []*Favorite
	var object *Favorite

	if singular {
		object = maybeFavorite.(*Favorite)
	} else {
		slice = *maybeFavorite.(*[]*Favorite)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &favoriteR{}
		}
		args = append(args, object.ContentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &favoriteR{}
			}

			for _, a := range args {
				if a == obj.ContentID {
					continue Outer
				}
			}

			args = append(args, obj.ContentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`content`),
		qm.WhereIn(`content.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Content")
	}

	var resultSlice []*Content
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Content")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for content")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for content")
	}

	if len(favoriteAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Content = foreign
		if foreign.R == nil {
			foreign.R = &contentR{}
		}
		foreign.R.Favorites = append(foreign.R.Favorites, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ContentID == foreign.ID {
				local.R.Content = foreign
				if foreign.R == nil {
					foreign.R = &contentR{}
				}
				foreign.R.Favorites = append(foreign.R.Favorites, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (favoriteL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFavorite interface{}, mods queries.Applicator) error {
	var slice []*Favorite
	var object *Favorite

	if singular {
		object = maybeFavorite.(*Favorite)
	} else {
		slice = *maybeFavorite.(*[]*Favorite)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &favoriteR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &favoriteR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user`),
		qm.WhereIn(`user.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(favoriteAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Favorites = append(foreign.R.Favorites, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Favorites = append(foreign.R.Favorites, local)
				break
			}
		}
	}

	return nil
}

// SetContent of the favorite to the related item.
// Sets o.R.Content to related.
// Adds o to related.R.Favorites.
func (o *Favorite) SetContent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Content) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `favorite` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"content_id"}),
		strmangle.WhereClause("`", "`", 0, favoritePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ContentID = related.ID
	if o.R == nil {
		o.R = &favoriteR{
			Content: related,
		}
	} else {
		o.R.Content = related
	}

	if related.R == nil {
		related.R = &contentR{
			Favorites: FavoriteSlice{o},
		}
	} else {
		related.R.Favorites = append(related.R.Favorites, o)
	}

	return nil
}

// SetUser of the favorite to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Favorites.
func (o *Favorite) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `favorite` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, favoritePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &favoriteR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Favorites: FavoriteSlice{o},
		}
	} else {
		related.R.Favorites = append(related.R.Favorites, o)
	}

	return nil
}

// Favorites retrieves all the records using an executor.
func Favorites(mods ...qm.QueryMod) favoriteQuery {
	mods = append(mods, qm.From("`favorite`"))
	return favoriteQuery{NewQuery(mods...)}
}

// FindFavorite retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFavorite(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Favorite, error) {
	favoriteObj := &Favorite{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `favorite` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, favoriteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from favorite")
	}

	return favoriteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Favorite) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no favorite provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(favoriteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	favoriteInsertCacheMut.RLock()
	cache, cached := favoriteInsertCache[key]
	favoriteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			favoriteAllColumns,
			favoriteColumnsWithDefault,
			favoriteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(favoriteType, favoriteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(favoriteType, favoriteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `favorite` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `favorite` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `favorite` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, favoritePrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into favorite")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for favorite")
	}

CacheNoHooks:
	if !cached {
		favoriteInsertCacheMut.Lock()
		favoriteInsertCache[key] = cache
		favoriteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Favorite.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Favorite) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	favoriteUpdateCacheMut.RLock()
	cache, cached := favoriteUpdateCache[key]
	favoriteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			favoriteAllColumns,
			favoritePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update favorite, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `favorite` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, favoritePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(favoriteType, favoriteMapping, append(wl, favoritePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update favorite row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for favorite")
	}

	if !cached {
		favoriteUpdateCacheMut.Lock()
		favoriteUpdateCache[key] = cache
		favoriteUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q favoriteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for favorite")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for favorite")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FavoriteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), favoritePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `favorite` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, favoritePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in favorite slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all favorite")
	}
	return rowsAff, nil
}

var mySQLFavoriteUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Favorite) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no favorite provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(favoriteColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFavoriteUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	favoriteUpsertCacheMut.RLock()
	cache, cached := favoriteUpsertCache[key]
	favoriteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			favoriteAllColumns,
			favoriteColumnsWithDefault,
			favoriteColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			favoriteAllColumns,
			favoritePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert favorite, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`favorite`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `favorite` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(favoriteType, favoriteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(favoriteType, favoriteMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for favorite")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(favoriteType, favoriteMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for favorite")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for favorite")
	}

CacheNoHooks:
	if !cached {
		favoriteUpsertCacheMut.Lock()
		favoriteUpsertCache[key] = cache
		favoriteUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Favorite record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Favorite) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Favorite provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), favoritePrimaryKeyMapping)
	sql := "DELETE FROM `favorite` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from favorite")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for favorite")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q favoriteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no favoriteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from favorite")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for favorite")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FavoriteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(favoriteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), favoritePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `favorite` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, favoritePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from favorite slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for favorite")
	}

	if len(favoriteAfterDeleteHooks) != 0 {
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
func (o *Favorite) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFavorite(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FavoriteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FavoriteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), favoritePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `favorite`.* FROM `favorite` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, favoritePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FavoriteSlice")
	}

	*o = slice

	return nil
}

// FavoriteExists checks if the Favorite row exists.
func FavoriteExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `favorite` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if favorite exists")
	}

	return exists, nil
}
