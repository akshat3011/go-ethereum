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
	"fmt"
	//"io/ioutil"
	//"math/big"
	//"math/rand"
	//"os"
	//"reflect"
	"bytes"
	"testing"
	//"testing/quick"

   // "github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
   // "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	//"github.com/ethereum/go-ethereum/rlp"
)
func newEmpty() *Trie {
	trie, _ := New(common.Hash{}, NewDatabase(ethdb.NewMemDatabase()))
	return trie
}

func TestGet(t *testing.T) {
	trie := newEmpty()
	updateString(trie, "doe", "reindeer","hey")
	updateString(trie, "dog", "puppy","sad")
	updateString(trie, "dogglesworth", "cat","meow")
	trie.Commit(nil)
	for i := 0; i < 2; i++ {
		res,data := getString(trie, "dog")
		if !bytes.Equal(res, []byte("puppy")) {
			t.Errorf("expected puppy got %x", res)
		}
		fmt.Println("here is data")
		fmt.Println(string(data))
		if i == 1 {
			return
		}
		
	}
}

func TestDelete(t *testing.T) {
	trie := newEmpty()
	vals := []struct{ k, v string }{
		{"do", "verb"},
		{"ether", "wookiedoo"},
		{"horse", "stallion"},
		{"shaman", "horse"},
		{"doge", "coin"},
		{"ether", ""},
		{"dog", "puppy"},
		{"shaman", ""},
	}

	for _, val := range vals {
		if val.v != "" {
			updateString(trie, val.k, val.v,"supposed")
		} else {
			res,_ := getString(trie,val.k)
			fmt.Println(string(res))
			deleteString(trie, val.k)
			fmt.Println("after deletion")
			res,_ = getString(trie,val.k)
			fmt.Println(string(res))
		}
	}


	}






func getString(trie *Trie, k string) ([]byte ,[]byte){
	return trie.Get([]byte(k))
}

func updateString(trie *Trie, k, v string,d string) {
	trie.Update([]byte(k), []byte(v), []byte(d))
}

func deleteString(trie *Trie, k string) {
trie.Delete([]byte(k))
}
