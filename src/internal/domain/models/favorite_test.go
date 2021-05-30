// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testFavorites(t *testing.T) {
	t.Parallel()

	query := Favorites()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFavoritesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFavoritesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Favorites().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFavoritesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FavoriteSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFavoritesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FavoriteExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Favorite exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FavoriteExists to return true, but got false.")
	}
}

func testFavoritesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	favoriteFound, err := FindFavorite(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if favoriteFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFavoritesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Favorites().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFavoritesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Favorites().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFavoritesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	favoriteOne := &Favorite{}
	favoriteTwo := &Favorite{}
	if err = randomize.Struct(seed, favoriteOne, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}
	if err = randomize.Struct(seed, favoriteTwo, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = favoriteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = favoriteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Favorites().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFavoritesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	favoriteOne := &Favorite{}
	favoriteTwo := &Favorite{}
	if err = randomize.Struct(seed, favoriteOne, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}
	if err = randomize.Struct(seed, favoriteTwo, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = favoriteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = favoriteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func favoriteBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Favorite) error {
	*o = Favorite{}
	return nil
}

func favoriteAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Favorite) error {
	*o = Favorite{}
	return nil
}

func favoriteAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Favorite) error {
	*o = Favorite{}
	return nil
}

func favoriteBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Favorite) error {
	*o = Favorite{}
	return nil
}

func favoriteAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Favorite) error {
	*o = Favorite{}
	return nil
}

func favoriteBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Favorite) error {
	*o = Favorite{}
	return nil
}

func favoriteAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Favorite) error {
	*o = Favorite{}
	return nil
}

func favoriteBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Favorite) error {
	*o = Favorite{}
	return nil
}

func favoriteAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Favorite) error {
	*o = Favorite{}
	return nil
}

func testFavoritesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Favorite{}
	o := &Favorite{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, favoriteDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Favorite object: %s", err)
	}

	AddFavoriteHook(boil.BeforeInsertHook, favoriteBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	favoriteBeforeInsertHooks = []FavoriteHook{}

	AddFavoriteHook(boil.AfterInsertHook, favoriteAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	favoriteAfterInsertHooks = []FavoriteHook{}

	AddFavoriteHook(boil.AfterSelectHook, favoriteAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	favoriteAfterSelectHooks = []FavoriteHook{}

	AddFavoriteHook(boil.BeforeUpdateHook, favoriteBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	favoriteBeforeUpdateHooks = []FavoriteHook{}

	AddFavoriteHook(boil.AfterUpdateHook, favoriteAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	favoriteAfterUpdateHooks = []FavoriteHook{}

	AddFavoriteHook(boil.BeforeDeleteHook, favoriteBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	favoriteBeforeDeleteHooks = []FavoriteHook{}

	AddFavoriteHook(boil.AfterDeleteHook, favoriteAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	favoriteAfterDeleteHooks = []FavoriteHook{}

	AddFavoriteHook(boil.BeforeUpsertHook, favoriteBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	favoriteBeforeUpsertHooks = []FavoriteHook{}

	AddFavoriteHook(boil.AfterUpsertHook, favoriteAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	favoriteAfterUpsertHooks = []FavoriteHook{}
}

func testFavoritesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFavoritesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(favoriteColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFavoriteToOneContentUsingContent(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Favorite
	var foreign Content

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, contentDBTypes, false, contentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Content struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ContentID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Content().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FavoriteSlice{&local}
	if err = local.L.LoadContent(ctx, tx, false, (*[]*Favorite)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Content == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Content = nil
	if err = local.L.LoadContent(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Content == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFavoriteToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Favorite
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FavoriteSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Favorite)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFavoriteToOneSetOpContentUsingContent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Favorite
	var b, c Content

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, favoriteDBTypes, false, strmangle.SetComplement(favoritePrimaryKeyColumns, favoriteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, contentDBTypes, false, strmangle.SetComplement(contentPrimaryKeyColumns, contentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contentDBTypes, false, strmangle.SetComplement(contentPrimaryKeyColumns, contentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Content{&b, &c} {
		err = a.SetContent(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Content != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Favorites[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ContentID != x.ID {
			t.Error("foreign key was wrong value", a.ContentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ContentID))
		reflect.Indirect(reflect.ValueOf(&a.ContentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ContentID != x.ID {
			t.Error("foreign key was wrong value", a.ContentID, x.ID)
		}
	}
}
func testFavoriteToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Favorite
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, favoriteDBTypes, false, strmangle.SetComplement(favoritePrimaryKeyColumns, favoriteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Favorites[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testFavoritesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFavoritesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FavoriteSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFavoritesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Favorites().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	favoriteDBTypes = map[string]string{`ID`: `varchar`, `UserID`: `varchar`, `ContentID`: `varchar`}
	_               = bytes.MinRead
)

func testFavoritesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(favoritePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(favoriteAllColumns) == len(favoritePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoritePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFavoritesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(favoriteAllColumns) == len(favoritePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Favorite{}
	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, favoriteDBTypes, true, favoritePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(favoriteAllColumns, favoritePrimaryKeyColumns) {
		fields = favoriteAllColumns
	} else {
		fields = strmangle.SetComplement(
			favoriteAllColumns,
			favoritePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := FavoriteSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFavoritesUpsert(t *testing.T) {
	t.Parallel()

	if len(favoriteAllColumns) == len(favoritePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLFavoriteUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Favorite{}
	if err = randomize.Struct(seed, &o, favoriteDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Favorite: %s", err)
	}

	count, err := Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, favoriteDBTypes, false, favoritePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Favorite struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Favorite: %s", err)
	}

	count, err = Favorites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
