package redirect

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	bolt "go.etcd.io/bbolt"
)

type Redirect struct {
	label       string
	destination string
}

func NewRedirect(label, destination string) (*Redirect, error) {
	u, err := url.Parse(destination)
	if err != nil {
		return nil, err
	}

	e := url.QueryEscape(label)
	if e != label {
		return nil, errors.New("invalid label")
	}

	return &Redirect{
		label:       e,
		destination: u.String(),
	}, nil
}

func (r Redirect) String() string {
	return fmt.Sprintf("%s (%s)", r.label, r.destination)
}

func (r Redirect) GetLabel() string {
	return r.label
}

func (r Redirect) GetDestination() string {
	return r.destination
}

func (r Redirect) Save() error {
	db, err := bolt.Open(dbname, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(redirectBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		b.Put([]byte(r.label), []byte(r.destination))

		return nil
	})
}

func GetRedirect(label string) (Redirect, error) {
	r := Redirect{}

	db, err := bolt.Open(dbname, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return r, err
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(redirectBucket))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		v := b.Get([]byte(label))
		if v == nil {
			return errors.New("key not found")
		}

		r.label = label
		r.destination = string(v)

		return nil
	})

	return r, err
}

func DeleteRedirect(label string) error {
	db, err := bolt.Open(dbname, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(redirectBucket))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.Delete([]byte(label))
	})
}

func GetAllRedirects() ([]Redirect, error) {
	db, err := bolt.Open(dbname, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rs := make([]Redirect, 0)

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(redirectBucket))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		b.ForEach(func(k, v []byte) error {
			r := Redirect{
				label:       string(k),
				destination: string(v),
			}
			rs = append(rs, r)
			return nil
		})

		return nil
	})

	return rs, err

}
