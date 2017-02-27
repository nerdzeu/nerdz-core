Nerdz Core
=========

Nerdz Core is a slimmed-down fork of Nerdz API which strives to implement a simplified, streamlined backend service for Nerdz, to be used as the entrypoint for several other services (including API and Web itself).

Work in progress.

# Roadmap

1. Trim away all of the REST/API logic from the Nerdz API db access package (`nerdz`->`db`)
2. Define Nerdz transfer objects using Protocol Buffer 
3. Implement a TLS authenticated gRPC service to provide an unified access point to a Nerdz instance
4. Create clients

# Contributing

Feel free to contribute with code, documentation, by running tests or by reporting issues. 

# License
Copyright (C) 2016 Paolo Galeone; nessuno@nerdz.eu

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
