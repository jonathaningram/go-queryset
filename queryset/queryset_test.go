package queryset

import (
	"os"
	"testing"
)

func TestUserSelect(t *testing.T) {
	//test.UserQuerySet{}.All(ret)
}

func TestMain(m *testing.M) {
	err := GenerateQuerySets("test/models.go", "test/autogenerated_models.go")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
