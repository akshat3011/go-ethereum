// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	//"encoding/binary"
	//"errors"
	//"fmt"
	//"io/ioutil"
	//"math/big"
	//"math/rand"
	//"os"
	//"reflect"
	//
	//"testing/quick"

    "github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
)
func newEmpty() *Trie {
	trie, _ := New(common.Hash{}, NewDatabase(ethdb.NewMemDatabase()))
	return trie
}

func main() {
	trie := newEmpty()
	updateString(trie, "doe", "reindeer")
	updateString(trie, "dog", "puppy")
	updateString(trie, "dogglesworth", "cat")

}

//func getString(trie *Trie, k string) []byte {
//	return trie.Get([]byte(k))
//}

func updateString(trie *Trie, k, v string,d string) {
	trie.Update([]byte(k), []byte(v), []byte(d))
}

//func deleteString(trie *Trie, k string) {
//	trie.Delete([]byte(k))
//}
