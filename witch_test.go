package witch

import (
    "testing"

    "github.com/i0Ek3/asrt"
)

func TestWitch(t *testing.T) {
    t.Run("test int with unsetted", func(t *testing.T){
        _, got, _ := Piu(3, 2, "ptr")
        want := true
        asrt.Equal(got, want)
    })

    t.Run("test int with setted", func(t *testing.T){
        _, got, _ := Piu(3, 2, "set")
        want := true
        asrt.Equal(got, want)
    })

    t.Run("test string with unsetted", func(t *testing.T){
        _, got, _ := Piu("hi", "hello", "ptr")
        want := true
        asrt.Equal(got, want)
    })

    t.Run("test string with setted", func(t *testing.T){
        _, got, _ := Piu("hi", "hello", "set")
        want := true
        asrt.Equal(got, want)
    })

    t.Run("test float with unsetted", func(t *testing.T){
        _, got, _ := Piu(3.14, 3.1415, "ptr")
        want := true
        asrt.Equal(got, want)
    })

    t.Run("test float with setted", func(t *testing.T){
        _, got, _ := Piu(3.14, 3.1415, "set")
        want := true
        asrt.Equal(got, want)
    })

    t.Run("test wrong string with unsetted", func(t *testing.T){
        _, got, _ := Piu("hi", 3, "ptr")
        want := false
        asrt.Equal(got, want)
    })

    t.Run("test wrong string with setted", func(t *testing.T){
        _, got, _ := Piu("hi", 2, "set")
        want := false
        asrt.Equal(got, want)
    })
}
