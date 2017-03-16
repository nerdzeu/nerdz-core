/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package db

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/galeone/igor"
	"github.com/spf13/viper"
)

var dbInst *igor.Database

// Init initialises the internal Database instance.
// Using the package before calling Init will cause the application to panic.
func Init() error {
	connectionString, err := connectionString()
	if err != nil {
		return err
	}

	if dbInst, err = igor.Connect(connectionString); err != nil {
		return err
	}

	logger := log.New(os.Stdout, "query-logger: ", log.LUTC)
	dbInst.Log(logger)

	return nil
}

// connectionString uses viper to access the database configuration, to create a global db instance.
func connectionString() (string, error) {
	setDefaults()

	username := viper.GetString(unameKey)
	if username == "" {
		return "", errors.New("empty database username")
	}

	name := viper.GetString(dbKey)
	if name == "" {
		return "", errors.New("Empty db name")
	}

	var ret bytes.Buffer
	ret.WriteString("user=" + username + " dbname=" + name + " host=" + viper.GetString(hostKey))

	passwd := viper.GetString(passKey)

	if passwd != "" {
		ret.WriteString(" password=" + passwd)
	}

	ret.WriteString(" sslmode=" + viper.GetString(sslKey))

	ret.WriteString(" port=" + strconv.Itoa(viper.GetInt(portKey)))

	return ret.String(), nil
}

// db is used by this package to access an *igor.Database
func db() *igor.Database {
	if dbInst == nil {
		panic("db not yet initialised")
	}

	return dbInst
}

const (
	viperScope = "db."

	unameKey = viperScope + "user"
	dbKey    = viperScope + "name"
	hostKey  = viperScope + "host"
	passKey  = viperScope + "password"
	portKey  = viperScope + "port"
	sslKey   = viperScope + "ssl"
)

// setDefaults sets into viper the default values to access the database.
// This packages namespaces each one of its keys with 'db.'.
func setDefaults() {
	viper.SetDefault(hostKey, "localhost")
	viper.SetDefault(portKey, 5432)
	viper.SetDefault(sslKey, "disable")
}
