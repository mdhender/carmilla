// carmilla - clean up a document
// Copyright (C) 2020  Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	pg, err := ioutil.ReadFile("10007-8.txt")
	if err != nil {
		return err
	}
	pg = append(pg, '\n')
	for _, pp := range strings.Split(string(pg), "\n\n") {
		pp = strings.TrimSpace(strings.Replace(strings.TrimSpace(pp), "\n", " ", -1))
		fmt.Printf("%s\n\n", strings.TrimSpace(pp))
	}
	return nil
}
