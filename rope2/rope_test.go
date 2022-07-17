package rope2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const stringEn = "A sling (rope, cord) is a data structure" +
	" that allows you to efficiently" +
	" store and process long strings, such as text."

const stringRu = "Строп (канат, корд) — структура данных," +
	" которая позволяет эффективно" +
	" хранить и обрабатывать длинные строки, например текст."

func Test_NewRope(t *testing.T) {
	t.Log("creating new rope from a ASCII-string")
	{
		r, err := NewGoRopeFromString(stringEn, 8)
		require.NoError(t, err)
		require.Equal(t, r.String(), stringEn)
	}

	t.Log("creating new rope from a non-ASCII-string")
	{
		r, err := NewGoRopeFromString(stringRu, 8)
		require.NoError(t, err)
		require.Equal(t, r.String(), stringRu)
	}
}
