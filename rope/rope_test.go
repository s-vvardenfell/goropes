package rope

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const stringEn = "A sling (rope, cord) is a data structure" +
	" that allows you to efficiently" +
	" store and process long strings, such as text."

func Test_NewRope(t *testing.T) {
	t.Log("creating new rope from a string")
	{
		r, err := NewRopeFromString(stringEn, 8)
		require.NoError(t, err)
		require.Equal(t, r.String(), stringEn)
	}

	t.Log("creating new rope from a Reader")
	{
		wd, err := os.Getwd()
		require.NoError(t, err)
		path := filepath.Join(filepath.Join(wd, ".."), "resources/test.txt")
		f, err := os.Open(path)
		require.NoError(t, err)

		r, err := NewRopeFromReader(f, 8)
		require.NoError(t, err)
		require.Equal(t, r.String(), stringEn) // i removed \n from file ending
	}
}
