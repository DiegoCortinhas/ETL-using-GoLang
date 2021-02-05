package etl

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConverter(t *testing.T) {

	file, _ := os.Open("RelExtraPorPeriodo.csv")

	defer file.Close()

	vo := extract(file)

	workDays := convert(vo)

	assert.NotNil(t, workDays, "workDays is null")
	assert.Equal(t, 72, len(workDays), "Expected length: 72")

	for i := 0; i < len(workDays); i++ {

		assert.NotNil(t, workDays[i].Employee)
		assert.NotNil(t, workDays[i].Date)
		assert.NotNil(t, workDays[i].Punches)
		assert.True(t, len(workDays[i].Punches) > 0)
	}
}