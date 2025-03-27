package services

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// Helper function to reset the users slice before each test
func resetUsers() {
	users = []User{}
}

func TestRegisterUser(t *testing.T) {
	resetUsers() // Ensure a clean state before each test

	// Test case 1: Successful registration
	newUser1 := User{Username: "testuser1", Password: "password1"}
	err := RegisterUser(newUser1)
	if err != nil {
		t.Errorf("Test Case 1 Failed: Expected no error, but got %v", err)
	}
	if len(users) != 1 || users[0] != newUser1 {
		t.Errorf("Test Case 1 Failed: User not registered correctly")
	}

	// Test case 2: Attempt to register with an existing username
	newUser2 := User{Username: "testuser1", Password: "password2"}
	err = RegisterUser(newUser2)
	if err == nil {
		t.Errorf("Test Case 2 Failed: Expected an error for existing username, but got none")
	}
	if len(users) != 1 {
		t.Errorf("Test Case 2 Failed: Should not have added a new user")
	}
}

func TestLoginUser(t *testing.T) {
	resetUsers() // Ensure a clean state

	// Add a user for testing login
	existingUser := User{Username: "testuser", Password: "testpassword"}
	users = append(users, existingUser)

	// Test case 1: Successful login
	user, err := LoginUser("testuser", "testpassword")
	if err != nil {
		t.Errorf("Test Case 1 Failed: Expected no error, but got %v", err)
	}
	if user == nil || !reflect.DeepEqual(*user, existingUser) {
		t.Errorf("Test Case 1 Failed: Login successful, but user details are incorrect")
	}

	// Test case 2: Incorrect password
	user, err = LoginUser("testuser", "wrongpassword")
	if err == nil {
		t.Errorf("Test Case 2 Failed: Expected an error for incorrect password, but got none")
	}
	if user != nil {
		t.Errorf("Test Case 2 Failed: Should not have returned a user for incorrect password")
	}

	// Test case 3: Non-existent username
	user, err = LoginUser("nonexistentuser", "anypassword")
	if err == nil {
		t.Errorf("Test Case 3 Failed: Expected an error for non-existent username, but got none")
	}
	if user != nil {
		t.Errorf("Test Case 3 Failed: Should not have returned a user for non-existent username")
	}
}

func TestFlakyRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := rng.Intn(10)
	t.Logf("Generated number: %d", random)

	// Test for even number
	if random%2 != 0 {
		t.Errorf("Expected an even number, but got %d", random)
	}
}
