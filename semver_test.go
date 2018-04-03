package semver

import (
	"testing"
)

type StringIntFixture struct {
	String string
	Int    int64
}

func TestGetAsInt(t *testing.T) {
	fixtures := []StringIntFixture{
		StringIntFixture{
			String: "",
			Int:    0,
		},
		StringIntFixture{
			String: "1",
			Int:    1,
		},
		StringIntFixture{
			String: "32",
			Int:    32,
		},
		StringIntFixture{
			String: "1000",
			Int:    1000,
		},
		StringIntFixture{
			String: "abc",
			Int:    0,
		},
	}

	for _, fixture := range fixtures {
		intVal, _ := getAsInt(fixture.String)
		if intVal != fixture.Int {
			t.Errorf("Expected int value (%v) did not match received (%v)", fixture.Int, intVal)
			t.Fail()
		}
	}
}

type SemversionFixture struct {
	String string
	Semver SemVersion
}

func TestGetAsSemversion(t *testing.T) {
	fixtures := []SemversionFixture{
		SemversionFixture{
			String: "1.2.3",
			Semver: SemVersion{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
		},
		SemversionFixture{
			String: "0.2.3",
			Semver: SemVersion{
				Major: 0,
				Minor: 2,
				Patch: 3,
			},
		},
		SemversionFixture{
			String: "1",
			Semver: SemVersion{
				Major: 1,
				Minor: 0,
				Patch: 0,
			},
		},
		SemversionFixture{
			String: "",
			Semver: SemVersion{
				Major: 0,
				Minor: 0,
				Patch: 0,
			},
		},
		SemversionFixture{
			String: "0.0.13",
			Semver: SemVersion{
				Major: 0,
				Minor: 0,
				Patch: 13,
			},
		},
		SemversionFixture{
			String: "1.2.0",
			Semver: SemVersion{
				Major: 1,
				Minor: 2,
				Patch: 0,
			},
		},
	}

	for _, fixture := range fixtures {
		semver, _ := GetAsSemversion(fixture.String)
		if semver != fixture.Semver {
			t.Errorf("Returned semver does not match fixture for string %s", fixture.String)
			t.Fail()
		}
	}
}

type SemVersionCompareFixture struct {
	Version string
	Compare string
	Newer   bool
}

func TestIsNewer(t *testing.T) {
	fixtures := []SemVersionCompareFixture{
		SemVersionCompareFixture{
			Version: "0",
			Compare: "1",
			Newer:   true,
		},
		SemVersionCompareFixture{
			Version: "1",
			Compare: "0",
			Newer:   false,
		},
		SemVersionCompareFixture{
			Version: "1.0.0",
			Compare: "1.0.1",
			Newer:   true,
		},
		SemVersionCompareFixture{
			Version: "1.0.1",
			Compare: "1.1.0",
			Newer:   true,
		},
		SemVersionCompareFixture{
			Version: "1.0.10",
			Compare: "1.0.12",
			Newer:   true,
		},
		SemVersionCompareFixture{
			Version: "1.0.99",
			Compare: "1.1",
			Newer:   true,
		},
		SemVersionCompareFixture{
			Version: "1",
			Compare: "1.0.0",
			Newer:   false,
		},
		SemVersionCompareFixture{
			Version: "1.1.0",
			Compare: "1.2.0",
			Newer:   true,
		},
		SemVersionCompareFixture{
			Version: "2.0.0",
			Compare: "1.2.3",
			Newer:   false,
		},
		SemVersionCompareFixture{
			Version: "1.2.3",
			Compare: "2.0.0",
			Newer:   true,
		},
	}

	for _, fixture := range fixtures {
		newer, _ := IsNewer(fixture.Version, fixture.Compare)
		if newer != fixture.Newer {
			t.Errorf("Results for IsNewer (%v) not what was expected (%v) for version %s", newer, fixture.Newer, fixture.Version)
			t.Fail()
		}
	}
}
