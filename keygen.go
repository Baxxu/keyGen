package keygen

import "crypto/rand"

const encoding string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_"

func GetRandKey(len int) (key []byte, err error) {
	key = make([]byte, len)
	_, err = rand.Read(key)
	if err != nil {
		return nil, err
	}

	for i := range key {
		key[i] = encoding[key[i]&63]
	}

	return key, nil
}
