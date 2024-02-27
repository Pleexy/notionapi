package notionapi

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCreatedTimeProperty_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		data string
		want CreatedTimeProperty
	}{
		{
			name: "Create Time. Correct time string",
			data: `{
				"id": "12345",
				"type": "created_time",
				"created_time": "2023-02-02T00:00:00Z"
			}`,
			want: CreatedTimeProperty{
				ID:          "12345",
				Type:        PropertyTypeCreatedTime,
				CreatedTime: time.Date(2023, 02, 02, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Create Time. Incorrect time string",
			data: `{
				"id": "12345",
				"type": "created_time",
				"created_time": "{}"
			}`,
			want: CreatedTimeProperty{
				ID:          "12345",
				Type:        PropertyTypeCreatedTime,
				CreatedTime: time.Time{},
			},
		},
		{
			name: "Create Time. time is empty",
			data: `{
				"id": "12345",
				"type": "created_time"
			}`,
			want: CreatedTimeProperty{
				ID:          "12345",
				Type:        PropertyTypeCreatedTime,
				CreatedTime: time.Time{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got CreatedTimeProperty
			err := json.Unmarshal([]byte(tt.data), &got)
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("got = %+v, want = %+v", got, tt.want)
			}
			if got.GetID() != tt.want.ID.String() {
				t.Errorf("got = %+v, want = %+v", got, tt.want)
			}
		})
	}
}
