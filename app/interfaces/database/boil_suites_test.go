// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package database

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("UserPasswords", testUserPasswords)
	t.Run("UserTokens", testUserTokens)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsDelete)
	t.Run("UserTokens", testUserTokensDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsQueryDeleteAll)
	t.Run("UserTokens", testUserTokensQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsSliceDeleteAll)
	t.Run("UserTokens", testUserTokensSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsExists)
	t.Run("UserTokens", testUserTokensExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsFind)
	t.Run("UserTokens", testUserTokensFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsBind)
	t.Run("UserTokens", testUserTokensBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsOne)
	t.Run("UserTokens", testUserTokensOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsAll)
	t.Run("UserTokens", testUserTokensAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsCount)
	t.Run("UserTokens", testUserTokensCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsHooks)
	t.Run("UserTokens", testUserTokensHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsInsert)
	t.Run("UserPasswords", testUserPasswordsInsertWhitelist)
	t.Run("UserTokens", testUserTokensInsert)
	t.Run("UserTokens", testUserTokensInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("UserPasswordToUserUsingUser", testUserPasswordToOneUserUsingUser)
	t.Run("UserTokenToUserUsingUser", testUserTokenToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("UserToUserPasswordUsingUserPassword", testUserOneToOneUserPasswordUsingUserPassword)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("UserToUserTokens", testUserToManyUserTokens)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("UserPasswordToUserUsingUserPassword", testUserPasswordToOneSetOpUserUsingUser)
	t.Run("UserTokenToUserUsingUserTokens", testUserTokenToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("UserToUserPasswordUsingUserPassword", testUserOneToOneSetOpUserPasswordUsingUserPassword)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("UserToUserTokens", testUserToManyAddOpUserTokens)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsReload)
	t.Run("UserTokens", testUserTokensReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsReloadAll)
	t.Run("UserTokens", testUserTokensReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsSelect)
	t.Run("UserTokens", testUserTokensSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsUpdate)
	t.Run("UserTokens", testUserTokensUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("UserPasswords", testUserPasswordsSliceUpdateAll)
	t.Run("UserTokens", testUserTokensSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
