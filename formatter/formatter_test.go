package formatter

import (
	"strings"
	"testing"
)

// convertTabsToSpaces はテキスト内のタブをスペース4個に変換する
func convertTabsToSpaces(text string) string {
	return strings.ReplaceAll(text, "\t", "    ")
}

func TestFormatKotlinData1(t *testing.T) {
	input := `User(name=name, age=20)`
	expected := `User(
	name = name,
	age = 20
)`
	expected = convertTabsToSpaces(expected)

	actual := formatKotlinData(input)
	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestFormatKotlinData2(t *testing.T) {
	input := `[User(name=name1, age=20), User(name=name2, age=30), User(name=name3, age=40)]`
	expected := `[
    User(
		name = name1,
		age = 20
	),
	User(
		name = name2,
		age = 30
	),
	User(
		name = name3,
		age = 40
	)
]`
	expected = convertTabsToSpaces(expected)

	actual := formatKotlinData(input)
	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestFormatKotlinData3(t *testing.T) {
	input := `MyPage(user=User(name=name, age=20), coin=100, posts=4)`
	expected := `MyPage(
	user = User(
		name = name,
		age = 20
	),
	coin = 100,
	posts = 4
)`
	expected = convertTabsToSpaces(expected)

	actual := formatKotlinData(input)
	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestFormatKotlinData4(t *testing.T) {
	input := `[Category(id=1, name=category1, locations=[Location(id=1, name=location1)]), Category(id=2, name=category2, locations=[Location(id=2, name=location2), Location(id=3, name=location3)]), Category(id=3, name=category3, locations=[Location(id=4, name=location4)])]`
	expected := `[
	Category(
		id = 1,
		name = category1,
		locations = [
			Location(
				id = 1,
				name = location1
			)
		]
	),
	Category(
		id = 2,
		name = category2,
		locations = [
			Location(
				id = 2,
				name = location2
			),
			Location(
				id = 3,
				name = location3
			)
		]
	),
	Category(
		id = 3,
		name = category3,
		locations = [
			Location(
				id = 4,
				name = location4
			)
		]
	)
]`
	expected = convertTabsToSpaces(expected)

	actual := formatKotlinData(input)
	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestFormatKotlinData5(t *testing.T) {
	input := `"User(name=name, age=20)"`
	expected := `User(
	name = name,
	age = 20
)`
	expected = convertTabsToSpaces(expected)

	actual := formatKotlinData(input)
	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}
