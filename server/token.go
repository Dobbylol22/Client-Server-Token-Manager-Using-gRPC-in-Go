package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"sync"
)

// Hash concatentates a message and a nonce and generates a hash value.
func Hash(name string, nonce uint64) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s %d", name, nonce)))
	return binary.BigEndian.Uint64(hasher.Sum(nil))
}

//to calculate the partial value (when write function is called) and (final value when read function is called)
func MinHash(name string, start uint64, end uint64) (min uint64) {
	min = math.MaxUint64
	for i := start; i < end; i++ {
		current := Hash(name, i)
		if current < min {
			min = current
		}
	}
	return
}

type TokenDomain struct {
	Low  uint64
	Mid  uint64
	High uint64
}

type TokenState struct {
	Partial uint64
	Final   uint64
}

type Token struct {
	Id     string
	Name   string
	Domain TokenDomain
	State  *TokenState

	lock sync.RWMutex
}

func (token *Token) Write(name string, low uint64, mid uint64, high uint64) {
	token.Name = name
	token.Domain.Low = low
	token.Domain.Mid = mid
	token.Domain.High = high
	state := &TokenState{}
	state.Partial = MinHash(name, low, mid)
	state.Final = 0 // reset to zero
	token.State = state
}

func (token *Token) Read() {
	min := MinHash(token.Name, token.Domain.Mid, token.Domain.High)
	if token.State != nil {
		if token.State.Partial < min {
			min = token.State.Partial
		}
		token.State.Final = min
	} else {
		token.State = &TokenState{Final: min}
	}
}

func (token *Token) Log() {
	log.Printf("Token: (Id=%+v, Name=%+v, Domain=%+v, State=%+v)\n", token.Id, token.Name, token.Domain, token.State)
}

type TokenManager struct {
	tokens *sync.Map //assigning Map package by importing sync
}

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewTokenManager() (t *TokenManager) {
	t = &TokenManager{tokens: &sync.Map{}}
	return
}

func (t *TokenManager) Create(id string) (err error) {
	defer t.LogTokenIds()
	token := &Token{Id: id}
	token.Log()
	_, loaded := t.tokens.LoadOrStore(id, token) //Map function (LOADORSTORE, LOADANDDELETE,LOAD)
	if loaded {
		err = &Error{Message: "Duplicate Token Id. Token Id is already added."}
		return
	}
	return
}

func (t *TokenManager) Drop(id string) (err error) {
	defer t.LogTokenIds()
	tok, loaded := t.tokens.LoadAndDelete(id)
	if loaded {
		token, ok := tok.(*Token)
		if !ok {
			err = &Error{Message: "Failed to cast back to token"}
			return
		}
		token.Log()
	}
	return
}

func (t *TokenManager) Write(id string, name string, low uint64, mid uint64, high uint64) (partial uint64, err error) {
	defer t.LogTokenIds()
	tok, ok := t.tokens.Load(id)
	if !ok {
		err = &Error{Message: "Token Id not found."}
		return
	}
	token, ok := tok.(*Token)
	if !ok {
		err = &Error{Message: "Failed to cast back to token"}
		return
	}

	token.lock.Lock()
	defer token.lock.Unlock() //RWMutex functions

	token.Write(name, low, mid, high)
	token.Log()

	return token.State.Partial, nil
}

func (t *TokenManager) Read(id string) (final uint64, err error) {
	defer t.LogTokenIds()
	tok, ok := t.tokens.Load(id)
	if !ok {
		err = &Error{Message: "Token Id not found."}
		return
	}
	token, ok := tok.(*Token)
	if !ok {
		err = &Error{Message: "Failed to cast back to token"}
		return
	}

	token.lock.Lock() //RWMutex functions
	defer token.lock.Unlock()

	token.Read()
	token.Log()

	return token.State.Final, nil
}

func (t *TokenManager) LogTokenIds() {
	ids := []string{}
	t.tokens.Range(func(key, value interface{}) bool {
		ids = append(ids, key.(string))
		return true
	})
	log.Printf("List of ids: %+v\n", ids)
}
