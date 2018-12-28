package userpb

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/fzerorubigd/balloon/pkg/assert"
	"github.com/fzerorubigd/balloon/pkg/kv"
	"github.com/fzerorubigd/balloon/pkg/random"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// From the bcrypt package
const (
	minHashSize  = 59
	noPassString = "NO" // Size must be less than 6 character
)

//  TODO: NEEDS COMMENT INFO
var (
	isBcrypt = regexp.MustCompile(`^\$[^$]+\$[0-9]+\$`)
)

func (m *User) cryptPassword() {
	// TODO : Watch it if this creepy code is dangerous :)
	if (len(m.Password) < minHashSize || !isBcrypt.MatchString(m.Password)) && m.Password != noPassString {
		p, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
		assert.Nil(err)
		m.Password = string(p)
	}
}

// PreInsert the user on create
func (m *User) PreInsert() {
	m.cryptPassword()
}

// PreUpdate the user on update
func (m *User) PreUpdate() {
	m.cryptPassword()
}

// VerifyPassword try to verify password for given hash
func (m *User) VerifyPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password)) == nil
}

// FindUserByEmailPassword try to login user with username and password
func (m *Manager) FindUserByEmailPassword(ctx context.Context, email, password string) (*User, error) {
	u, err := m.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if u.Status == UserStatus_USER_STATUS_BANNED {
		return nil, errors.New("sorry, but you are banned")
	}

	if u.VerifyPassword(password) {
		return u, nil
	}

	return nil, errors.New("user not found or wrong password")
}

// FindUserByEmail is a function to find user based on app
func (m *Manager) FindUserByEmail(ctx context.Context, e string) (*User, error) {
	q := fmt.Sprintf(
		"SELECT %s FROM %s WHERE email = $1 ",
		strings.Join(m.getUserFields(), ","),
		UserTableFull,
	)

	r := m.GetDbMap().QueryRowxContext(ctx, q, e)

	return m.scanUser(r)
}

// RegisterUser is to register new user
func (m *Manager) RegisterUser(ctx context.Context, email, name, pass string) (*User, error) {
	u := User{
		Email:       email,
		DisplayName: name,
		Password:    pass,
		Status:      UserStatus_USER_STATUS_REGISTERED,
	}

	if err := m.CreateUser(ctx, &u); err != nil {
		return nil, errors.Wrap(err, "already registered")
	}

	return &u, nil
}

// CreateToken TODO: NEEDS COMMENT INFO
func (m *Manager) CreateToken(_ context.Context, u *User, d time.Duration) string {
	t := <-random.ID
	v, err := proto.Marshal(u)
	assert.Nil(err)
	kv.MustStoreKey(t, string(v), d)
	return t
}

// FindUserByIndirectToken TODO: NEEDS COMMENT INFO
func (m *Manager) FindUserByIndirectToken(ctx context.Context, token string) (*User, error) {
	t, err := kv.FetchKey(token)
	if err != nil {
		return nil, err
	}
	var u User
	// Invalid data is a bug
	assert.Nil(proto.Unmarshal([]byte(t), &u))

	return &u, nil
}

// DeleteToken TODO: NEEDS COMMENT INFO
func (m *Manager) DeleteToken(_ context.Context, token string) {
	kv.MustDeleteKey(token)
}

// ChangePassword TODO: NEEDS COMMENT INFO
func (m *Manager) ChangePassword(ctx context.Context, u *User, newPassword string) error {
	u.Password = newPassword
	return m.UpdateUser(ctx, u)
}