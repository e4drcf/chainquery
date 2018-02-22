// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testTransactionAddresses(t *testing.T) {
	t.Parallel()

	query := TransactionAddresses(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testTransactionAddressesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = transactionAddress.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionAddressesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = TransactionAddresses(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionAddressesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TransactionAddressSlice{transactionAddress}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testTransactionAddressesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TransactionAddressExists(tx, transactionAddress.ID)
	if err != nil {
		t.Errorf("Unable to check if TransactionAddress exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TransactionAddressExistsG to return true, but got false.")
	}
}
func testTransactionAddressesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	transactionAddressFound, err := FindTransactionAddress(tx, transactionAddress.ID)
	if err != nil {
		t.Error(err)
	}

	if transactionAddressFound == nil {
		t.Error("want a record, got nil")
	}
}
func testTransactionAddressesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = TransactionAddresses(tx).Bind(transactionAddress); err != nil {
		t.Error(err)
	}
}

func testTransactionAddressesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := TransactionAddresses(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTransactionAddressesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddressOne := &TransactionAddress{}
	transactionAddressTwo := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddressOne, transactionAddressDBTypes, false, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionAddressTwo, transactionAddressDBTypes, false, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddressOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = transactionAddressTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := TransactionAddresses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTransactionAddressesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	transactionAddressOne := &TransactionAddress{}
	transactionAddressTwo := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddressOne, transactionAddressDBTypes, false, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionAddressTwo, transactionAddressDBTypes, false, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddressOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = transactionAddressTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTransactionAddressesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionAddressesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx, transactionAddressColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionAddressToManyInputs(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a TransactionAddress
	var b, c Input

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, inputDBTypes, false, inputColumnsWithDefault...)
	randomize.Struct(seed, &c, inputDBTypes, false, inputColumnsWithDefault...)

	b.TransactionAddressID = a.ID
	c.TransactionAddressID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	input, err := a.Inputs(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range input {
		if v.TransactionAddressID == b.TransactionAddressID {
			bFound = true
		}
		if v.TransactionAddressID == c.TransactionAddressID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TransactionAddressSlice{&a}
	if err = a.L.LoadInputs(tx, false, (*[]*TransactionAddress)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Inputs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Inputs = nil
	if err = a.L.LoadInputs(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Inputs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", input)
	}
}

func testTransactionAddressToManyOutputsAddresses(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a TransactionAddress
	var b, c OutputsAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, outputsAddressDBTypes, false, outputsAddressColumnsWithDefault...)
	randomize.Struct(seed, &c, outputsAddressDBTypes, false, outputsAddressColumnsWithDefault...)

	b.TransactionAddressID = a.ID
	c.TransactionAddressID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	outputsAddress, err := a.OutputsAddresses(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range outputsAddress {
		if v.TransactionAddressID == b.TransactionAddressID {
			bFound = true
		}
		if v.TransactionAddressID == c.TransactionAddressID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TransactionAddressSlice{&a}
	if err = a.L.LoadOutputsAddresses(tx, false, (*[]*TransactionAddress)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OutputsAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OutputsAddresses = nil
	if err = a.L.LoadOutputsAddresses(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OutputsAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", outputsAddress)
	}
}

func testTransactionAddressToManyAddOpInputs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a TransactionAddress
	var b, c, d, e Input

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionAddressDBTypes, false, strmangle.SetComplement(transactionAddressPrimaryKeyColumns, transactionAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Input{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Input{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddInputs(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TransactionAddressID {
			t.Error("foreign key was wrong value", a.ID, first.TransactionAddressID)
		}
		if a.ID != second.TransactionAddressID {
			t.Error("foreign key was wrong value", a.ID, second.TransactionAddressID)
		}

		if first.R.TransactionAddress != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.TransactionAddress != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Inputs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Inputs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Inputs(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testTransactionAddressToManyAddOpOutputsAddresses(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a TransactionAddress
	var b, c, d, e OutputsAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionAddressDBTypes, false, strmangle.SetComplement(transactionAddressPrimaryKeyColumns, transactionAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OutputsAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, outputsAddressDBTypes, false, strmangle.SetComplement(outputsAddressPrimaryKeyColumns, outputsAddressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*OutputsAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOutputsAddresses(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TransactionAddressID {
			t.Error("foreign key was wrong value", a.ID, first.TransactionAddressID)
		}
		if a.ID != second.TransactionAddressID {
			t.Error("foreign key was wrong value", a.ID, second.TransactionAddressID)
		}

		if first.R.TransactionAddress != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.TransactionAddress != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OutputsAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OutputsAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OutputsAddresses(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testTransactionAddressToOneAddressUsingAddress(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local TransactionAddress
	var foreign Address

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, transactionAddressDBTypes, false, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, addressDBTypes, false, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AddressID = foreign.Address
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Address(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Address != foreign.Address {
		t.Errorf("want: %v, got %v", foreign.Address, check.Address)
	}

	slice := TransactionAddressSlice{&local}
	if err = local.L.LoadAddress(tx, false, (*[]*TransactionAddress)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Address == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Address = nil
	if err = local.L.LoadAddress(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Address == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTransactionAddressToOneTransactionUsingTransaction(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local TransactionAddress
	var foreign Transaction

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, transactionAddressDBTypes, false, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.TransactionID = foreign.Hash
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Transaction(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Hash != foreign.Hash {
		t.Errorf("want: %v, got %v", foreign.Hash, check.Hash)
	}

	slice := TransactionAddressSlice{&local}
	if err = local.L.LoadTransaction(tx, false, (*[]*TransactionAddress)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Transaction = nil
	if err = local.L.LoadTransaction(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTransactionAddressToOneSetOpAddressUsingAddress(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a TransactionAddress
	var b, c Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionAddressDBTypes, false, strmangle.SetComplement(transactionAddressPrimaryKeyColumns, transactionAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Address{&b, &c} {
		err = a.SetAddress(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Address != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TransactionAddresses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AddressID != x.Address {
			t.Error("foreign key was wrong value", a.AddressID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AddressID))
		reflect.Indirect(reflect.ValueOf(&a.AddressID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AddressID != x.Address {
			t.Error("foreign key was wrong value", a.AddressID, x.Address)
		}
	}
}
func testTransactionAddressToOneSetOpTransactionUsingTransaction(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a TransactionAddress
	var b, c Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionAddressDBTypes, false, strmangle.SetComplement(transactionAddressPrimaryKeyColumns, transactionAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Transaction{&b, &c} {
		err = a.SetTransaction(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Transaction != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TransactionAddresses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TransactionID != x.Hash {
			t.Error("foreign key was wrong value", a.TransactionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TransactionID))
		reflect.Indirect(reflect.ValueOf(&a.TransactionID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TransactionID != x.Hash {
			t.Error("foreign key was wrong value", a.TransactionID, x.Hash)
		}
	}
}
func testTransactionAddressesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = transactionAddress.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTransactionAddressesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TransactionAddressSlice{transactionAddress}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testTransactionAddressesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := TransactionAddresses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	transactionAddressDBTypes = map[string]string{`AddressID`: `varchar`, `CreditAmount`: `decimal`, `DebitAmount`: `decimal`, `ID`: `varchar`, `LatestTransactionTime`: `datetime`, `TransactionID`: `varchar`}
	_                         = bytes.MinRead
)

func testTransactionAddressesUpdate(t *testing.T) {
	t.Parallel()

	if len(transactionAddressColumns) == len(transactionAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	if err = transactionAddress.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTransactionAddressesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(transactionAddressColumns) == len(transactionAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	transactionAddress := &TransactionAddress{}
	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, transactionAddress, transactionAddressDBTypes, true, transactionAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(transactionAddressColumns, transactionAddressPrimaryKeyColumns) {
		fields = transactionAddressColumns
	} else {
		fields = strmangle.SetComplement(
			transactionAddressColumns,
			transactionAddressPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(transactionAddress))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TransactionAddressSlice{transactionAddress}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testTransactionAddressesUpsert(t *testing.T) {
	t.Parallel()

	if len(transactionAddressColumns) == len(transactionAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	transactionAddress := TransactionAddress{}
	if err = randomize.Struct(seed, &transactionAddress, transactionAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionAddress.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert TransactionAddress: %s", err)
	}

	count, err := TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &transactionAddress, transactionAddressDBTypes, false, transactionAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TransactionAddress struct: %s", err)
	}

	if err = transactionAddress.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert TransactionAddress: %s", err)
	}

	count, err = TransactionAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
