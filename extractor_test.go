package etl

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExtract(t *testing.T) {

	file, _ := os.Open("RelExtraPorPeriodo.csv")

	defer file.Close()

	vo := extract(file)

	assert.NotNil(t, vo, "vo is null")
	assert.Equal(t, 72, len(vo), "Expected length: 72")

	for i := 0; i < len(vo); i++ {

		assert.NotNil(t, vo[i].Hours)
		assert.NotNil(t, vo[i].Employee)
		assert.NotNil(t, vo[i].Date)
	}
}