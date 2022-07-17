package rope

import (
	"math/rand"
	"testing"
	"time"

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

func Test_Index(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	t.Log("getting a symbol by a given index with ASCII-string")
	{
		r, err := NewGoRopeFromString(stringEn, 8)
		require.NoError(t, err)
		idx := r1.Intn(r.Size() - 1)
		c, err := r.Index(idx)
		require.NoError(t, err)
		require.Equal(t, c, string(stringEn[idx]))
	}

	t.Log("getting a symbol by a given index non-ASCII-string")
	{
		r, err := NewGoRopeFromString(stringRu, 8)
		require.NoError(t, err)
		idx := r1.Intn(r.Size() - 1)
		c, err := r.Index(idx)
		require.NoError(t, err)
		require.Equal(t, c, string([]rune(stringRu)[idx]))
	}
}

func Test_Concat(t *testing.T) {
	t.Log("concatenating two strings")
	{
		r1, err := NewGoRopeFromString(stringEn, 8)
		require.NoError(t, err)
		r2, err := NewGoRopeFromString(stringRu, 8)
		require.NoError(t, err)

		r3 := r1.Concat(r2)
		require.Equal(t, r3.String(), stringEn+stringRu)
	}
}
