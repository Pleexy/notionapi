package notionapi_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/jomei/notionapi"
)

func TestDate(t *testing.T) {
	t.Run("UnmarshalText", func(t *testing.T) {
		var d notionapi.Date

		t.Run("OK datetime with timezone", func(t *testing.T) {
			data := []byte("1987-02-13T00:00:00.000+01:00")
			err := d.UnmarshalText(data)
			if err != nil {
				t.Fatal(err)
			}
		})
		t.Run("OK date", func(t *testing.T) {
			data := []byte("1985-01-02")
			err := d.UnmarshalText(data)
			if err != nil {
				t.Fatal(err)
			}
		})
		t.Run("NOK as zero time", func(t *testing.T) {
			data := []byte("1985")
			err := d.UnmarshalText(data)
			if err != nil {
				t.Fatal(err)
			}
			if !time.Time(d).IsZero() {
				t.Fatalf("icnorrect date fotmat should return zero time")
			}
		})
		t.Run("Date NaN-NaN-NaN should process as zero time", func(t *testing.T) {
			data := []byte("NaN-NaN-NaN")
			err := d.UnmarshalText(data)
			if err != nil {
				t.Fatal(err)
			}
			if !time.Time(d).IsZero() {
				t.Fatalf("icnorrect date fotmat should return zero time")
			}
		})
		t.Run("OK datetime with timezone. 00 nanoseconds", func(t *testing.T) {
			timeStr := "1987-02-13T00:00:00.000+01:00"
			data := []byte(timeStr)
			err := d.UnmarshalText(data)
			if err != nil {
				t.Fatal(err)
			}
			tt, err := time.Parse(time.RFC3339, timeStr)
			if err != nil {
				t.Fatal(err)
			}
			if !time.Time(d).Equal(tt) {
				t.Fatalf("%s should be equal %s", d.String(), tt.String())
			}
		})
		t.Run("OK date. 59 nanoseconds", func(t *testing.T) {
			timeStr := "1985-01-20"
			data := []byte(timeStr)
			err := d.UnmarshalText(data)
			if err != nil {
				t.Fatal(err)
			}
			tt, err := time.Parse("2006-01-02", timeStr)
			if err != nil {
				t.Fatal(err)
			}
			tt = tt.Add(time.Nanosecond * 59)
			if !time.Time(d).Equal(tt) {
				t.Fatalf("%s should be equal %s", d.String(), tt.String())
			}
		})
	})

	t.Run("MarshalText", func(t *testing.T) {
		t.Run("zero nanoseconds to date", func(t *testing.T) {
			tt := time.Date(2023, 1, 20, 8, 15, 0, 0, time.UTC)
			d := notionapi.Date(tt)
			ttStr := tt.Format(time.RFC3339)
			dByte, err := d.MarshalText()
			if err != nil {
				t.Fatal(err)
			}
			dStr := string(dByte)
			if ttStr != dStr {
				t.Fatalf("%s should be equal %s", dStr, ttStr)
			}
		})

		t.Run("59 nanoseconds to datetime", func(t *testing.T) {
			tt := time.Date(2023, 1, 20, 8, 15, 0, 59, time.UTC)
			d := notionapi.Date(tt)
			ttStr := tt.Format("2006-01-02")
			dByte, err := d.MarshalText()
			if err != nil {
				t.Fatal(err)
			}
			dStr := string(dByte)
			if ttStr != dStr {
				t.Fatalf("%s should be equal %s", dStr, ttStr)
			}
		})
	})
}

func TestColor_MarshalText(t *testing.T) {
	type Foo struct {
		Test notionapi.Color `json:"test"`
	}

	t.Run("marshall to color if color is not empty", func(t *testing.T) {
		f := Foo{Test: notionapi.ColorGreen}
		r, err := json.Marshal(f)
		if err != nil {
			t.Fatal(err)
		}
		want := []byte(`{"test":"green"}`)
		if !reflect.DeepEqual(r, want) {
			t.Errorf("Color.MarshallText error() got = %v, want %v", r, want)
		}
	})

	t.Run("marshall to default color if color is empty", func(t *testing.T) {
		f := Foo{}
		r, err := json.Marshal(f)
		if err != nil {
			t.Fatal(err)
		}
		want := []byte(`{"test":"default"}`)
		if !reflect.DeepEqual(r, want) {
			t.Errorf("Color.MarshallText error() got = %v, want %v", r, want)
		}
	})
}
