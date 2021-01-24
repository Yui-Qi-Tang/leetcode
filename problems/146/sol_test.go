package sol

import "testing"

type testValue struct {
	action string
	key    int
	value  int // used in get as 'want' of testcase
}

type testcase struct {
	size    int
	actions []testValue
}

func TestBasic(t *testing.T) {
	lr := Constructor(2)
	lr.Put(1, 2)

	r := lr.Get(1)

	if r != 2 {
		t.Fatalf("it should be %d, but got: %d", 2, r)
	}

	t.Log("passed")
}

func TestCase1(t *testing.T) {
	testcases := testcase{
		size: 3,
		actions: []testValue{
			{action: "put", key: 1, value: 1},
			{action: "put", key: 2, value: 2},
			{action: "put", key: 3, value: 3},
			{action: "put", key: 4, value: 4},

			{action: "get", key: 4, value: 4},
			{action: "get", key: 3, value: 3},
			{action: "get", key: 2, value: 2},
			{action: "get", key: 1, value: -1},

			{action: "put", key: 5, value: 5},

			{action: "get", key: 1, value: -1},
			{action: "get", key: 2, value: 2},
			{action: "get", key: 3, value: 3},
			{action: "get", key: 4, value: -1},
			{action: "get", key: 5, value: 5},
		},
	}

	lru := Constructor(testcases.size)

	for _, action := range testcases.actions {

		if action.action == "put" {
			lru.Put(action.key, action.value)
		}

		if action.action == "get" {
			want := action.value
			result := lru.Get(action.key)

			if result != want {
				t.Fatalf("expected %d, but got %d", want, result)
			}
		}

	}

	t.Log("passed")
}

func TestCase2(t *testing.T) {
	testcases := testcase{
		size: 2,
		actions: []testValue{
			{action: "put", key: 1, value: 1},
			{action: "put", key: 2, value: 2},

			{action: "get", key: 1, value: 1},

			{action: "put", key: 3, value: 3},

			{action: "get", key: 2, value: -1},

			{action: "put", key: 4, value: 4},

			{action: "get", key: 1, value: -1},

			{action: "get", key: 3, value: 3},

			{action: "get", key: 4, value: 4},
		},
	}

	lru := Constructor(testcases.size)

	for _, action := range testcases.actions {

		if action.action == "put" {
			lru.Put(action.key, action.value)
		}

		if action.action == "get" {
			want := action.value
			result := lru.Get(action.key)

			if result != want {
				t.Fatalf("expected %d, but got %d", want, result)
			}
		}

	}

	t.Log("passed")
}
